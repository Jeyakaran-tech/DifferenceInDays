package types

//Date holds all the three members. (i.e.,) Day, Month and Year
type Date struct {
	Day   int
	Month int
	Year  int
}

//NewDate creates New Date object
func NewDate(day, month, year int) *Date {
	return &Date{
		Day:   day,
		Month: month,
		Year:  year,
	}

}
