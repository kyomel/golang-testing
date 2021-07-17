package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Michael")

	if result != "Hello Michael" {
		// error
		panic("Result is not 'Hello Michael'")
	}
}

func TestHelloWorldStevan(t *testing.T) {
	result := HelloWorld("Stevan")

	if result != "Hello Stevan" {
		// error
		panic("Result is not 'Hello Stevan'")
	}
}
