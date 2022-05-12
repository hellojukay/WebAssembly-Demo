# WebAssembly-Demo

本项目展示如何在 javascript 中使用 Go 导出的 WebAssembly 模块

## 编译 WebAssembly
```shell
GOOS=js GOARCH=wasm go build -o main.wasm
```