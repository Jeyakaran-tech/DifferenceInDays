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

outerloop:
	for {
		fmt.Print("Enter First Date(dd/MM/YYYY): ")

		//getting the first date
		firstDate, _, _ := reader.ReadLine()
		firstDateString := string(firstDate)

		//validation for first date happens
		ok1 := computation.IsValidDate((string(firstDate)))

	Innerloop:
		if ok1 {
			fmt.Print("Enter Second Date(dd/MM/YYYY): ")
			//getting the second date
			secondDate, _, _ := reader.ReadLine()
			secondDateString := string(secondDate)

			//validation for second date happens
			ok2 := computation.IsValidDate((string(secondDate)))

			if ok2 { //If the validation passes
				//Two strings gets splitted into three seperate integer parts (date, month, year)
				integerArrayFirstDate := computation.SplitString(firstDateString)
				integerArraySecondDate := computation.SplitString(secondDateString)

				//declaring the date, month and year into struct
				date1 := types.NewDate(integerArrayFirstDate[0], integerArrayFirstDate[1], integerArrayFirstDate[2])
				date2 := types.NewDate(integerArraySecondDate[0], integerArraySecondDate[1], integerArraySecondDate[2])

				//we get the difference here
				diff := computation.NumberOfDaysBetweenTheDates(date1, date2)

				//Always the result is subtracted by 1, because we should not take the starting day into account
				if diff < 0 {
					fmt.Println((diff*-1)-1, "days") //If the result is negative, we will multiply by -1
				} else if diff == 0 {
					fmt.Println("0 days (Both dates provided are same)") //If the result is 0, then both the dates are same
				} else {
					fmt.Println(diff-1, "days")
				}
			} else {
				// This message gets printed if any of the the dates were invalid
				fmt.Println("The dates were of invalid format. Check the format once")
				goto Innerloop
			}

		} else {
			// This message gets printed if any of the the dates were invalid
			fmt.Println("The dates were of invalid format. Check the format once")
			goto outerloop
		}

		//This block helps to run the program till the user press "2"
		fmt.Println("1. Press 1 to continue\n2. Press 2 to exit")
		fmt.Scan(&option)
		switch option {
		case 1:
			goto outerloop
		default:
			break outerloop
		}

	}
}
