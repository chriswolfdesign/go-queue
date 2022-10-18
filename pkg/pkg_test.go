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
				Expect(queue.Size).To(Equal(0))
			})
		})
	})

	Describe("enqueue tests", func() {
		When("enqueuing to empty queue", func() {
			BeforeEach(func() {
				queue.Enqueue(1)
			})

			It("has size one", func() {
				Expect(queue.Size).To(Equal(1))
			})
		})

		When("enqueuing to list with one element", func() {
			BeforeEach(func() {
				queue.Enqueue(1)
				queue.Enqueue(2)
			})

			It("Has size two", func() {
				Expect(queue.Size).To(Equal(2))
			})
		})
	})
})
