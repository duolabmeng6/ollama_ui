package ollama_chat

import (
	"changeme/action/globlo"
	"github.com/duolabmeng6/goefun/ecore"
)

type E机器人连续聊天 struct {
	问答列表          []*H
	问答轮数          int
	系统提示          *H
	SecretKey     string
	Name          string
	ModelName     string
	ApiServer     string
	ModelNameList []string
}

func New机器人连续聊天(apiServer string, secretKey string, modelName string) *E机器人连续聊天 {
	o := &E机器人连续聊天{
		问答轮数:      5,
		SecretKey: secretKey,
		ApiServer: apiServer,
	}
	o.问答列表 = []*H{} // 初始化为空的有序map切片
	o.系统提示 = NewH() // 使用NewH创建一个新的有序map

	//使用逗号分割 modelName
	o.ModelNameList = ecore.E分割文本(modelName, ",")
	o.SetModelName(o.ModelNameList[0])
	o.SetName("ChatGptCustomize")

	o.E清空对话()
	return o
}
func (c *E机器人连续聊天) Clone() globlo.E机器人连续聊天接口 {
	newChatbot := *c
	return &newChatbot
}
func (c *E机器人连续聊天) GetName() string {
	return c.Name
}

func (c *E机器人连续聊天) SetName(Name string) {
	c.Name = Name
}

func (c *E机器人连续聊天) E设定聊天内容(聊天内容 string) {
	c.问答列表 = []*H{} // 初始化为空的有序map切片
	c.系统提示 = NewH() // 使用NewH创建一个新的有序map
	c.系统提示.Set("role", "system")
	c.系统提示.Set("content", 聊天内容)

	//c.问答列表 = []H{}
	//c.系统提示 = H{"role": "system", "content": 聊天内容}
	//c.问答列表 = append(c.问答列表, H{"role": "system", "content": 聊天内容})
}

func (c *E机器人连续聊天) E清空对话() {
	c.E设定聊天内容("你是AI.请直接回答.使用markdown格式回答")
}

func (c *E机器人连续聊天) _获取机器人回答(内容 []*H) string {
	message, _ := 连续聊天(内容, c.ApiServer, c.SecretKey, c.ModelName)
	return message
}

func (c *E机器人连续聊天) E发送消息(问题 string) string {
	用户问题 := NewH()
	用户问题.Set("role", "user")
	用户问题.Set("content", 问题)
	c.问答列表 = append(c.问答列表, 用户问题)

	// 创建一个新的切片来存储有序的问答列表
	临时问答列表 := []*H{c.系统提示}
	临时问答列表 = append(临时问答列表, c.问答列表...)

	机器人回答 := c._获取机器人回答(临时问答列表)

	机器人回答H := NewH()
	机器人回答H.Set("role", "assistant")
	机器人回答H.Set("content", 机器人回答)
	c.问答列表 = append(c.问答列表, 机器人回答H)
	if len(c.问答列表) > c.问答轮数*2 {
		c.问答列表 = c.问答列表[2:]
	}

	// ecore.E调试输出等调试代码可以保持不变
	ecore.E调试输出("问答列表", c.问答列表)
	ecore.E调试输出("问答次数", len(c.问答列表))

	return 机器人回答
}

func (c *E机器人连续聊天) E发送图片消息(问题 string, 图片base64 string) string {
	//     {
	//        "role": "user", "content":
	//                [
	//                  {"type": "text", "text": "What’s in this image?"},
	//                  {
	//                    "type": "image_url",
	//                    "image_url": {
	//                      "url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/dd/Gfp-wisconsin-madison-the-nature-boardwalk.jpg/2560px-Gfp-wisconsin-madison-the-nature-boardwalk.jpg"
	//                    }
	//                  }
	//              ]
	//    }
	图片问题 := []map[string]interface{}{
		//"role": "user",
		//"content": []interface{}{
		map[string]interface{}{
			"type": "text",
			"text": 问题,
		},
		map[string]interface{}{
			"type": "image_url",
			"image_url": map[string]interface{}{
				"url": 图片base64,
			},
		},
		//},
	}
	用户问题 := NewH()
	用户问题.Set("role", "user")
	用户问题.Set("content", 图片问题)

	c.问答列表 = append(c.问答列表, 用户问题)

	// 创建一个新的切片来存储有序的问答列表
	//临时问答列表 := []*H{c.系统提示}
	临时问答列表 := []*H{}
	临时问答列表 = append(临时问答列表, c.问答列表...)

	机器人回答 := c._获取机器人回答(临时问答列表)

	机器人回答H := NewH()
	机器人回答H.Set("role", "assistant")
	机器人回答H.Set("content", 机器人回答)
	c.问答列表 = append(c.问答列表, 机器人回答H)
	if len(c.问答列表) > c.问答轮数*2 {
		c.问答列表 = c.问答列表[2:]
	}

	// ecore.E调试输出等调试代码可以保持不变
	ecore.E调试输出("问答列表", c.问答列表)
	ecore.E调试输出("问答次数", len(c.问答列表))

	return 机器人回答
}

func (c *E机器人连续聊天) GetModelNames() []string {
	return c.ModelNameList
}
func (c *E机器人连续聊天) SetModelName(ModelName string) {
	c.ModelName = ModelName
}
func (c *E机器人连续聊天) GetModelName() string {
	return c.ModelName
}
