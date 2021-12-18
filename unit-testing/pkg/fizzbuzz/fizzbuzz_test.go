package fizzbuzz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	// zeros test
	fizBuzz, err := FizzBuzz(0, 1, 1 )
	assert.Nil(t, err)
	assert.Equal(t, len(fizBuzz), 0, "Zero test should return zero length slice")

	// Test mod by zero
	fizBuzz, err = FizzBuzz(10, 0, 5)
	assert.NotNil(t, err)
	assert.Nil(t, fizBuzz, "buzzAt of zero should result in null slice")
	fmt.Println(fizBuzz)

	// Test mod by zero
	fizBuzz, err = FizzBuzz(10, 5, 0)
	assert.NotNil(t, err, nil)
	assert.Nil(t, fizBuzz, "fizzAt of zero should result in null slice")

	// Limit test
	fizBuzz, err = FizzBuzz(65, 3, 5)
	assert.NotNil(t, err)
	assert.Nil(t, fizBuzz, "Total greater than 64 should result in limit error")

	// classic test
	fizBuzzClassic := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	fizBuzz, err = FizzBuzz(15, 3, 5)
	assert.Nil(t, err)
	assert.Equal(t, len(fizBuzz), len(fizBuzzClassic))
	for i, val := range fizBuzz{
		assert.Equal(t, val, fizBuzzClassic[i])
	}


}

