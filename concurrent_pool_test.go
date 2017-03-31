package spawning_test

import (
	. "github.com/ess/spawning"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ConcurrentPool", func() {
	var pool Pool

	BeforeEach(func() {
		pool = NewConcurrentPool()
	})

	Describe("Add", func() {
		It("returns the pool itself", func() {
			Expect(pool.Add("something")).To(Equal(pool))
		})
	})

	Describe("Run", func() {
		It("runs the commands concurrently", func() {
			results := pool.
				Add("sleep 0.4").
				Add("sleep 0.3").
				Add("sleep 0.2").
				Add("sleep 0.1").
				Run()

			Expect(results[0].Command).To(Equal("sleep 0.1"))
			Expect(results[1].Command).To(Equal("sleep 0.2"))
			Expect(results[2].Command).To(Equal("sleep 0.3"))
			Expect(results[3].Command).To(Equal("sleep 0.4"))
		})

		Context("when a command is not successful", func() {
			It("reflects the failure in that command's result", func() {
				result := pool.Add("false").Run()[0]

				Expect(result.Success).To(Equal(false))
			})
		})
	})

})
