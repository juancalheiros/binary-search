package dojo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckCentralRoot1(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3}
	value := 3
	position := 0
	resp := checkCentralRoot(value, position, arr)

	expected := false
	assert.Equal(expected, resp, "must array [1,2,3] check position in array should be false with value for search equals 3")
}

func TestCheckCentralRoot2(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3}
	value := 2
	position := 1
	resp := checkCentralRoot(value, position, arr)

	expected := true
	assert.Equal(expected, resp, "must array [1,2,3] check position in array should be false with value for search equals 2")
}

func TestCheckCentralRoot3(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3}
	value := 1
	position := 2
	resp := checkCentralRoot(value, position, arr)

	expected := false
	assert.Equal(expected, resp, "must array [1,2,3] check position in array should be false with value for search equals 1")
}

func TestBinarySearch1(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3}
	value := 0
	resp := binarySearch(value, arr)

	expected := false
	assert.Equal(expected, resp, "must array [1,2,3] should be not found value 0")
}

func TestBinarySearch2(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3}
	value := 1
	resp := binarySearch(value, arr)

	expected := true
	assert.Equal(expected, resp, "must array [1,2,3] should be found value 1")
}

func TestBinarySearch3(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1}
	value := 2
	resp := binarySearch(value, arr)

	expected := false
	assert.Equal(expected, resp, "must array [1] should be not found value 2")
}

func TestBinarySearch4(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	value := 1
	resp := binarySearch(value, arr)

	expected := true
	assert.Equal(expected, resp, "must array [1,2,3,4,5,6,7] should be found value 1")
}

func TestBinarySearch5(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	value := 6
	resp := binarySearch(value, arr)

	expected := true
	assert.Equal(expected, resp, "must array [1,2,3,4,5,6,7] should be found value 7")
}
