package main

import (
	"context"
	"huhobot/internal/handlers"
	"log"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/token"
)

func main() {
	// 初始化token
	botToken := token.New(token.QQBot, "your_app_id", "your_app_secret")

	// 创建API实例
	api := botgo.NewOpenAPI(botToken)

	// 创建websocket连接
	ws, err := api.WS(context.Background(), nil, "")
	if err != nil {
		log.Fatalf("WS错误: %v", err)
	}

	// 初始化处理器
	handler := handlers.NewBot(api)

	// 启动事件监听
	handler.Start(ws)
}
