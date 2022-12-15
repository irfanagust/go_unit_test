package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Irfan")

	assert.Equal(t, "Hello Irfan", result)

	fmt.Println("Jalan sampe akhir")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Irfan")

	require.Equal(t, "Hello Irfan", result)

	fmt.Println("Jalan sampe akhir")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Unit test tidak dapat berjalan di Linux, Gratisan mulu kontol")
	}

	result := HelloWorld("Irfan")

	require.Equal(t, "Hello Irfan", result)
}

// func TestMain(m *testing.M) {
// 	fmt.Println("Sebelum unit test dijalankan")

// 	m.Run()

// 	fmt.Println("Sesudah unit test dijalankan")
// }

func TestSubTest(t *testing.T) {
	t.Run("subtest_1", func(t *testing.T) {
		result := HelloWorld("Irfan")

		require.Equal(t, "Hello Irfan", result)
	})

	t.Run("subtest_2", func(t *testing.T) {
		result := HelloWorld("Irfan")

		require.Equal(t, "Hello Awk", result)
	})
}

func TestTableTesting(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Irfan",
			request:  "Irfan",
			expected: "Hello Irfan",
		},

		{
			name:     "Tiawan",
			request:  "Tiawan",
			expected: "Hello Tiawan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
