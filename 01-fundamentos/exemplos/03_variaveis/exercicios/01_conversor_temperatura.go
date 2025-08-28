/*
*
* Exercício 1:
* Crie uma variável em °C e converta para °F. Fórmula: `F = (C * 9/5) + 32`.
*
 */

package main

import "fmt"

/*
1. **Conversor de temperatura**

   * Crie uma variável em °C e converta para °F.
     Fórmula: `F = (C * 9/5) + 32`.
*/

func main() {
	var celsius float64
	fmt.Print("Digite a temperatura em °C: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("Temperatura em °F: %.2f\n", fahrenheit)
}
