package edgarimporter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEDGARImporter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EDGAR Importer Suite")
}
