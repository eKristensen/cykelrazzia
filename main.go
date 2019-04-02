package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

func main() {
	// Start nu
	start := time.Now()

	// Cyeklrazzia dato - min varsel 6 uger
	razzia := start.AddDate(0, 0, 7*6)

	// Platkat ned dato - må tages ned 2 uger efter razzia
	end := razzia.AddDate(0, 0, 7*2)

	fmt.Print(end.Format("Mon Jan 2 15:04:05 MST 2006"))

	// Seddel der hænges op
	note := template.Must(template.ParseFiles("note.tmpl"))

	// Output til note.tex
	out, err := os.Create("note.tex")
	if err != nil {
		panic(err)
	}
	note.Execute(out, map[string]interface{}{
		"opsat":                    "Skriv studienr (KUN TAL) i boksen",
		"nedtages":                 "Forslag/fejl mv.? Send en mail til ek@pf.dk. Denne website er sidst ændret 6. januar 2019.",
		"DateWithWeekdayAllcapsDK": "",
	})

	fmt.Print("Hello World")
}
