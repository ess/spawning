package spawning_test

import (
	. "github.com/ess/spawning"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	mockable "github.com/ess/mockable"
)

var _ = Describe("MockedPool", func() {
	var pool Pool

	BeforeEach(func() {
		mockable.Enable()
		pool = NewConcurrentPool()
	})

	AfterEach(func() {
		mockable.Disable()
	})

	Describe("Add", func() {
		It("returns the pool itself", func() {
			Expect(pool.Add("something")).To(Equal(pool))
		})
	})

	Describe("Run", func() {
		It("runs the commands in their addition order", func() {
			results := pool.
				Add("sleep 0.4").
				Add("sleep 0.3").
				Add("sleep 0.2").
				Add("sleep 0.1").
				Run()

			Expect(results[0].Command).To(Equal("sleep 0.4"))
			Expect(results[1].Command).To(Equal("sleep 0.3"))
			Expect(results[2].Command).To(Equal("sleep 0.2"))
			Expect(results[3].Command).To(Equal("sleep 0.1"))
		})

		Context("when a command is not successful", func() {
			It("ignores the failure in that command's result", func() {
				result := pool.Add("false").Run()[0]
				Expect(result.Success).To(Equal(true))
			})
		})
	})

})
