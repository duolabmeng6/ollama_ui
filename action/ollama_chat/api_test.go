package ollama_chat

import (
	"fmt"
	"testing"
)

func Test_连续聊天3(t *testing.T) {
	chatbot := New机器人连续聊天("http://localhost:11434/api/chat", "",
		"qwen2:0.5b",
	)
	chatbot.E清空对话()
	var 回答 string
	chatbot.E设定聊天内容("你是是一个中文翻译助手我给你发送的内容你都翻译为中文")
	回答 = chatbot.E发送消息("I Like Apple")
	fmt.Println(回答)
	//回答 = chatbot.E发送消息("谁是最大的数？")
	//fmt.Println(回答)
	//回答 = chatbot.E发送消息("谁是最小的数？")
	//fmt.Println(回答)
	//回答 = chatbot.E发送消息("从小到大排序")
	//fmt.Println(回答)
	//回答 = chatbot.E发送消息("从大到小排序")
	//fmt.Println(回答)
}
