/* -/* - Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.*/
package main

import "fmt"

func main() {
	x := `isso é um 
	    raw 
	          string 
	            literal`
	
	fmt.Println(x)
}
