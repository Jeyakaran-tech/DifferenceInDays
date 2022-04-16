package computation

import (
	"regexp"
	"time"

	"github.com/Jeyakaran-tech/DifferenceInDays/types"
)

//NumberOfDaysBetweenTheDates - will calculate the days between two passed dates. The starting and ending date will be omitted
func NumberOfDaysBetweenTheDates(d1 *types.Date, d2 *types.Date) int {

	//Counting all the days within the year. Days of current month also been taken into account
	n1 := d1.Year*365 + d1.Day

	//The days of all the previous months have been added here 2/3/2000
	for i := 0; i < d1.Month-1; i++ {
		n1 += monthDays[i]
	}

	//Takes all the leap years within the year into account
	n1 += countLeapYears(d1)

	//Counting all the days within the year. Days of current month also been taken into account
	n2 := d2.Year*365 + d2.Day

	//The days of all the previous months have been added here
	for i := 0; i < d2.Month-1; i++ {
		n2 += monthDays[i]
	}

	//Takes all the leap years within the year into account
	n2 += countLeapYears(d2)

	//difference between the total days will give us the difference
	return (n2 - n1)
}

func countLeapYears(date *types.Date) int {
	years := date.Year
	if date.Month <= 2 {
		years--
	}
	return years/4 - years/100 + years/400
}

//IsValidDate - Checks if the date is valid based on
//1. Year range between 1900 - 2999
//2. If the format is dd/mm/YYYY (if it is dd-mm-yyyy or dd.mm.yyyy, the function will return false)
func IsValidDate(date string) bool {

	re := regexp.MustCompile(`(((0[1-9]|[12][0-9]|3[01])([\/])(0[13578]|10|12)([\/])(?:(?:19|2[0-9])[0-9]{2}))|(([0][1-9]|[12][0-9]|30)([\/])(0[469]|11)([\/])(?:(?:19|2[0-9])[0-9]{2}))|((0[1-9]|1[0-9]|2[0-8])([\/])(02)([\/])(?:(?:19|2[0-9])[0-9]{2})))`)
	return re.MatchString(date) || checkForLeapYear(date)
}

func findDifferenceUsingTimePackage(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func checkForLeapYear(date string) bool {

	leapYearFlag := false
	dates := SplitString(date)
	day, year := dates[0], dates[2]

	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				leapYearFlag = true
			} else {
				leapYearFlag = false
			}
		} else {
			leapYearFlag = true
		}
	} else {
		leapYearFlag = false
	}

	return day == 29 && leapYearFlag && year >= 1900 && year <= 2999

}
