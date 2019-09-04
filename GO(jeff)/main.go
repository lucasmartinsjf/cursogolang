package main

import "fmt"
import "github.com/jeffprestes/cursodego/fundamentos/funcoes_basico/matematica"

func main() {
	resultado := multiplicador(2, 2)
	fmt.Printf("O resultado da multiplicação foi %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da Soma foi %d\r\n", resultado)
}
func multiplicador(x int, y int) int {
	return x * y

}
