package kabisa_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKabisa(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kabisa Suite")
}
