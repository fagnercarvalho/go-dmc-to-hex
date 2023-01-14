package dmc

import (
	_ "embed"
	"encoding/json"
)

type DMC struct {
	ColorName string
	DMCCode   string
	HexCode   string
}

//go:embed dmcs.json
var dmcFile []byte

type Converter struct {
	dmcs []DMC
}

func NewDMCConverter() (Converter, error) {
	dmcs, err := parseFile()
	if err != nil {
		return Converter{}, err
	}

	return Converter{dmcs: dmcs}, nil
}

func (c Converter) GetHexByDMC(dmcCode string) string {
	for _, dmc := range c.dmcs {
		if dmc.DMCCode == dmcCode {
			return dmc.HexCode
		}
	}

	return ""
}

func (c Converter) GetDMCByHex(hexCode string) string {
	for _, dmc := range c.dmcs {
		if dmc.HexCode == hexCode {
			return dmc.DMCCode
		}
	}

	return ""
}

func (c Converter) GetColorByDMC(dmcCode string) string {
	for _, dmc := range c.dmcs {
		if dmc.DMCCode == dmcCode {
			return dmc.ColorName
		}
	}

	return ""
}

func parseFile() ([]DMC, error) {
	var dmcs []DMC

	err := json.Unmarshal(dmcFile, &dmcs)
	if err != nil {
		return nil, err
	}

	return dmcs, nil
}
