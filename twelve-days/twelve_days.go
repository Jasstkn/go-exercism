package twelve

import (
	"fmt"
	"strings"
)

var days = [...]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var count = [...]string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
}

var gifts = [...]string{
	"a Partridge in a Pear Tree",
	"Turtle Doves",
	"French Hens",
	"Calling Birds",
	"Gold Rings",
	"Geese-a-Laying",
	"Swans-a-Swimming",
	"Maids-a-Milking",
	"Ladies Dancing",
	"Lords-a-Leaping",
	"Pipers Piping",
	"Drummers Drumming",
}

func getGifts(i int, str string) string {
	if i == 1 && str == "" {
		return fmt.Sprintf(" %s.", gifts[0])
	}
	if i == 1 && str != "" {
		return fmt.Sprintf("%s and %s.", str, gifts[0])
	}
	return getGifts(i-1, fmt.Sprintf("%s %s %s,", str, count[i-1], gifts[i-1]))
}

func Verse(i int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me:%s", days[i-1], getGifts(i, ""))
}

func Song() string {
	var res strings.Builder
	for i := 1; i <= len(count); i++ {
		if i != len(count) {
			res.WriteString(Verse(i) + "\n")
		} else {
			res.WriteString(Verse(i))
		}
	}
	return res.String()
}
