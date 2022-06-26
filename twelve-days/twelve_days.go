package twelve

import (
	"strings"
)

type gift struct {
	day      string
	quantity string
	giftName string
}

var gifts = []gift{
	{"first", "", "a Partridge in a Pear Tree"},
	{"second", "two", "Turtle Doves"},
	{"third", "three", "French Hens"},
	{"fourth", "four", "Calling Birds"},
	{"fifth", "five", "Gold Rings"},
	{"sixth", "six", "Geese-a-Laying"},
	{"seventh", "seven", "Swans-a-Swimming"},
	{"eighth", "eight", "Maids-a-Milking"},
	{"ninth", "nine", "Ladies Dancing"},
	{"tenth", "ten", "Lords-a-Leaping"},
	{"eleventh", "eleven", "Pipers Piping"},
	{"twelfth", "twelve", "Drummers Drumming"},
}

func getGifts(line int) (res string) {
	line -= 1

	// if it's first line
	if line == 0 {
		return " " + gifts[0].giftName + "."
	}

	for i := line; i >= 0; i-- {
		if i == 0 {
			res += " and " + gifts[0].giftName + "."
		} else {
			res += " " + gifts[i].quantity + " " + gifts[i].giftName + ","
		}
	}
	return res
}

func Verse(i int) string {
	return "On the " + gifts[i-1].day + " day of Christmas my true love gave to me:" + getGifts(i)
}

func Song() string {
	verses := make([]string, len(gifts))

	n := len("\n") * (len(gifts) - 1)
	for i := 1; i <= len(gifts); i++ {
		verse := Verse(i)
		verses[i-1] = verse
		n += len(verse)
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(verses[0])

	for i := 1; i <= len(gifts)-1; i++ {
		b.WriteString("\n")
		b.WriteString(verses[i])
	}
	return b.String()
}
