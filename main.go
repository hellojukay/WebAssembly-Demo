package main

import (
	"encoding/json"
	"html/template"
	"strings"
	"syscall/js"
	"time"
)

func fib(this js.Value, args []js.Value) any {
	var n = args[0].Int()
	if n <= 2 {
		return 1
	}
	// 这里类型转换， any 转 int
	return add(this, []js.Value{js.ValueOf(n - 1)}).(int) + add(this, []js.Value{js.ValueOf(n - 2)}).(int)
}

func add(this js.Value, args []js.Value) any {
	return args[0].Float() + args[1].Float()
}

// 根据模板和数据，渲染字符串
func render(this js.Value, args []js.Value) any {
	var text = args[0].String()
	var jsonText = args[1].String()
	var data = make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonText), &data); err != nil {
		return err.Error()
	}
	data["builtin"] = 1000
	data["builtinfunc"] = func() string {
		return time.Now().Format(`2006-01-02T15:04:05.000Z`)
	}
	t := template.New("test")
	t, err := t.Parse(text)
	if err != nil {
		return err.Error()
	}
	var buf strings.Builder
	if err = t.Execute(&buf, data); err != nil {
		return err.Error()
	}
	return buf.String()
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("addFun", js.FuncOf(add))
	js.Global().Set("fib", js.FuncOf(fib))
	js.Global().Set("template", js.FuncOf(render))
	js.Global().Set("name", "licong")
	<-done
}
