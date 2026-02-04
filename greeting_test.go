package greetings

import (
	"regexp"
	"testing"
)

// teste la funcion Hello() cuando le pasamos un nombre
func TexthelloName(t *testing.T) {
	name := "Santiago"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Santiago")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiuere coincidencia para %#q`, msg, err, want)
	}
}

// Comprueba el manejo de error de la funcion Hello()
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
