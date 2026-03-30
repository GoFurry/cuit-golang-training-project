package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 7, Add(3, 4))
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	require.NoError(t, err)
	assert.Equal(t, 5.0, result)
}

func TestDivideByZero(t *testing.T) {
	result, err := Divide(10, 0)
	require.Error(t, err)
	assert.Equal(t, 0.0, result)
	assert.Equal(t, "division by zero", err.Error())
}
