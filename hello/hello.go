//declare a main package
//packages are ways to group functions
//, all files in same dir
package main

//import fmt, which contains functions for
//formatting text.
import "fmt"

import "rsc.io/quote"

//main function executes by default to print message
func main() {
	fmt.Println(quote.Go())
}
