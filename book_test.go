package book

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_Create_First_Book_Price_8_EUR(t *testing.T) {
	bookPart1 := NewBook("Harry Potter Part 1", 8)

	assert.Equal(t, 8, bookPart1.getPrice())
	assert.Equal(t, "Harry Potter Part 1", bookPart1.getName())
}

func Test_Create_Second_Book_Price_8_EUR(t *testing.T) {
	bookPart2 := NewBook("Harry Potter Part 2", 8)

	assert.Equal(t, 8, bookPart2.getPrice())
	assert.Equal(t, "Harry Potter Part 2", bookPart2.getName())
}
