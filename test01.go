/*
Faça um programa que peça o tamanho de um arquivo para download (em MB) e a velocidade de um link de Internet (em Mbps), calcule e informe o tempo aproximado de download do arquivo usando este link (em minutos).
*/

package main

import "fmt"

var InternetSpeed  float64
var FileSize float64
var tempoAproximado float32

func main() {

	fmt.Printf("Digite o tamanho em MB do arquivo: ")
	fmt.Scanln(&FileSize)

	fmt.Printf("Digite a velocidade da internet em Mbps: ")
	fmt.Scanln(&InternetSpeed)
}