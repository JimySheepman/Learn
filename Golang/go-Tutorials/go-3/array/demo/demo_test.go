package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateArray1(t *testing.T) {
	testArray := [2]string{"Value1", "Value2"}
	UpdateArray1(testArray)
	assert.Equal(t, NewValue, testArray[0])
}

func TestUpdateArray2(t *testing.T) {
	testArray := [2]string{"Value1", "Value2"}
	UpdateArray2(&testArray)
	assert.Equal(t, NewValue2, testArray[0])
}

func TestArrayCopy(t *testing.T) {
	testArray := [2]string{"Value1", "Value2"}
	newCopy := testArray
	testArray[1] = "updated"
	assert.Equal(t, "updated", newCopy[1])
}

func TestArrayReference(t *testing.T) {
	testArray := [2]string{"Value1", "Value2"}
	reference := &testArray
	testArray[1] = "updated"
	assert.Equal(t, "updated", reference[1])
}
