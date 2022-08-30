package foo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	assert.Equal(t, "Foo", Foo(), "they should be equal")
}
