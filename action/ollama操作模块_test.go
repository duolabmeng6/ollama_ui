package action

import (
	"testing"
)

func TestOllama操作_E获取模型列表(t *testing.T) {
	m := NeOllama操作()
	data := m.E获取模型列表()
	println(data)

}

func TestOllama操作_E复制模型(t *testing.T) {
	m := NeOllama操作()
	data := m.E复制模型("qwen2x:latest", "qwen:latest")
	println(data)

}

func TestOllama操作_下载模型(t *testing.T) {
	//fmt.Println(字节转友好字符串(2142590208)) // 输出: "2.00 GB"
	//fmt.Println(字节转友好字符串(1024))       // 输出: "1.00 KB"
	//fmt.Println(字节转友好字符串(1500000))    // 输出: "1.43 MB"
	//fmt.Println(字节转友好字符串(1500))       // 输出: "1.46 KB"
	//fmt.Println(字节转友好字符串(1))          // 输出: "1.00 B"
	m := NeOllama操作()
	data := m.E下载模型("qwen2:0.5b", func(总大小, 已下载, 总进度 string) {
		println(总大小, 已下载, 总进度)
	})
	println(data)

}
func TestOllama操作_E删除模型(t *testing.T) {
	m := NeOllama操作()
	data := m.E删除模型("phi:2.7b")
	println(data)

}
func TestOllama操作_E对话(t *testing.T) {
	m := NeOllama操作()
	data := m.E对话("qwen2:0.5b", "你是什么模型")
	println(data)

}
