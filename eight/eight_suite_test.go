package eight_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEight(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Eight Suite")
}
