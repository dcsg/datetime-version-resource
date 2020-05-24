package resource_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDatetimeVersionResource(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Datetime Version Resource Suite")
}
