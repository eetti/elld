package configdir_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfigdir(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Configdir Suite")
}
