/* Dependency for external packages */
package utils

import "fmt"

// Must start with capital case to be used in external packages
func HelloWorld() {
	fmt.Println("Hello World from my new package")
}
