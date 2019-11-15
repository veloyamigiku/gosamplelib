package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/veloyamigiku/gosamplelib"
)

var _ = Describe("Test", func() {

	Context("test1", func() {
		It("test1_1", func() {
			
			p := Person { Name: "sample", Age: 10 }
			//Expect(person.)
			Expect(p.Say2()).To(Equal("My name is sample, I'm 10"))
		})
	})

})
