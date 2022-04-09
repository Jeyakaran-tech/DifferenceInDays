package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Jeyakaran-tech/DifferenceInDays/computation"
	"github.com/Jeyakaran-tech/DifferenceInDays/types"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter First Date: (dd/MM/YYYY)")
	firstDate, _, _ := reader.ReadLine()
	integerArrayFirstDate := computation.SplitString(string(firstDate), "/")

	fmt.Print("Enter Second Date: (dd/MM/YYYY)")
	secondDate, _, _ := reader.ReadLine()
	integerArraySecondDate := computation.SplitString(string(secondDate), "/")

	beforeDate := types.NewDate(integerArrayFirstDate[0], integerArrayFirstDate[1], integerArrayFirstDate[2])
	AfterDate := types.NewDate(integerArraySecondDate[0], integerArraySecondDate[1], integerArraySecondDate[2])

	diff := computation.NumberOfDaysBetweenTheDates(beforeDate, AfterDate)

	if diff < 0 {
		fmt.Println((diff*-1)-1, "days")
	} else {
		fmt.Println(diff-1, "days")
	}

}
