package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 100000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// NÚMEROS REAIS

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// FIM NÚMEROS REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	// VALOR ZERO - NO GO TODA VARIÁVEL NÂO INICIALIZADA TEM SEU VALOR ZERO, NO CASO DE STRING É VAZIO

	var booleano1 bool = true //O VALOR ZERO DO BOOLEANO É FALSE
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno") // O VALOR ZERO DO ERROR É NIL
	fmt.Println(erro)

}
