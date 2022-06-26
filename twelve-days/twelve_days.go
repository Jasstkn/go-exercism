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
	{"first", "one", "a Partridge in a Pear Tree"},
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

func getGifts(i int, str string) string {
	if i == 1 && str == "" {
		return " " + gifts[0].giftName + "."
	}
	if i == 1 && str != "" {
		return str + " and " + gifts[0].giftName + "."
	}
	return getGifts(i-1, str+" "+gifts[i-1].quantity+" "+gifts[i-1].giftName+",")
}

func Verse(i int) string {
	return "On the " + gifts[i-1].day + " day of Christmas my true love gave to me:" + getGifts(i, "")
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
