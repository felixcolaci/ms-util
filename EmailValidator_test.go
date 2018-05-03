package ms_util_test

import (
	"testing"
	"github.com/felixcolaci/ms-util"
)

var emails = []struct {
	name string
	email  string
	result bool
}{
	{"valid mail", "felix.colaci@onlinehome.de", true},
	{"multiple dots in tld", "carlos@example.co.uk", true},
	{"domain only", "carlos.com", false},
	{"without @", "carlos", false},
	{"valid mail without tld", "carlos@example", true},
}

func TestValidateEmail(t *testing.T) {

	for _, tc := range emails {
		t.Run(tc.name, func(t *testing.T) {
			result := ms_util.ValidateEmail(tc.email)
			if result != tc.result {
				t.Errorf("Email validation failed for %v | Expected %v got %v", tc.email, tc.result, result)
			}
		})
	}
}
