package computation

import (
	ginkgo "github.com/onsi/ginkgo/v2"
	gomega "github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Computation", func() {
	ginkgo.When("Dates were passed", func() {
		ginkgo.It("should return two seperate arrays with date, month and year", func() {
			date1 := "3/1/2000"
			seperator := "/"
			dates := SplitString(date1, seperator)
			gomega.Expect(dates[0]).To(gomega.BeEquivalentTo(3))
		})
	})

})
