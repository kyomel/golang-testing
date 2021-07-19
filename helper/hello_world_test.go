package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldMichael(t *testing.T) {
	result := HelloWorld("Michael")

	if result != "Hello Michael" {
		// error
		t.Error("Result must be 'Hello Michael'")
	}

	fmt.Println("TestHelloWorldMichael Done")
}

func TestHelloWorldStevan(t *testing.T) {
	result := HelloWorld("Stevan")

	if result != "Hello Stevan" {
		// error
		t.Fatal("Result must be 'Hello Stevan")
	}
	fmt.Println("TestHelloWorldStevan Done")
}
