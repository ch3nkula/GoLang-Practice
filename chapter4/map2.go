/* Maps, Slices and Arrays */

package main

import "fmt"

func main() {
	elements := make(map[string]string)

	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Berylium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements)
	fmt.Println(elements["Ne"])

	//name, ok := elements["Un"]
	//fmt.Println(name, ok)

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}