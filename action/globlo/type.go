package globlo

type E机器人连续聊天接口 interface {
	E设定聊天内容(聊天内容 string)
	E清空对话()
	E发送消息(问题 string) string
	E发送图片消息(问题 string, 图片base64 string) string
	GetName() string
	GetModelName() string
	GetModelNames() []string
	SetModelName(ModelName string)
	SetName(Name string)
	Clone() E机器人连续聊天接口
}
type E机器人连续聊天 struct {
	E机器人 E机器人连续聊天接口
}

func New机器人连续聊天(机器人 E机器人连续聊天接口) *E机器人连续聊天 {
	return &E机器人连续聊天{
		E机器人: 机器人,
	}
}

func (c E机器人连续聊天) E设定聊天内容(聊天内容 string) {
	c.E机器人.E设定聊天内容(聊天内容)
}
func (c E机器人连续聊天) E清空对话() {
	c.E机器人.E清空对话()
}
func (c E机器人连续聊天) E发送消息(问题 string) string {
	return c.E机器人.E发送消息(问题)
}
func (c E机器人连续聊天) E发送图片消息(问题 string, 图片base64 string) string {
	return c.E机器人.E发送图片消息(问题, 图片base64)
}
func (c E机器人连续聊天) GetName() string {
	return c.E机器人.GetName()
}
func (c E机器人连续聊天) GetModelNames() []string {
	return c.E机器人.GetModelNames()
}
func (c E机器人连续聊天) GetModelName() string {
	return c.E机器人.GetModelName()
}
func (c E机器人连续聊天) SetModelName(ModelName string) {
	c.E机器人.SetModelName(ModelName)
}

func (c E机器人连续聊天) SetName(name string) {
	c.E机器人.SetName(name)
}
func (c *E机器人连续聊天) Clone() *E机器人连续聊天 {
	return &E机器人连续聊天{
		E机器人: c.E机器人.Clone(),
	}
}
