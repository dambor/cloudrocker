package compressor_test

import (
	. "github.com/cloudcredo/cloudrocker/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/cloudcredo/cloudrocker/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestCompressor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Compressor Suite")
}
