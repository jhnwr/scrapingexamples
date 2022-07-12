package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)

type Result struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func main() {

	output := []Result{}

	c := colly.NewCollector()

	d := c.Clone()

	c.OnHTML("span.COLviolaChiaro", func(h *colly.HTMLElement) {
		if strings.HasPrefix(h.Text, "F") {
			d.Visit("https://www.federbridge.it/regioni/dettAss.asp?codice=" + strings.TrimSpace(h.Text) + "&sigla=TOS")
		}

	})

	d.OnHTML("body", func(h *colly.HTMLElement) {
		email := h.ChildText("td.ALLbaseR a[href^='mailto']")
		name := h.ChildText("td.FNTbase12")
		output = append(output, Result{Email: email, Name: name})

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	d.OnRequest(func(r *colly.Request) {
		fmt.Println("d visiting", r.URL.String())
	})

	c.Visit("https://federbridge.it/SocSp/ChkFS.asp?FS=F")

	json_data, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("output.json", json_data, 0644)
	if err != nil {
		fmt.Println(err)
	}

}
