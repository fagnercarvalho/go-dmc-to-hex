package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"go-dmc-to-hex/dmc"
	"io/fs"
	"net/http"
	"os"
)

func main() {
	dmcs, err := scrapeDMCs()
	if err != nil {
		panic(err)
	}

	err = saveToJSON(dmcs)
	if err != nil {
		panic(err)
	}
}

func scrapeDMCs() ([]dmc.DMC, error) {
	resp, err := http.Get("https://floss.maxxmint.com/index.php")
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var dmcs []dmc.DMC
	names := scrapeColumn(doc, ".tableborder td:nth-child(2)")
	dmcCodes := scrapeColumn(doc, ".tableborder td:nth-child(3)")
	hexCodes := scrapeColumn(doc, ".tableborder td:nth-child(8)")

	for i := 0; i < len(names); i++ {
		dmc := dmc.DMC{
			ColorName: names[i],
			DMCCode:   dmcCodes[i],
			HexCode:   hexCodes[i],
		}

		dmcs = append(dmcs, dmc)
	}

	return dmcs, nil
}

func scrapeColumn(doc *goquery.Document, selector string) []string {
	var results []string

	doc.Find(selector).Each(func(i int, selection *goquery.Selection) {
		results = append(results, selection.Text())
	})

	return results
}

func saveToJSON(dmcs []dmc.DMC) error {
	bytes, err := json.MarshalIndent(dmcs, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile("dmc/dmcs.json", bytes, fs.ModePerm)
}
