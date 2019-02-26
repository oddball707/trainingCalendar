package model

//A Schedule is an array of Weeks
type Schedule []*Week

func (s Schedule) Print() {
	fmt.Println("Week |  StartDate | Monday|  Tuesday  |Wednesday| Thursday | Friday | Saturday | Sunday")
	for i, week := range s {
		fmt.Print(i)
		fmt.Print("    | ")
		fmt.Print(week.weekStart.Format(dateLayout))
		fmt.Print(" | ")
		for _, day := range week.days {
			fmt.Print(day.description)
			fmt.Print("  | ")
		}
		fmt.Println(" ")
	}
}

func (s Schedule) WriteCsv() {
	file, err := os.OpenFile("calendar.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	checkError("Failed to open calendar.csv", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()
    writer.Write(header)
	for _, week := range s {
		line := make([]string, 8)
		line[0] = week.weekStart.Format(dateLayout)
		for i, day := range week.days {
			line[i+1] = day.description
		}
		fmt.Println(line)
		err := writer.Write(line)
		checkError("Failed to write line", err)
	}
}

// EmitICal implements the interface for goics
func (s Schedule) EmitICal() goics.Componenter {
	component := goics.NewComponent()
	component.SetType("VCALENDAR")
	component.AddProperty("CALSCAL", "GREGORIAN")

	for _, week := range s {
		for _, day := range week.days {
			if day.date.Before(time.Now()) { continue }
			c := goics.NewComponent()
			c.SetType("VEVENT")
			k, v := goics.FormatDateField("DTEND", day.date)
			c.AddProperty(k, v)
			k, v = goics.FormatDateField("DTSTART", day.date)
			c.AddProperty(k, v)
			c.AddProperty("SUMMARY", day.description)

			component.AddComponent(c)
		}
	}
	return component
}