package simpleinterest

// Init functions initializes the package
// This function is called when the package is imported and is used to set up
// any necessary state or configuration for the package.
// It is not mandatory to have this function in a package but it can be useful
// for logging or initializing package-level variables.
// In this case, we are just printing a message to indicate that the package
// has been initialized.
import (
	"fmt"
)

func init() {
	fmt.Println("Simpleinterest package initialized")
}

// Calculate calculates and returns the simple interest
// for a principal p, rate of interest r for time duration t years
func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
