import {ElMessage, ElMessageBox} from "element-plus";
import * as systemFc from "../../wailsjs/runtime";
import * as goFc from "../../wailsjs/go/main/App";
import {__load_data} from './__load_data'

export function BindWindowEvent() {
    const c = __load_data()
    let comps = c.comps
    c.WinCreated = async function () {
        console.log("Win创建完毕")
        comps.Win.text = "OllamaManager " + await goFc.GetVersion();
        comps.选择夹3.value = "0"
        comps.表格1.data = []
        c.按钮_刷新模型列表被单击()

        systemFc.EventsOn("downInfo", function (data) {
            let jsondata = JSON.parse(data);
            console.log(jsondata)
            //{"total_size":"%s","downloaded":"%s","progress":"%s"}
            comps.Progress1.percentage = jsondata['progress']
            comps.标签_下载进度.text = "总大小:"+jsondata['total_size']+" / 已下载:"+jsondata['downloaded']

        });


    }
    c.按钮_刷新模型列表被单击 = async function () {
        console.log("按钮_刷新模型列表被单击")
        let data = await goFc.E获取模型列表()
        console.log("表格数据",data)
        data = JSON.parse(data)
        console.log("表格数据",data)

        comps.表格1.data = data

        ElMessage.success('已刷新');

        comps.Select_模型名称.options = []
        for (let i = 0; i < data.length; i++) {
            comps.Select_模型名称.options.push({
                label: data[i].name,
                value: data[i].name,
            })
        }

    }



    c.按钮_删除模型被单击 = async function () {
        console.log("按钮_删除模型被单击")
        // goFc.E删除模型(comps.)
        let list = comps.表格1.value
        for (let i = 0; i < list.length; i++) {
            let data = await goFc.E删除模型(list[i].name)
            ElMessage.success(data);
        }
        c.按钮_刷新模型列表被单击()


    }

    c.按钮_下载被单击 = async function () {
        console.log("按钮_下载被单击")
        if (comps.按钮_下载.text == "下载模型") {
            comps.按钮_下载.text = "停止下载模型"
            comps.标签_下载进度.text = "正在下载"
            let 下载状态 = await goFc.E下载模型(comps.编辑框_下载模型名称.text)
            comps.标签_下载进度.text = 下载状态
            comps.按钮_下载.text = "下载模型"
        }else{
            comps.按钮_下载.text = "下载模型"
            await goFc.E停止下载()
        }


    }


    c.按钮_模型改名被单击 = async function () {
        console.log("按钮_模型改名被单击")
        let list = comps.表格1.value
        var data = ""
        console.log(list)
        if (list.length === 1){
            data = list[0].name
        }

        ElMessageBox.prompt("请输入新的模型名称", "提示", {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            inputErrorMessage: '模型名称',
        }).then(async ({value}) => {
            let data2 = await goFc.E模型改名(data,value)
            ElMessage.success(data2);
            c.按钮_刷新模型列表被单击()

        })
        
    }

    c.按钮_发送对话被单击 = async function () {
        console.log("按钮_发送对话被单击")
        comps.编辑框_显示对话结果.text = comps.编辑框_显示对话结果.text + "用户: " + comps.编辑框_发送对话.text + "\r\n"
        comps.编辑框_显示对话结果.text = comps.编辑框_显示对话结果.text + "AI: "
        let data = await goFc.E对话(comps.Select_模型名称.value,comps.编辑框_显示对话结果.text)
        console.log("返回结果",data)
        comps.编辑框_显示对话结果.text = comps.编辑框_显示对话结果.text + data + "\r\n"



        
    }

    c.按钮_清空被单击 = async function () {
        console.log("按钮_清空被单击")
        comps.编辑框_显示对话结果.text = ""
        
    }

    c.标签1被单击 = async function () {
        console.log("标签1被单击")
        systemFc.BrowserOpenURL("https://ollama.com/library")
    }

    c.按钮_检查更新被单击 = async function () {
        console.log("按钮_检查更新被单击")
        ElMessage.success('检查更新中');
        await goFc.E检查更新();
    }
//Don't delete the event function flag
}