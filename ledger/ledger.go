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
	changeWidth      = 13

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
	separator := "..."
	if len(input) > descriptionWidth {
		return input[:descriptionWidth-len(separator)] + separator
	}
	return input + strings.Repeat(" ", descriptionWidth-len(input))
}

func isNegative(input int) bool {
	return input < 0
}

func getCurrency(currency string) (string, error) {
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

func formatIntegerChange(input string, separator string) (out string) {
	var parts []string

	if len(input) > 3 {
		parts = append(parts, input[len(input)-3:])
		input = input[:len(input)-3]
	}

	if len(input) > 0 {
		parts = append(parts, input)
	}

	// reverse array
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	return strings.Join(parts, separator)
}

func generateChange(locale, currency string, change int, isNegative bool) (string, error) {
	var out string

	changeStr := strconv.Itoa(change)

	switch locale {
	case "nl-NL":
		outCurrency, err := getCurrency(currency)
		if err != nil {
			return "", err
		}
		out += outCurrency + " "

		fullChange := generateChangeLeadingZeros(changeStr)
		intChange, decimal := fullChange[:len(fullChange)-2], fullChange[len(fullChange)-2:]

		// concatenate integer and decimal parts
		out += formatIntegerChange(intChange, ".") + "," + decimal

		if isNegative {
			out += "-"
		} else {
			out += " "
		}
	case "en-US":
		outCurrency, err := getCurrency(currency)
		if err != nil {
			return "", err
		}
		out += outCurrency

		fullChange := generateChangeLeadingZeros(changeStr)
		// get integer and decimal parts of the number
		intChange, decimal := fullChange[:len(fullChange)-2], fullChange[len(fullChange)-2:]

		// concatenate integer and decimal parts
		out += formatIntegerChange(intChange, ",") + "." + decimal

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

func currencyIsValid(currency string) bool {
	switch currency {
	case
		"USD",
		"EUR":
		return true
	default:
		return false
	}
}

func FormatLedger(currency string, locale string, inputEntries []Entry) (string, error) {
	if !currencyIsValid(currency) {
		return "", errors.New("")
	}

	entries := make([]Entry, len(inputEntries))
	copy(entries, inputEntries)

	// sort entries
	sort.Slice(entries[:], func(i, j int) bool {
		return entries[i].Date < entries[j].Date || entries[i].Description < entries[j].Description || entries[i].Change < entries[j].Change
	})

	// generate heading
	heading, err := generateHeading(locale)
	if err != nil {
		return "", err
	}

	// Parallelism, always a great idea
	ch := make(chan struct {
		i int
		s string
		e error
	})

	for i, entry := range entries {
		go func(i int, entry Entry) {
			date, err := parseDate(entry.Date, locale)
			if err != nil {
				ch <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}

			description := generateDescription(entry.Description)

			changeAbs := math.Abs(float64(entry.Change))
			change, err := generateChange(locale, currency, int(changeAbs), isNegative(entry.Change))
			if err != nil {
				ch <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
			}

			ch <- struct {
				i int
				s string
				e error
			}{i: i, s: date + strings.Repeat(" ", dateWidth-len(date)) + " | " + description + " | " +
				strings.Repeat(" ", changeWidth-len([]rune(change))) + change + "\n"}
		}(i, entry)
	}
	body := make([]string, len(entries))
	for range entries {
		entry := <-ch
		if entry.e != nil {
			return "", entry.e
		}
		body[entry.i] = entry.s
	}

	return heading + strings.Join(body, ""), nil
}
