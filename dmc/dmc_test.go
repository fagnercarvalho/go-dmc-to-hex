package dmc

import (
	"fmt"
	"testing"
)

func TestConverter_GetHexByDMC(t *testing.T) {
	converter, err := NewDMCConverter()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		dmc         string
		expectedHex string
	}{
		{
			dmc:         "991",
			expectedHex: "#477B6E",
		},
		{
			dmc:         "970",
			expectedHex: "#F78B13",
		},
		{
			dmc:         "3865",
			expectedHex: "#F9F7F1",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v to %v", test.dmc, test.expectedHex), func(t *testing.T) {
			hex := converter.GetHexByDMC(test.dmc)
			if hex != test.expectedHex {
				t.Fatal()
			}
		})
	}
}

func TestConverter_GetDMCByHex(t *testing.T) {
	converter, err := NewDMCConverter()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		hex         string
		expectedDMC string
	}{
		{
			hex:         "#9C599C",
			expectedDMC: "33",
		},
		{
			hex:         "#CD2F63",
			expectedDMC: "600",
		},
		{
			hex:         "#AB3357",
			expectedDMC: "3803",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v to %v", test.hex, test.expectedDMC), func(t *testing.T) {
			dmc := converter.GetDMCByHex(test.hex)
			if dmc != test.expectedDMC {
				t.Fatal()
			}
		})
	}
}
