package main

import (
	"fmt"
)

func main() {
	// TODO: create a map called "modules"

	modules := map[int]string{
		104: "Modul 104 ist genial",
		117: "Modul 117 ist hammer",
		346: "Modul 346 ist mega",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117) // TODO: delete one
	modules[123] = "Modul 123" // TODO: add one
	modules[104] = "Modul 104 ist supergenial"// TODO: replace one
	fmt.Println(modules)
}
