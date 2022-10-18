package queue_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "go-queue/pkg"
)

var _ = Describe("Pkg", func() {
	var queue Queue

	BeforeEach(func() {
		queue = Queue{}
	})

	Describe("initialization tests", func() {
		When("queue is initialized", func() {
			It("has size 0", func() {
				Expect(queue.Size).To(Equal(1))
			})
		})
	})
})
