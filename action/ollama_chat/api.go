package ollama_chat

import (
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/ehttp"
	"github.com/duolabmeng6/goefun/etool"
)

type H struct {
	data map[string]interface{}
	keys []string
}

// NewH 创建并返回一个新的有序map H
func NewH() *H {
	return &H{
		data: make(map[string]interface{}),
	}
}

// Set 方法在H中设置键值对
// 如果键不存在，则添加到keys切片中
func (h *H) Set(key string, value interface{}) {
	if _, ok := h.data[key]; !ok {
		h.keys = append(h.keys, key)
	}
	h.data[key] = value
}

// Get 方法从H中获取值
func (h *H) Get(key string) (interface{}, bool) {
	val, ok := h.data[key]
	return val, ok
}

// Keys 方法返回H中所有的键，按照插入顺序
func (h *H) Keys() []string {
	return h.keys
}

func 连续聊天(content []*H, ApiServer string, secret_key string, model string) (string, error) {

	eh := ehttp.NewHttp()
	// 将content转换为有序的键值对序列
	orderedContent := make([]map[string]interface{}, len(content))
	for i, h := range content {
		messageMap := make(map[string]interface{})
		for _, key := range h.Keys() {
			value, _ := h.Get(key)
			messageMap[key] = value
		}
		orderedContent[i] = messageMap
	}

	url := ApiServer

	data := map[string]interface{}{
		"model":    model,
		"stream":   false,
		"messages": orderedContent,
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + secret_key,
	}
	//打印data的json数据
	ecore.E调试输出(etool.Json美化(data))

	ret, err := eh.Post(url, data, headers)
	println(ret)
	if err != nil {
		fmt.Printf("err:%v", err)
		return "服务器访问错误", err
	}
	if eh.Response.StatusCode != 200 {
		return "服务器返回错误 状态码:" + eh.Response.Status, err
	}
	回答 := etool.Json解析文本(ret, "message.content")
	if 回答 == "" {
		return "取回答案失败", nil
	}
	return 回答, nil
}
