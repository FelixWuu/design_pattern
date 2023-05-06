package simple_factory

import (
	"testing"
)

func TestNewAnimal(t *testing.T) {
	testCases := []struct {
		animalType string
		expected   string
	}{
		{"Dog", "Woof!"},
		{"Cat", "Meow!"},
	}

	for idx, testCase := range testCases {
		animal := NewAnimal(testCase.animalType)
		if animal.Speak() != testCase.expected {
			t.Errorf("case%d: expected %s, but got %s", idx+1, testCase.expected, animal.Speak())
		}
	}
}
