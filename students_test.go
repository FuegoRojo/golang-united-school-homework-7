
package coverage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

const (
	single            = "single row"
	multiple          = "multiple row"
	singleString      = "0 1 2 3"
	multipleString    = "0 1 2\n3 4 5\n6 7 8"
	testLessErrorText = "testLess error"
	testSwapErrorText = "testSwap error"
)

var (
	a = People{Person{"Adam", "Smith", time.Date(1988, 02, 01, 0, 0, 0, 0, time.UTC)}}
	b = append(a, Person{"Boris", "Johnson", time.Date(1988, 02, 01, 0, 0, 0, 0, time.UTC)})
	p = append(b, Person{"Jack", "Butler", time.Date(1988, 02, 03, 0, 0, 0, 0, time.UTC)})
)

func TestLen(t *testing.T) {
	assert.Equal(t, p.Len(), 3, "Ошибка в определении длины")
}

func TestLess(t *testing.T) {
	assert.Equal(t, p.Less(0, 1), true, testLessErrorText)
	assert.Equal(t, p.Less(0, 2), false, testLessErrorText)
	assert.Equal(t, p.Less(1, 1), false, testLessErrorText)
}

func TestSwap(t *testing.T) {
	person1, person2 := p[0], p[1]
	p.Swap(0, 1)
	swap1, swap2 := p[0], p[1]
	assert.NotEqual(t, swap1, person1, testSwapErrorText)
	assert.NotEqual(t, swap2, person2, testSwapErrorText)
}

func TestNew(t *testing.T) {
	tData := map[string]string{
		"zeroV": "",
		"someV": "1 1\n0 2 2",
	}
	for _, v := range tData {
		actual, err := New(v)
		assert.NotNil(t, err)
		assert.Nil(t, actual)
	}
}

func TestRows(t *testing.T) {
	tData := map[string]struct {
		input string
		rows  [][]int
	}{
		single:   {singleString, [][]int{{0, 1, 2, 3}}},
		multiple: {multipleString, [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}},
	}
	for _, v := range tData {
		matrix, _ := New(v.input)
		assert.Equal(t, v.rows, matrix.Rows())
	}
}

func TestCols(t *testing.T) {
	tData := map[string]struct {
		input string
		cols  [][]int
	}{
		single:   {singleString, [][]int{{0}, {1}, {2}, {3}}},
		multiple: {multipleString, [][]int{{0, 3, 6}, {1, 4, 7}, {2, 5, 8}}}}

	for _, v := range tData {
		matrix, _ := New(v.input)
		assert.Equal(t, v.cols, matrix.Cols())
	}
}
func TestSet(t *testing.T) {
	tData := map[string]struct {
		input string
		row   int
		col   int
		value int
	}{
		single:   {"0 1 2", 0, 1, 3},
		multiple: {multipleString, 2, 2, 9}}

	for _, v := range tData {
		matrix, _ := New(v.input)
		assert.True(t, matrix.Set(v.row, v.col, v.value))
		assert.Equal(t, v.value, matrix.Rows()[v.row][v.col])
	}
}

func TestMatrixFalse(t *testing.T) {
	matrix, _ := New("0 1 2")
	tData := map[string]struct {
		row int
		col int
	}{
		"negative row":     {-1, 0},
		"row out of range": {1, 0},
		"negative col":     {0, -1},
		"col out of range": {0, 3},
	}

	for _, v := range tData {
		assert.False(t, matrix.Set(v.row, v.col, 100))
	}
}
