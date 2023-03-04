package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Artur"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Artur")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Artur) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmptry(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
