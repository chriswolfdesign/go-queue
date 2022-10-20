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
				Expect(queue.Size()).To(Equal(0))
			})
		})
	})

	Describe("enqueue tests", func() {
		When("enqueuing to empty queue", func() {
			BeforeEach(func() {
				queue.Enqueue(1)
			})

			It("has size one", func() {
				Expect(queue.Size()).To(Equal(1))
				head, err := queue.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(head).To(Equal(1))
			})
		})

		When("enqueuing to list with one element", func() {
			BeforeEach(func() {
				queue.Enqueue(1)
				queue.Enqueue(2)
			})

			It("Has size two", func() {
				Expect(queue.Size()).To(Equal(2))
				head, err := queue.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(head).To(Equal(1))
			})
		})
	})

	Describe("Peek tests", func() {
		When("peeking into empty queue", func() {
			It("returns error", func() {
				head, err := queue.Peek()
				Expect(err).To(HaveOccurred())
				Expect(head).To(Equal(0))
			})
		})

		When("peeking into queue with one element", func() {
			BeforeEach(func() {
				queue.Enqueue(1)
			})

			It("returns that element with no error", func() {
				head, err := queue.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(head).To(Equal(1))
			})
		})

		When("peeking into queue with multiple elements", func() {
			BeforeEach(func() {
				queue.Enqueue(1)
				queue.Enqueue(2)
				queue.Enqueue(3)
			})

			It("returns the first element with no error", func() {
				head, err := queue.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(head).To(Equal(1))
			})
		})
	})

	Describe("Dequeue tests", func() {
		When("dequeuing from empty queue", func() {
			It("throws an error", func() {
				val, err := queue.Dequeue()
				Expect(err).To(HaveOccurred())
				Expect(val).To(Equal(0))
			})
		})

		When("dequeuing from queue with only one element", func() {
			BeforeEach(func() {
				queue.Enqueue(1)
			})

			It("returns and removes this element without throwing an error", func() {
				val, err := queue.Dequeue()
				Expect(err).ToNot(HaveOccurred())
				Expect(val).To(Equal(1))
				head, err := queue.Peek()
				Expect(err).To(HaveOccurred())
				Expect(head).To(Equal(0))
				Expect(queue.Size()).To(Equal(0))
			})
		})

		When("dequeuing from queue with multiple elements", func() {
			BeforeEach(func() {
				queue.Enqueue(1)
				queue.Enqueue(2)
				queue.Enqueue(3)
			})

			It("returns and removes the head of the queue without throwing an error", func() {
				val, err := queue.Dequeue()
				Expect(err).ToNot(HaveOccurred())
				Expect(val).To(Equal(1))
				head, err := queue.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(head).To(Equal(2))
				Expect(queue.Size()).To(Equal(2))
			})
		})
	})
})
