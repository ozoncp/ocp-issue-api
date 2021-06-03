package utils_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"testing"
)

func TestUtils(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Utils Suite")
}
