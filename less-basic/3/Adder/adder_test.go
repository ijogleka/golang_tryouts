package adder_test

import (
	"."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("adder", func() {
	It("should add 2 and 3 and return 5", func() {
		Expect(adder.Add(2, 3)).To(Equal(5))
	})
	It("should add 2 and 3 and return 6", func() {
		Expect(adder.Add(2, 3)).To(Equal(6))
	})
	PIt("should add 2 and 3 and return 7", func() {
		Expect(adder.Add(2, 3)).To(Equal(7))
	})
})
