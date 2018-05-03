package ms_util_test

import (
	"testing"
	"github.com/felixcolaci/ms-util"
)

func TestHashPassword(t *testing.T) {

	salt, hash, err := ms_util.HashPassword("password")
	if len(salt) < 1 || len(hash) < 1 || err != nil {
		t.Errorf("hashing failed because of %v", err)
	}

}

func TestCheckPasswordHash(t *testing.T) {

	var testcases = []struct{
		name string
		pwToHash string
		pwToCompare string
		result bool
	}{
		{"valid password", "password", "password", true},
		{"valid short password", "asd", "asd", true},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			salt, hash, err := ms_util.HashPassword(tc.pwToHash)
			if len(salt) < 1 || len(hash) < 1 || err != nil {
				t.Fatalf("hashing failed because of %v", err)
			}

			result := ms_util.CheckPasswordHash(salt, tc.pwToCompare, hash)

			if result != tc.result {
				t.Errorf("password mismatched expectation; expected %v got %v", tc.result, result)
			}
		})
	}


}

