package handlers

import (
	"context"
	"log"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/websocket"
)

type Bot struct {
	api openapi.OpenAPI
}

func NewBot(api openapi.OpenAPI) *Bot {
	return &Bot{api: api}
}

func (b *Bot) Start(client *websocket.Client) {
	// 设置intents
	intent := dto.IntentGuildMessages |
		dto.IntentGuildMembers |
		dto.IntentGuildMessageReactions

	// 启动事件监听
	client.Start(context.Background(), b, b, intent)
}

// 实现事件处理接口
func (b *Bot) OnReady(event *dto.WSPayload, data *dto.WSReadyData) {
	log.Println("机器人已就绪")
}

func (b *Bot) OnMessageCreate(event *dto.WSPayload, data *dto.WSGuildMessageData) {
	log.Printf("收到消息: %s", data.Content)
}
