package main

//Importo para imprimir y para usar random
import (
	p "fmt"
	"math/rand"
)

//Inicializo los tamaños de un ejemplo de arreglos
var matrizA [4][5]int
var matrizB [5][3]int
var matrizC [4][3]int

/**
 * Función que implementa la sumatoria que determina los componentes
 * ij de la matriz resultante del producto.
 * Argumentos:
 * i (Int): Es el número de filas de la matriz
 * j (Int): Es el número de columnas de la matriz
 * k (Int): Es el número de columas de la matrizA y filas de matrizB
 * Retorna:
 * Entero, es el componente ij resultante.
 */
func SumatoriaDeTerminos(i int, j int, k int) int {
	if k == 0 {
		return matrizA[i][0] * matrizB[0][j]
	}
	return matrizA[i][k]*matrizB[k][j] + SumatoriaDeTerminos(i, j, k-1)

}

//main
func main() {

	//Inicializo las componentes de matrizA
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			matrizA[i][j] = rand.Intn(20)
		}
	}

	//Inicializo las componentes de matrizB
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			matrizB[i][j] = rand.Intn(20)
		}
	}

	/**
	 * Función que implementa la multiplicación de matrices
	 */
	MultMatrices := func() {
		for i := 0; i < 4; i++ {
			for j := 0; j < 3; j++ {
				matrizC[i][j] = SumatoriaDeTerminos(i, j, 4)
			}
		}
	}

	MultMatrices()
	p.Println(matrizA)
	p.Println(matrizB)
	p.Println(matrizC)

}
