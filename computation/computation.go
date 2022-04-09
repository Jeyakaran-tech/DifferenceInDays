package computation

import (
	"regexp"
	"strconv"

	"github.com/Jeyakaran-tech/DifferenceInDays/types"
)

func NumberOfDaysBetweenTheDates(d1 *types.Date, d2 *types.Date) int {
	n1 := d1.Year*365 + d1.Day
	for i := 0; i < d1.Month-1; i++ {
		n1 += monthDays[i]
	}
	n1 += countLeapYears(d1)

	n2 := d2.Year*365 + d2.Day
	for i := 0; i < d2.Month-1; i++ {
		n2 += monthDays[i]
	}
	n2 += countLeapYears(d2)
	return (n2 - n1)

}

func countLeapYears(date *types.Date) int {
	years := date.Year
	if date.Month <= 2 {
		years--
	}
	return years/4 - years/100 + years/400
}

func IsValidDate(date string) bool {

	yearString := date[len(date)-4:]

	year, err := strconv.Atoi(yearString)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`^[0-3]?[0-9]\/[0-3]?[0-9]\/(?:[0-9]{4})$`)
	if re.MatchString(date) && year > 1900 && year < 2999 {
		return true
	}
	return false

}
