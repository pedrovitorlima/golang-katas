package phone_number

import (
	"testing"
)

func TestShouldCreatePhoneStringGivenValidInput(t *testing.T) {
	input := [10]uint{5, 5, 1, 2, 3, 5, 4, 5, 7, 0}
	expected := "(551) 235-4570"

	phone := CreatePhoneNumber(input)

	if phone != expected {
		t.Errorf("Expected phone number to be %s, but it was %s", expected, phone)
	}

}
