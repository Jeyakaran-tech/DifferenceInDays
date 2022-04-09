package computation

import (
	"github.com/Jeyakaran-tech/DifferenceInDays/types"

	ginkgo "github.com/onsi/ginkgo/v2"
	gomega "github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Computation", func() {
	ginkgo.When("Dates were passed", func() {
		ginkgo.It("should tell us the number of days between the dates", func() {
			date1 := types.NewDate(3, 1, 2000)
			date2 := types.NewDate(5, 1, 2000)
			differenceInDays := NumberOfDaysBetweenTheDates(date1, date2)
			gomega.Expect(differenceInDays).To(gomega.BeEquivalentTo(2))
		})
	})

	ginkgo.When("Date were passed to count the numberOfLeapYears", func() {
		ginkgo.It("should return the number of leap years within the year", func() {
			date1 := types.NewDate(3, 1, 2000)
			numberOfLeapYears := countLeapYears(date1)
			gomega.Expect(numberOfLeapYears).To(gomega.BeEquivalentTo(484))
		})
	})

})
