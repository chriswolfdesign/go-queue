package go_queue_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoQueue Suite")
}
