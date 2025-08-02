package main

import "fmt"

func main() {
	DemonstrateFactoryPattern()
	DemonstrateAbstractFactoryPattern()
	DemonstrateRegistryPattern()

	fmt.Println("=== PATTERN COMPARISON SUMMARY ===")
	fmt.Println("Factory: Simple, fixed types, compile-time binding")
	fmt.Println("Abstract Factory: Families of related products, provider-based")
	fmt.Println("Registry: Dynamic, runtime registration, most flexible")
}
