package str

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
)

func Concat(s1 string, s2 string) string {
	return fmt.Sprintf(s1 + s2)
}

//Converts a string to a boolean; defaults to false
func ToBool(s string) (bool, error) {
	var err error
	output := false
	arr := bytes.ToLower([]byte(s))
	str := string(arr)
	if "true" == str {
		output = true
	} else if "false" == str {
		output = false
	} else {
		msg := fmt.Sprintf("invalid input '%v'", s)
		err = errors.New(msg)
	}
	return output, err
}

func FromBool(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func subStrings(str string, size int)[]string {
	var out []string

	for i := 0; i < size; i++ {
		for j := i+1; j <= size; j++ {
			out = append(out, str[i:j])
		}
	}

	return out
}