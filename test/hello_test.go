package test

import (
	"testing"
)

/*
hello.go faylni ichidegi  testfaylda yani hello_test.go da korinad
lekin test hallo_test.go ni ichidegi bowqa hichqaysi faylda kornmdi
*/
func TestHello(t *testing.T) {
	//case
	t.Run("user 1", func(t *testing.T) {
		name := "Aziz"
		want := "Hello " + name
		assertHello(t, name, want)
	})
	//case 2
	t.Run("user 2", func(t *testing.T) {
		name := "nurmat"
		want := "Hello " + name
		assertHello(t, name, want)
	})
	//case 3

	t.Run("empty case", func(t *testing.T) {
		name := ""
		want := "Hello World" + name
		assertHello(t, name, want)

	})

	t.Run("number case",func(t *testing.T) {
		name:="2345"
		want:="Hello Adam"
		assertHello(t,name,want)
	})
}

// assert -tasdiqlash
func assertHello(t *testing.T, name, want string) {
	t.Helper()
	msg := Hello(name)
	if want != msg {
		t.Errorf("expected %q but wanted %q", want, msg)
	}
}
