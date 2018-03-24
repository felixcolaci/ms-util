package jnlx_util

import "testing"

var emails = []struct{
	email string
	result bool
}{
	{"felix.colaci@onlinehome.de", true},
	{"carlos@example.co.uk", true},
	{"carlos.com", false},
	{"carlos", false},
	{"carlos@example", true},
}

func TestValidateEmail(t *testing.T) {

	for _, tt := range emails {

		result := ValidateEmail(tt.email)
		if result != tt.result {
			t.Errorf("Email validation failed for %v | Expected %v got %v", tt.email, tt.result, result)
		}

	}

}
