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

func TestMain(m *testing.M) {
	fmt.Println("Sebelum unit test dijalankan")

	m.Run()

	fmt.Println("Sesudah unit test dijalankan")
}

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
