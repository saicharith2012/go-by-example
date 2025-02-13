package main

import (
	"fmt";
	"math"
)

// constants are like variables but their values cannot be changed after they are declared.
// used for values that are known at compile time and remain fixed throughout the execution.
// constants are typically declared at the package level but can also be declared within functions.

const s string = "anniyaa"

func main() {
	fmt.Println(s)

	const n = 500000000 // no explicit type is given...type inference happens.
	// constants in Go can also be untyped which are more flexible and can be used in more contexts.

	const d = 3e20 / n // calculations involving constants are also constants ?? --> calculation performed at compile time.
	// incase of using var for d, the calculation is performed at runtime.
	// compile time calculations are faster, since they are done only once during the compilation and not for every run.


	fmt.Println(d)

	fmt.Println(int64(d)) // type conversion/type casting into an integer type

	fmt.Println(math.Sin(n)) // sine function imported from the "math" package


}
