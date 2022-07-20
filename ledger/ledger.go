package ledger

import (
	"errors"
	"math"
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

func generateDescription(input string) string {
	if len(input) > descriptionWidth {
		return input[:22] + "..."
	}
	return input + strings.Repeat(" ", descriptionWidth-len(input))
}

func isNegative(input int) bool {
	return input < 0
}

func generateCurrency(currency string) (string, error) {
	var out string
	switch currency {
	case "EUR":
		out = "€"
	case "USD":
		out = "$"
	default:
		return "", errors.New("")
	}
	return out, nil
}

func generateChangeLeadingZeros(input string) string {
	switch len(input) {
	case 1:
		return "00" + input
	case 2:
		return "0" + input
	default:
		return input
	}
}

func generateChangeValue(input string, separator string) (out string) {
	var parts []string

	for len(input) > 3 {
		parts = append(parts, input[len(input)-3:])
		input = input[:len(input)-3]
	}

	if len(input) > 0 {
		parts = append(parts, input)
	}
	for i := len(parts) - 1; i >= 0; i-- {
		out += parts[i] + separator
	}

	return out
}

func generateChange(locale, currency string, change int, isNegative bool) (string, error) {
	var out string

	changeStr := strconv.Itoa(change)

	switch locale {
	case "nl-NL":
		outCurrency, err := generateCurrency(currency)
		if err != nil {
			return "", err
		}
		out += outCurrency + " "

		fullChange := generateChangeLeadingZeros(changeStr)
		intChange := fullChange[:len(fullChange)-2]

		out += generateChangeValue(intChange, ".")

		out = out[:len(out)-1] + "," + fullChange[len(fullChange)-2:]

		if isNegative {
			out += "-"
		} else {
			out += " "
		}
	case "en-US":

		outCurrency, err := generateCurrency(currency)
		if err != nil {
			return "", err
		}
		out += outCurrency

		fullChange := generateChangeLeadingZeros(changeStr)
		decimal := fullChange[len(fullChange)-2:]
		// remove last 2 digits
		intChange := fullChange[:len(fullChange)-2]

		out += generateChangeValue(intChange, ",")

		out = out[:len(out)-1] + "." + decimal

		if isNegative {
			out = "(" + out + ")"
		} else {
			out = out + " "
		}
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
			date, err := parseDate(entry.Date, locale)
			if err != nil {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}

			description := generateDescription(entry.Description)

			changeAbs := math.Abs(float64(entry.Change))
			change, err := generateChange(locale, currency, int(changeAbs), isNegative(entry.Change))
			if err != nil {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}

			var al int
			for range change {
				al++
			}
			co <- struct {
				i int
				s string
				e error
			}{i: i, s: date + strings.Repeat(" ", 10-len(date)) + " | " + description + " | " +
				strings.Repeat(" ", 13-al) + change + "\n"}
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
