package task_1

import (
	"crypto/sha256"
	"github.com/stretchr/testify/require"
	"testing"
)

var numDecimal int = 42           // Десятичная система
var numOctal int = 052            // Восьмеричная система
var numHexadecimal int = 0x2A     // Шестнадцатеричная система
var pi float64 = 3.14             // Тип float64
var name string = "Golang"        // Тип string
var isActive bool = true          // Тип bool
var complexNum complex64 = 1 + 2i // Тип complex64

func getArray() []interface{} {
	return []interface{}{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}
}

func TestArrayVars(t *testing.T) {
	expectedArr := getArray()
	actualArr := getArrayVars(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	require.Equal(t, expectedArr, actualArr)
}

func TestTransformToStr(t *testing.T) {
	arr := getArray()
	require.Equal(t, "4242423.14Golangtrue(1+2i)", transformToStr(arr))
}

func TestTransformToRuneSlice(t *testing.T) {
	require.Equal(t, []rune{'4', '2', ')'}, transformToRuneSlice("42)"))
}

func TestHashWithSaltEven(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("4-=go-2024-2)"))
	require.Equal(t, h.Sum(nil), hashWithSalt([]rune("4-=-2)")))
}

func TestHashWithSaltOdd(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("4-go-2024-2)"))
	require.Equal(t, h.Sum(nil), hashWithSalt([]rune("4--2)")))
}
