package ws

import (
	"ServerDog/info"
	"ServerDog/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Message 自定义消息
type Message struct {
	Status     bool
	Connection string
	SystemInfo info.SystemInfo
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Websocket连接池
var websocketPool map[string]*websocket.Conn

// MessageChannel Websocket消息通道
var MessageChannel chan Message

func init() {
	logfile, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	multiWriter := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime)

	websocketPool = make(map[string]*websocket.Conn)
	MessageChannel = make(chan Message)
}

/*
Websocket客户端
*/

// CreateWebsocket 发起Websocket连接
func CreateWebsocket(addr string) error {
	dialer := websocket.Dialer{}

	// 设置超时时间
	dialer.HandshakeTimeout = time.Second * 5
	conn, _, err := dialer.Dial(addr, nil)
	if err != nil {
		return err
	}

	// 将连接放入连接池
	websocketPool[conn.RemoteAddr().String()] = conn
	go readMessage(conn, MessageChannel)

	return nil
}

// readMessage 读取Websocket消息并放入消息通道
func readMessage(connection *websocket.Conn, messageChan chan Message) {
	addr := strings.Split(connection.RemoteAddr().String(), ":")[0]

	ipInfo := utils.GetIpInfo(addr)

	for {
		message := Message{
			Status:     false,
			Connection: addr,
		}

		_, rm, err := connection.ReadMessage()
		if err != nil {
			log.Println("Failed to read message from ", connection.RemoteAddr())
			delete(websocketPool, connection.RemoteAddr().String())
			_ = connection.Close()

			messageChan <- message
			return
		}

		var systemInfo info.SystemInfo
		err = json.Unmarshal(rm, &systemInfo)
		if err != nil {
			log.Println("Failed to unmarshal message: ", err)
			delete(websocketPool, connection.RemoteAddr().String())
			_ = connection.Close()

			messageChan <- message
			return
		}

		if ip, ok := ipInfo["ip"]; ok {
			systemInfo.Host.IP = ip
		}

		if location, ok := ipInfo["location"]; ok {
			systemInfo.Host.Location = location
		}

		if provider, ok := ipInfo["provider"]; ok {
			systemInfo.Host.Provider = provider
		}

		message.Status = true
		message.SystemInfo = systemInfo
		messageChan <- message
	}
}

/*
Websocket服务端
*/

// StartWebsocketServer 启动Websocket服务器
func StartWebsocketServer(port int) {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	http.HandleFunc("/info", handleWebSocketConnection)
	log.Println("Server dog's websocket is listening at: ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
		return
	}
}

// 处理每个连接
func handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection: ", err)
		closeConnection(conn)
		return
	}
	defer closeConnection(conn)

	for {
		systemInfo := info.GetSystemInfo()
		sm, err := json.Marshal(systemInfo)
		if err != nil {
			log.Println("Failed to marshal message: ", err)
			closeConnection(conn)
			return
		}

		err = conn.WriteMessage(websocket.TextMessage, sm)
		if err != nil {
			log.Println("Failed to write message to client: ", err)
			closeConnection(conn)
			return
		}

		time.Sleep(1100 * time.Millisecond)
	}
}

// closeConnection 关闭连接
func closeConnection(conn *websocket.Conn) {
	_ = conn.Close()
	log.Println("Connection closed: ", conn.RemoteAddr())
}
