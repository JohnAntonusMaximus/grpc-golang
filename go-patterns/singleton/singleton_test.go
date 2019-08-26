package creational

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("Expected pointer to a singleton after calling")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count should be 1, but got %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Errorf("Expected same instance in counter2, but got a different counter instance.\n")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling for the first time to count, the count should be 2, but got %d\n", currentCount)
	}

}
