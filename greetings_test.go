package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	// given
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	// when
	msg, err := Hello("Gladys")
	// then
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	// given
	name := ""
	// when
	msg, err := Hello(name)
	// then
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellosNames(t *testing.T) {
	// given
	names := []string{"Gladys", "Joe", "Jack"}
	// when
	greetingsMap, err := Hellos(names)
	// then
	for name, greeting := range greetingsMap {
		namePattern := regexp.MustCompile(`\b` + name + `\b`)
		containsName := namePattern.MatchString(greeting)
		if !containsName {
			t.Fatalf(`Hellos([]string{"Gladys, Joe, Jack"}) break for name %q, value %v`, name, greeting)
		}
	}
	if err != nil {
		t.Fatalf(`For Hellos([]string{"Gladys, Joe, Jack"}) err should be nil`)
	}
}

func TestHellosEmptyList(t *testing.T) {
	// given
	names := []string{}
	// when
	greetingsMap, err := Hellos(names)
	// then
	if greetingsMap != nil || err == nil {
		t.Fatalf(`Hellos([]string{}) = %q, %v, want nil, error`, greetingsMap, err)
	}
}

func TestHellosEmptyNameInList(t *testing.T) {
	// given
	names := []string{"Gladys", "", "Jack"}
	// when
	greetingsMap, err := Hellos(names)
	// then
	if greetingsMap != nil || err == nil {
		t.Fatalf(`Hellos([]string{"Gladys", "", "Jack"}) = %q, %v, want nil, error`, greetingsMap, err)
	}
}
