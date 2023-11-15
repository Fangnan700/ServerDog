package utils

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	logfile, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	multiWriter := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime)
}

func GetIpInfo(ip string) map[string]string {

	var target string
	if ip == "127.0.0.1" {
		target = "https://ip.chinaz.com/"
	} else {
		target = "https://ip.chinaz.com/" + ip
	}

	resp, err := http.Get(target)
	if err != nil {
		log.Println(err)
		return nil
	}

	data, _ := io.ReadAll(resp.Body)
	doc, _ := html.Parse(strings.NewReader(string(data)))

	targetID := "infoLocation"
	targetClasses := []string{"Whwtdhalf", "w15-0", "lh45"}

	textContent := extractTextContentByID(doc, targetID)
	textContents := extractTextContentsByClasses(doc, targetClasses)

	result := make(map[string]string)

	// 公网IP
	result["ip"] = textContents[0]
	// 服务商
	result["provider"] = textContents[3]
	// 归属地
	result["location"] = textContent

	return result
}

func extractTextContentsByClasses(n *html.Node, targetClasses []string) []string {
	var result []string

	if n.Type == html.ElementNode {
		// 检查元素是否具有指定的所有类
		hasAllClasses := true
		for _, targetClass := range targetClasses {
			hasClass := false
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, targetClass) {
					hasClass = true
					break
				}
			}
			if !hasClass {
				hasAllClasses = false
				break
			}
		}

		// 如果元素具有指定的所有类，提取文本内容
		if hasAllClasses {
			// 如果是文本节点，则提取文本内容
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				result = append(result, n.FirstChild.Data)
			}
		}
	}

	// 递归处理子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, extractTextContentsByClasses(c, targetClasses)...)
	}

	return result
}

func extractTextContentByID(n *html.Node, targetID string) string {
	var result string

	if n.Type == html.ElementNode {
		// 检查元素是否具有指定 id
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == targetID {
				// 如果是文本节点，则提取文本内容
				if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
					result = n.FirstChild.Data
				}
				break
			}
		}
	}

	// 递归处理子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result += extractTextContentByID(c, targetID)
	}

	return result
}
