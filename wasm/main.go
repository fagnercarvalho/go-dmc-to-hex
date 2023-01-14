package main

import (
	"syscall/js"

	"go-dmc-to-hex/dmc"
)

func main() {
	js.Global().Set("getNameByDMC", getNameWrapper())
	js.Global().Set("getHexByDMC", getHexWrapper())
	<-make(chan bool)
}

func getNameWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 {
			return "must provide hex code to convert"
		}

		argument := args[0]

		hex, err := getNameByDMC(argument.String())
		if err != nil {
			return "error when trying to get hex by DMC"
		}

		return hex
	})
}

func getHexWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 {
			return "must provide hex code to convert"
		}

		argument := args[0]

		hex, err := getHexByDMC(argument.String())
		if err != nil {
			return "error when trying to get name by DMC"
		}

		return hex
	})
}

func getNameByDMC(dmcCode string) (string, error) {
	converter, err := dmc.NewDMCConverter()
	if err != nil {
		return "", err
	}

	return converter.GetColorByDMC(dmcCode), nil
}

func getHexByDMC(dmcCode string) (string, error) {
	converter, err := dmc.NewDMCConverter()
	if err != nil {
		return "", err
	}

	return converter.GetHexByDMC(dmcCode), nil
}
