package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Francesca",
			request: "Francesca",
		}, {
			name:    "Vania",
			request: "Vania",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Francesca", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Francesca")
		}
	})
	b.Run("Vania", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Vania")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Michael")
	}
}

func BenchmarkHelloWorldVania(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Vania")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Vania",
			request:  "Vania",
			expected: "Hello Vania",
		},
		{
			name:     "Stevan",
			request:  "Stevan",
			expected: "Hello Stevan",
		},
		{
			name:     "Francesca",
			request:  "Francesca",
			expected: "Hello Francesca",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Stevan", func(t *testing.T) {
		result := HelloWorld("Stevan")
		require.Equal(t, "Hello Stevan", result, "Result must be 'Hello Stevan'")
	})
	t.Run("Vania", func(t *testing.T) {
		result := HelloWorld("Vania")
		require.Equal(t, "Hello Vania", result, "Result must be 'Hello Vania'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Francesca")
	require.Equal(t, "Hello Francesca", result, "Result must be 'Hello Francesca'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Francesca")
	require.Equal(t, "Hello Francesca", result, "Result must be 'Hello Francesca'")
	fmt.Println("TestHelloWorld with require done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Vania")
	assert.Equal(t, "Hello Vania", result, "Result must be 'Hello Vania'")
	fmt.Println("TestHelloWorld with assert done")
}

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
	assert.Equal(t, "Hello Stevan", result)
	// error
	t.Fatal("Result must be 'Hello Stevan'")
	fmt.Println("TestHelloWorldStevan Done")
}
