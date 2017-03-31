package mockable_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMockable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mockable Suite")
}
