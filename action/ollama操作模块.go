package action

import (
	"bufio"
	"changeme/action/ollama_chat"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/ehttp"
	"github.com/duolabmeng6/goefun/etool"
	"io"
	"math"
	"net/http"
	"strings"
)

type ModelDetails struct {
	ParentModel       string   `json:"parent_model"`
	Format            string   `json:"format"`
	Family            string   `json:"family"`
	Families          []string `json:"families"`
	ParameterSize     string   `json:"parameter_size"`
	QuantizationLevel string   `json:"quantization_level"`
}

type Model struct {
	Name       string       `json:"name"`
	Model      string       `json:"model"`
	ModifiedAt string       `json:"modified_at"`
	Size       int64        `json:"size"`
	Digest     string       `json:"digest"`
	Details    ModelDetails `json:"details"`
}

type ModelResponse struct {
	Models []Model `json:"models"`
}
type Ollama操作接口 interface {
	E获取模型列表() string
	E下载模型(模型名称 string) string
	E删除模型(模型名称 string) string
	E复制模型(模型名称 string, 目标名称 string) string
	E停止下载() string
	E对话(模型名称 string, 内容 string) string
	E搜索模型(模型名称 string) string
}

type Ollama操作 struct {
	Ollama操作接口
	停止下载   bool
	E服务器地址 string
}

func NeOllama操作() *Ollama操作 {
	m := new(Ollama操作)
	m.E服务器地址 = "http://localhost:11434"
	return m
}

func (this *Ollama操作) E获取模型列表() string {
	eh := ehttp.NewHttp()
	response, err := eh.Get(this.E服务器地址 + "/api/tags")

	var modelResponse ModelResponse
	err = json.Unmarshal([]byte(response), &modelResponse)
	if err != nil {
		return `[]` // 返回空数组，如果解析失败
	}

	result := make([]map[string]interface{}, 0)
	for _, model := range modelResponse.Models {
		modelInfo := map[string]interface{}{
			"name": model.Name,
			"id":   model.Digest[:8], // 使用摘要的前8个字符作为ID
			"size": formatSize(model.Size),
			"time": formatTime(model.ModifiedAt),
		}
		result = append(result, modelInfo)
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		return `[]` // 返回空数组，如果序列化失败
	}

	return string(jsonResult)
}

// 辅助函数，用于格式化大小
func formatSize(size int64) string {
	// 实现大小格式化逻辑，例如将字节转换为GB
	// 这里简化处理，直接返回GB
	return fmt.Sprintf("%.2fGB", float64(size)/(1024*1024*1024))
}

// 辅助函数，用于格式化时间
func formatTime(timeStr string) string {
	t := ecore.E到时间(timeStr)

	return t.E到友好时间()
}

func (this *Ollama操作) E下载模型(模型名称 string, fn func(总大小, 已下载, 总进度 string)) string {
	url := this.E服务器地址 + "/api/pull"
	payload := fmt.Sprintf(`{"name": "%s", "stream": true}`, 模型名称)
	this.停止下载 = false
	resp, err := http.Post(url, "application/json", strings.NewReader(payload))
	if err != nil {
		return fmt.Sprintf("请求错误: %v", err)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		if this.停止下载 {
			return "停止下载"
		}
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Sprintf("读取响应错误: %v", err)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(line, &response); err != nil {
			return fmt.Sprintf("解析JSON错误: %v", err)
		}

		status, ok := response["status"].(string)
		if !ok {
			continue
		}
		switch status {
		case "pulling manifest":
			fmt.Println("开始拉取模型...")
		case "verifying sha256 digest":
			fmt.Println("验证SHA256摘要...")
		case "writing manifest":
			fmt.Println("写入清单...")
		case "removing any unused layers":
			fmt.Println("移除未使用的层...")
		case "success":
			return "模型下载成功"
		default:
			if strings.HasPrefix(status, "pulling") {
				//digest := etool.Json解析(ecore.E到文本(line), "digest").String()
				total := etool.Json解析(ecore.E到文本(line), "total").Int()
				completed := etool.Json解析(ecore.E到文本(line), "completed").Int()
				progress := float64(completed) / float64(total) * 100
				//println(total, completed, 字节转友好字符串(total), 字节转友好字符串(completed))
				//fmt.Printf("pulling 下载进度 (%s): %.2f%%\n", digest, progress)
				fn(字节转友好字符串(total), 字节转友好字符串(completed), fmt.Sprintf("%.2f", progress))
			}
			if strings.HasPrefix(status, "downloading") {
				digest := response["digest"].(string)
				total := int64(response["total"].(float64))
				completed := int64(response["completed"].(float64))
				progress := float64(completed) / float64(total) * 100
				fmt.Printf("下载进度 (%s): %.2f%%\n", digest, progress)
				fn(字节转友好字符串(total), 字节转友好字符串(completed), fmt.Sprintf("%.2f", progress))
			}
		}
	}

	return "下载过程异常结束"
}
func (this *Ollama操作) E停止下载() string {
	this.停止下载 = true
	return "停止下载"
}
func (this *Ollama操作) E复制模型(模型名称 string, 目标名称 string) string {
	eh := ehttp.NewHttp()
	response, _ := eh.Post(this.E服务器地址+"/api/copy", `{
  "source": "`+模型名称+`",
  "destination": "`+目标名称+`"
}`)
	if eh.E取状态码() == 200 {
		return "成功"
	}

	return etool.Json解析文本(response, "error")
}

func (this *Ollama操作) E对话(模型名称 string, 内容 string) string {
	chatbot := ollama_chat.New机器人连续聊天(this.E服务器地址+"/api/chat", "",
		模型名称,
	)
	chatbot.E清空对话()
	var 回答 string
	chatbot.E设定聊天内容("你是助手")
	回答 = chatbot.E发送消息(内容)
	fmt.Println(回答)

	return 回答
}

func (this *Ollama操作) E删除模型(name string) string {
	url := this.E服务器地址 + "/api/delete"
	payload := fmt.Sprintf(`{"name": "%s"}`, name)

	// 创建一个新的请求
	req, err := http.NewRequest("DELETE", url, strings.NewReader(payload))
	if err != nil {
		return fmt.Sprintf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode == 200 {
		return "模型删除成功"
	} else if resp.StatusCode == 404 {
		return "模型不存在"
	} else {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Sprintf("删除失败，状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}
}

func 字节转友好字符串(字节 int64) string {
	if 字节 == 0 {
		return "0 B"
	}

	单位 := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	i := int(math.Floor(math.Log(float64(字节)) / math.Log(1024)))

	// 如果 i 超出了单位数组的范围，就使用最大的单位
	if i > len(单位)-1 {
		i = len(单位) - 1
	}

	大小 := float64(字节) / math.Pow(1024, float64(i))

	// 根据大小决定小数点后的位数
	var 格式化字符串 string
	if 大小 >= 100 {
		格式化字符串 = "%.0f %s"
	} else if 大小 >= 10 {
		格式化字符串 = "%.1f %s"
	} else {
		格式化字符串 = "%.2f %s"
	}

	return fmt.Sprintf(格式化字符串, 大小, 单位[i])
}

type ModelInfo struct {
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	UpdatedTime string   `json:"updated_time"`
}

func extractModelInfo(html string) ([]ModelInfo, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return []ModelInfo{}, err
	}

	var info ModelInfo
	var infos []ModelInfo

	doc.Find("#repo li").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find("h2 span").Text())
		info.Name = name

		// nfo.Sizes = string[]{s.Find(".bg-\\[\\#ddf4ff\\]").Text()}
		info.Tags = s.Find(".bg-\\[\\#ddf4ff\\]").Map(func(i int, sel *goquery.Selection) string {
			return strings.TrimSpace(sel.Text())
		})

		info.Description = strings.TrimSpace(s.Find("p.max-w-md").Text())

		updatedTime := strings.TrimSpace(s.Find("span:contains('Updated')").Text())
		info.UpdatedTime = strings.TrimSpace(strings.Replace(updatedTime, "Updated", "", -1))
		infos = append(infos, info)
	})

	return infos, nil
}

func (this *Ollama操作) E搜索模型(模型名称 string) string {
	//访问 https://ollama.com/library?q=phi&sort=featured
	url := fmt.Sprintf("https://ollama.com/library?q=%s&sort=featured", 模型名称)
	eh := ehttp.NewHttp()
	response, err := eh.Get(url)
	if err != nil {
		return ""
	}
	//ecore.E写到文件("1.html", []byte(response))
	//response := ecore.E读入文本("1.html")
	data, _ := extractModelInfo(response)
	ecore.E调试输出(data)
	return etool.E到Json(data)
}
