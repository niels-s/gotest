package kabisa_test

import (
	. "kabisa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
		var (
			longBook Book
		)

		BeforeEach(func() {
			longBook = Book{"Fabeltjes krant", "Rui", 300}
		})

		Describe("Books", func() {
				Context("With me", func() {
						It("should verify something", func() {
								Expect(1).To(Equal(1))
							})
					})
			})

	})
