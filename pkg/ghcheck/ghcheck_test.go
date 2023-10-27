package ghcheck

import (
	"reflect"
	"testing"
)

func TestGithubApiCall(t *testing.T) {
	str := reflect.TypeOf("string")
	res := GithubCall("/users/absozero")
	if reflect.TypeOf(res) != str {
		t.Errorf("Incorrect data type, got %v instead of %v", reflect.TypeOf(res), str)
	}
}
