package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Jeyakaran-tech/DifferenceInDays/computation"
	"github.com/Jeyakaran-tech/DifferenceInDays/types"
)

var option int

func main() {
	reader := bufio.NewReader(os.Stdin)

loop:
	for {
		fmt.Print("Enter First Date(dd/MM/YYYY): ")
		firstDate, _, _ := reader.ReadLine()
		firstDateString := string(firstDate)
		ok1 := computation.IsValidDate((string(firstDate)))

		fmt.Print("Enter Second Date(dd/MM/YYYY): ")
		secondDate, _, _ := reader.ReadLine()
		secondDateString := string(secondDate)
		ok2 := computation.IsValidDate((string(secondDate)))

		if ok1 && ok2 {
			integerArrayFirstDate := computation.SplitString(firstDateString, "/")
			integerArraySecondDate := computation.SplitString(secondDateString, "/")

			AfterDate := types.NewDate(integerArraySecondDate[0], integerArraySecondDate[1], integerArraySecondDate[2])
			beforeDate := types.NewDate(integerArrayFirstDate[0], integerArrayFirstDate[1], integerArrayFirstDate[2])

			diff := computation.NumberOfDaysBetweenTheDates(beforeDate, AfterDate)

			if diff < 0 {
				fmt.Println((diff*-1)-1, "days")
			} else {
				fmt.Println(diff-1, "days")
			}
		} else {
			fmt.Println("The dates were of invalid format. Check the format once")
			goto loop
		}

		fmt.Println("1. Press 1 to continue\n2. Press 2 to exit")
		fmt.Scan(&option)
		switch option {
		case 1:
			goto loop
		default:
			break loop
		}

	}
}
