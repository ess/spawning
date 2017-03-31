package spawning_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSpawning(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spawning Suite")
}
