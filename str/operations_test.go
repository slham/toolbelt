package str

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	tables := []struct {
		input1    string
		input2    string
		expected1 string
	}{
		{"a", "b", "ab"},
	}

	for _, table := range tables {
		s := Concat(table.input1, table.input2)
		assert.Equal(t, s, table.expected1, "status does not match")
	}
}

func TestToBool(t *testing.T) {
	tables := []struct {
		input     string
		expected1 bool
		expected2 error
	}{
		{"true", true, nil},
		{"false", false, nil},
		{"wrong", false, errors.New("invalid input 'wrong'")},
	}

	for _, table := range tables {
		a, b := ToBool(table.input)
		assert.Equal(t, a, table.expected1, "bool doesn't match")
		if b != nil {
			assert.Equal(t, b.Error(), table.expected2.Error(), "error doesn't match")
		}
	}
}
