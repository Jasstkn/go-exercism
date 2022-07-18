package ledger

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

const (
	descriptionWidth = 25
	dateWidth        = 10

	layout = "2006-01-02"
)

func generateHeading(locale string) (s string, err error) {
	switch locale {
	case "nl-NL":
		s = "Datum" +
			strings.Repeat(" ", dateWidth-len("Datum")) +
			" | " +
			"Omschrijving" +
			strings.Repeat(" ", descriptionWidth-len("Omschrijving")) +
			" | " + "Verandering" + "\n"
	case "en-US":
		s = "Date" +
			strings.Repeat(" ", dateWidth-len("Date")) +
			" | " +
			"Description" +
			strings.Repeat(" ", descriptionWidth-len("Description")) +
			" | " + "Change" + "\n"
	default:
		return "", errors.New("")
	}
	return s, err
}

func parseDate(date string, locale string) (string, error) {
	d, err := time.Parse(layout, date)
	if err != nil {
		return "", errors.New("")
	}

	var out string
	switch locale {
	case "nl-NL":
		out = d.Format("02-01-2006")
	case "en-US":
		out = d.Format("01/02/2006")
	default:
		return "", errors.New("")
	}
	return out, nil
}

func FormatLedger(currency string, locale string, inputEntries []Entry) (string, error) {
	if len(inputEntries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	entries := make([]Entry, len(inputEntries))
	copy(entries, inputEntries)

	// sort entries
	sort.Slice(entries[:], func(i, j int) bool {
		return entries[i].Date < entries[j].Date || entries[i].Description < entries[j].Description || entries[i].Change < entries[j].Change
	})

	// generate heading
	s, err := generateHeading(locale)
	if err != nil {
		return "", err
	}

	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})

	for i, et := range entries {
		go func(i int, entry Entry) {
			d, err := parseDate(entry.Date, locale)
			if err != nil {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}

			de := entry.Description
			if len(de) > descriptionWidth {
				de = de[:22] + "..."
			} else {
				de += strings.Repeat(" ", descriptionWidth-len(de))
			}

			negative := false
			cents := entry.Change
			if cents < 0 {
				cents *= -1
				negative = true
			}

			var a string
			if locale == "nl-NL" {
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
				}
				a += " "
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + "."
				}
				a = a[:len(a)-1]
				a += ","
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += "-"
				} else {
					a += " "
				}
			} else if locale == "en-US" {
				if negative {
					a += "("
				}
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
				}
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + ","
				}
				a = a[:len(a)-1]
				a += "."
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += ")"
				} else {
					a += " "
				}
			} else {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}
			var al int
			for range a {
				al++
			}
			co <- struct {
				i int
				s string
				e error
			}{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
				strings.Repeat(" ", 13-al) + a + "\n"}
		}(i, et)
	}
	ss := make([]string, len(entries))
	for range entries {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}
	for i := 0; i < len(entries); i++ {
		s += ss[i]
	}
	return s, nil
}
