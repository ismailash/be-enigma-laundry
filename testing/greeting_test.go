package testing

import "testing"

func TestSayHello_Success(t *testing.T) {
	// ekspektasi
	expected := "Hello Budi"

	// eksekusi
	actual, _ := SayHello("Budi")

	// assertion
	if expected != actual {
		t.Fatalf(`SayHello() failed, expected: %v, actual: %v`, expected, actual)
	}
}

func TestSayHello_Fail_Empty_Name(t *testing.T) {
	// ekspektasi
	expected := "name can't be empty"

	// eksekusi
	_, actual := SayHello("")

	// assertion
	if actual.Error() != expected {
		t.Fatalf(`SayHello() failed, expected: %v, actual: %v`, expected, actual)
	}
}

// TestSayHello_Fail_Invalid_Input
// TestSayHello_Fail_Minimal_Length
// Dan lain lain
