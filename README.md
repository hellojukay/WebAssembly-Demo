# WebAssembly-Demo

本项目展示如何在 javascript 中使用 Go 导出的 WebAssembly 模块

## 编译 WebAssembly
```shell
GOOS=js GOARCH=wasm go build -o main.wasm
```
复制运行环境到本地
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js"  .
```
## 启动项目
```
live-server
```

# doc
https://pkg.go.dev/text/template#hdr-Pipelines