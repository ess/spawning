package mockable_test

import (
	"os"

	. "github.com/ess/mockable"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const mockableEnvVar = "MOCKABLE"

var _ = Describe("Mockable", func() {
	Describe("Mocked", func() {
		var result bool

		Context("when MOCKABLE is not set", func() {
			BeforeEach(func() {
				os.Unsetenv(mockableEnvVar)

				result = Mocked()
			})

			It("is false", func() {
				Expect(result).To(Equal(false))
			})
		})

		Context("when MOCKABLE is not set", func() {
			BeforeEach(func() {
				os.Setenv(mockableEnvVar, "anything")

				result = Mocked()
			})

			It("is true", func() {
				Expect(result).To(Equal(true))
			})
		})
	})
})
