package main

import (
	"encoding/json"
	"html/template"
	"strings"
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	return args[0].Float() + args[1].Float()
}

// 更具 map 渲染字符串
func render(this js.Value, args []js.Value) interface{} {
	var text = args[0].String()
	var jsonText = args[1].String()
	var data = make(map[string]string)
	if err := json.Unmarshal([]byte(jsonText), &data); err != nil {
		return err.Error()
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
	js.Global().Set("template", js.FuncOf(render))
	js.Global().Set("name", "licong")
	<-done
}