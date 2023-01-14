package main

import (
	"syscall/js"

	"go-dmc-to-hex/dmc"
)

func main() {
	js.Global().Set("getHexByDMC", getWrapper())
	<-make(chan bool)
}

func getWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 {
			return "must provide hex code to convert"
		}

		argument := args[0]

		hex, err := getHexByDMC(argument.String())
		if err != nil {
			return "error when trying to get hex by DMC"
		}

		return hex
	})
}

func getHexByDMC(dmcCode string) (string, error) {
	converter, err := dmc.NewDMCConverter()
	if err != nil {
		return "", err
	}

	return converter.GetHexByDMC(dmcCode), nil
}
