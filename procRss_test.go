package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArgsPars(t *testing.T) {
	pid, _ := parsePID([]string{"exe", "1234"})
	assert.Equal(t, []int{1234}, pid, "wrong pid")
}

func TestArgsParsFail(t *testing.T) {
	_, err := parsePID([]string{"exe"})
	assert.Error(t, err)
}
