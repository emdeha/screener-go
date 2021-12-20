package company_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCompany(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Company Suite")
}
