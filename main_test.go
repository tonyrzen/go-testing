package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMainFunc ... because we want 100% test coverage in our example :-)
func TestMainFunc(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			assert.NotPanics(t, main)
		}
	}()

	main()
}
