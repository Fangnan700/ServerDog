package main

import (
	"ServerDog/ws"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	wsPort int
	uiPort int
)

//go:embed web
var web embed.FS
var services []ws.Message

func init() {
	logfile, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	multiWriter := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime)

	flag.IntVar(&uiPort, "P", 7776, "web panel port.")
	flag.IntVar(&wsPort, "p", 7777, "websocket listening port.")
	flag.Parse()
}

func main() {
	go ws.StartWebsocketServer(wsPort)
	go update()

	/*
		建立本机Websocket连接
	*/
	time.Sleep(1 * time.Second)
	wsAddr := fmt.Sprintf("ws://127.0.0.1:%d/info", wsPort)
	err := ws.CreateWebsocket(wsAddr)
	if err != nil {
		panic(err)
	}

	/*
		创建http服务
	*/
	uiAddr := fmt.Sprintf("0.0.0.0:%d", uiPort)

	// 开启静态文件服务
	root, _ := fs.Sub(web, "web")
	http.Handle("/", http.FileServer(http.FS(root)))

	// 更新数据
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := json.Marshal(services)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonData)
	})

	// 添加服务器
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		var data map[string]string
		body, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(body, &data)

		verifyAddr := fmt.Sprintf("http://%s/verify", data["addr"])

		client := &http.Client{
			Timeout: time.Second * 5, // 设置超时时间为10秒
		}

		resp, err := client.Get(verifyAddr)
		if err == nil {
			body, _ = io.ReadAll(resp.Body)
			text := string(body)

			if text == "yes!this is fangnan700's code!" {
				websocketAddr := fmt.Sprintf("ws://%s/info", data["addr"])
				err = ws.CreateWebsocket(websocketAddr)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		w.WriteHeader(http.StatusBadRequest)
	})

	// 验证服务
	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("yes!this is fangnan700's code!"))
	})

	log.Printf("Server dog's ui panel is runing at: %s", uiAddr)
	_ = http.ListenAndServe(uiAddr, nil)
}

// update 更新services列表
func update() {
	for {
		select {
		case message := <-ws.MessageChannel:
			if len(services) > 0 {
				var step = 0
				for index := range services {
					if services[index].Connection == message.Connection {
						services[index] = message
						if !services[index].Status {
							services = append(services[:index], services[index+1:]...)
						}
						break
					}
					if step >= len(services)-1 {
						services = append(services, message)
						break
					}
					step += 1
				}
			} else {
				services = append(services, message)
			}
		}
	}
}
