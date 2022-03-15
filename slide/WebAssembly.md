---
marp: true
---
# 起源

```javascript
sum(1,3)

function sum(a,b) {
    return a + b
}
```
---
![js_compile](assets/js_compile.png)
---
浏览器的的 v8 引擎性能已经达到了瓶颈，需要一种新的方式加速 web 端执行性能
![img](assets/node_benchmark.png)
---
# ASM.js
ASM 是 Javascript 的严格的子集，执行速度快
+ https://github.com/emscripten-core/emscripten
+ https://www.ruanyifeng.com/blog/2017/09/asmjs_emscripten.html
![asm.js_benchmark](assets/asm.js_benchmark.png)
---
# WebAssembly是什么
WebAssembly是一种新的编码方式，可以在现代的网络浏览器中运行 。
+ 二进制标准
+ 目标语言
---
# 现状
+ https://www.assemblyscript.org/ (typescript转WebAssembly)
+ https://wasi.dev/ (WebAssembly System Interface)
+ https://wapm.io/ (WebAssembly的包管理中心)
