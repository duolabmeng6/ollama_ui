<p align="center" style="text-align: center">

</p>

<p align="center">
</p>

<div align="center">
<strong>
<samp>

[English](README.md) · [简体中文](README.zh-Hans.md) 

</samp>
</strong>
</div>


# Demonstration project for GoEasyDesigner window designer

Used to run the interface and program designed by the window designer

![GIF](https://user-images.githubusercontent.com/59047063/270091148-a89d2ab9-9ba3-4efc-b0fa-0a7dcc3bcfc1.gif)

Create project:

```
wails init -n "go-easy-demo" -t <https://github.com/duolabmeng6/wails-template-vue-go-easy>
```

Run window:

```
cd go-easy-demo
wails dev
```


Download GoEasyDesigner window designer (<https://github.com/duolabmeng6/GoEasyDesigner/releases>).

GoEasyDesigner project address (<https://github.com/duolabmeng6/GoEasyDesigner>)

# Online experience of window designer

The web page only provides basic interface design. If you need better coding experience, please download the client.

Chinese address: <https://go.kenhong.com/>

Foreign address: <https://g.yx24.me>

Click save after designing the interface. It will download two files (`design.json`, `__aux_code.js`). Please note that your browser allows downloading multiple files.

Download the code of this project and copy the `go-easy-demo` folder as the development project.

`go-easy-demo/frontend/src/win/design.json`

`go-easy-demo/frontend/src/win/__aux_code.js`

Run the project to see the interface you designed.

```
cd go-easy-demo
wails dev
```