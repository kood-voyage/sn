package validator

import (
	"errors"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		Email      string   `validate:"required|email"`
		Username   string   `validate:"min_len:4|max_len:25|alphanumeric|lowercase"`
		TestField  []string `validate:"min_len:1"`
		TestField1 string   `validate:"alpha|uppercase"`
		Age        int
		want       error
	}{
		//Should be correct input
		{
			Email:      "johndoe@gmail.com",
			Username:   "johndoe",
			TestField:  []string{"hello", "world"},
			TestField1: "TEST",
			Age:        2,
			want:       nil,
		},
		//Test require field
		{
			Username: "johndoe",
			want:     errors.New("Email value is required"),
		},
		//Test email
		{
			Email: "SampleGmail@dsadasdas",
			want:  errors.New("Email invalid format"),
		},
		//test min len for string
		{
			Email:    "johndoe@gmail.com",
			Username: "a",
			want:     errors.New("Username length must be minimum of 4"),
		},
		//test max len
		{
			Email:    "johndoe@gmail.com",
			Username: strings.Repeat("a", 26),
			want:     errors.New("Username length must be maximum of 25"),
		},
		//test alpha
		{
			Email:      "johndoe@gmail.com",
			Username:   "johndoe",
			TestField:  []string{"hello", "world"},
			TestField1: "DSA2ASD",
			want:       errors.New("TestField1 string contains more than just letters"),
		},
		//test alphanumeric
		{
			Email:    "johndoe@gmail.com",
			Username: "asd123!",
			want:     errors.New("Username string contains more than just letters and numbers"),
		},
		//Test lowercase
		{
			Email:    "johndoe@gmail.com",
			Username: "AAAAA",
			want:     errors.New("Username only lowercase letters allowed"),
		},
		//test uppercase
		{
			Email:      "johndoe@gmail.com",
			Username:   "johndoe",
			TestField:  []string{"hello", "world"},
			TestField1: "asd",
			want:       errors.New("TestField1 only uppercase letters allowed"),
		},
		//test len for array
		{
			Email:     "johndoe@gmail.com",
			Username:  "johndoe",
			TestField: []string{},
			want:      errors.New("TestField length must be minimum of 1"),
		},
	}

	//test object as a pointer
	pointerTests := []*struct {
		Email    string `validate:"required|email"`
		Username string `validate:"min_len:4|max_len:25|alphanumeric|lowercase"`
		want     error
	}{
		{
			Email:    "JohnDoe@gmail.com",
			Username: "johndoe",
			want:     nil,
		},
	}

	for _, test := range pointerTests {
		err := Validate(test)
		if test.want != nil {
			assertErrorsEqual(t, test.want, err)
		}
	}

	for _, test := range tests {
		err := Validate(test)
		if test.want != nil {
			assertErrorsEqual(t, test.want, err)
		}
	}

}

func assertErrorsEqual(t *testing.T, expected, actual error) {
	if expected != nil && !strings.Contains(actual.Error(), expected.Error()) {
		t.Fatalf("Expected: %v, Got: %v", expected, actual)
	}
}
