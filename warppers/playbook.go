package warppers

import (
	"hellowasm/playbook"
	"syscall/js"
)

func ToPlaybookWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid arguments"
		}
		inputJSON := args[0].String()
		pretty, err := playbook.ToPlaybook(inputJSON)
		if err != nil {
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}

func ToPlaybookFromURLWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid arguments"
		}
		inputJSON := args[0].String()
		pretty, err := playbook.ToPlaybookFromURL(inputJSON)
		if err != nil {
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}
