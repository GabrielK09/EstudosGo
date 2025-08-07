package main

import (
	"fmt"
)

// uint8 = byte
// uint16
// uint32
// uint64

// int8
// int16
// int32 = rune
// int64
// Valor padrão = 0
// Se não definido o 64 32, e apenas deixar o int, vai pegar o do computador, 64bits = int64

// float32
// float64
// Valor padrão = 0
// Se não definido o 64 32, e apenas deixar o int, vai pegar o do computador, 64bits = int64

// string
// Uma string
// Valor padrão = ""

// bool
// Valor padrão = false

func main() {
	var eint8 int8
	var eint16 int16
	var eint32 int32
	var eint64 int64

	eint8, eint16 = 12, 12
	fmt.Println(eint8)  // 12
	fmt.Println(eint16) // 12
	fmt.Println(eint32) // 0
	fmt.Println(eint64) // 0

	// Alias
	// byte = uint8
	// rune = int16
	var byteEint8 byte
	var runeEint16 rune

	runeEint16 = 123
	fmt.Println(byteEint8)  // 0
	fmt.Println(runeEint16) // 123

	// uint
	var x uint16 = 65535
	fmt.Printf("Uint: %d\n", x) // 0

	// ----------------------------------------------------- \\

	var varFloat32 float32
	var varFloat64 float64
	fmt.Println(varFloat32) // 0
	fmt.Println(varFloat64) // 0

	// ----------------------------------------------------- \\

	var varString string
	fmt.Println(varString) // ""

	// ----------------------------------------------------- \\
}
