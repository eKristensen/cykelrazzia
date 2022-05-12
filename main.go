package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"
)

func main() {
	// Start nu
	start := time.Now()

	// Cyeklrazzia dato (fjerne cykler dag) - min varsel 4 uger
	// Format DD/MM-ÅÅÅÅ
	raid, err := time.Parse("2/1-06", "9/6-22")
	if err != nil {
		panic(err)
	}

	// Tjek at de 4 ugers varsel er overholdt! Dispensation er givet til juni 2022 razzia
	// Husk at tjekke om husorden er blevet rettet til 4 uger næste gang!
	// minus 1 dag, da det er okay at varsle præcis x antal uger før (altså på dagen!)
	if raid.Before(start.AddDate(0, 0, 7*4-1)) {
		fmt.Print("Der skal være mindst 4 uger imellem razzia og varsel!\n")
	}

	// Platkat ned dato - må tages ned 2 uger efter razzia
	end := raid.AddDate(0, 0, 7*2)

	// Referance dato: 20060102150405
	// Mon Jan 2 15:04:05 -0700 MST 2006

	//fmt.Print(end.Format("Mon Jan 2 15:04:05 MST 2006"))

	// Seddel der hænges op
	note := template.Must(template.ParseFiles("note.tmpl"))
	mail := template.Must(template.ParseFiles("mail.tmpl"))

	// Output til note.tex
	out, err := os.Create("note.tex")
	if err != nil {
		panic(err)
	}
	note.Execute(out, map[string]interface{}{
		"opsat":                    start.Format("2/1-06"),
		"nedtages":                 end.Format("2/1-06"),
		"DateWithWeekdayAllcapsDK": strings.ToUpper(ugedag(raid) + " D." + raid.Format(" 2. ") + maaned(raid) + raid.Format(" 2006")), // eksempel: TORSDAG D. 24. MAJ 2018
		"DateWithWeekdayAllcapsEN": strings.ToUpper(raid.Format("Monday January 2, 2006")),                                            // eksempel: TUESDAY MAY 24, 2018
		"raidOnDK":                 raid.Format("2. ") + maaned(raid),                                                                 // eksempel: 24. maj
		"raidOnEN":                 raid.Format("January 2"),
	})

	out, err = os.Create("mail.tex")
	if err != nil {
		panic(err)
	}
	mail.Execute(out, map[string]interface{}{
		"DateWithWeekdayDK": ugedag(raid) + " d." + raid.Format(" 2. ") + maaned(raid) + raid.Format(" 2006"), // eksempel: torsdag d. 24. maj 2018.
		"DateWithWeekdayEN": raid.Format("Monday January 2, 2006"),
		"SubjectDK":         raid.Format("2. ") + maaned(raid),                            // eksempel: 24. maj
		"SubjectEN":         raid.Format("January 2"),                                     // eksempel: May 24th
		"NoticeDK":          ugedag(start) + " d." + start.Format(" 2. ") + maaned(start), // Eksempel: torsdag d. 12. april
		"NoticeEN":          start.Format("Monday January 2"),                             // Eksempel: Thursday April 12
	})
}

func ugedag(date time.Time) string {
	weekday := date.Weekday()
	ugedag := int(weekday)
	switch ugedag {
	case 1:
		return "mandag"
	case 2:
		return "tirsdag"
	case 3:
		return "onsdag"
	case 4:
		return "torsdag"
	case 5:
		return "fredag"
	case 6:
		return "lørdag"
	case 7:
		return "søndag"
	}
	return ""
}

func maaned(date time.Time) string {
	month := date.Month()
	maaned := int(month)
	switch maaned {
	case 1:
		return "januar"
	case 2:
		return "februar"
	case 3:
		return "marts"
	case 4:
		return "april"
	case 5:
		return "maj"
	case 6:
		return "juni"
	case 7:
		return "juli"
	case 8:
		return "august"
	case 9:
		return "september"
	case 10:
		return "oktober"
	case 11:
		return "november"
	case 12:
		return "december"
	}
	return ""
}
