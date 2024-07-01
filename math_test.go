package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {

	total := soma(15, 15)
	assert.Equal(t, 30, total)

}
