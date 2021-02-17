package main

import (
	"testing"

	"github.com/franela/goblin"
)

func TestAdd(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Add", func() {
		g.It("it should return the sum of the received numbers", func() {
			currentResult := Add(1, 2)

			g.Assert(currentResult).Equal(3)
		})
	})
}
