package main

import (
	"fmt"
)

func containsElement(arr []string, target string) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}
func removeNode(neMap map[string][]string, key string, element string) {
	// Find the index of the element in the array
	index := -1
	for i, value := range neMap[key] {
		if value == element {
			index = i
			break
		}
	}
	// If the element is found, remove it from the array using slicing
	if index != -1 {
		neMap[key] = append(neMap[key][:index], neMap[key][index+1:]...)
	}
}

func main() {

	// Create a map with string keys and values as arrays of integers

	neMap := make(map[string][]string)

	// Add data to the map

	neMap["7750_SR"] = []string{"sr_1s", "sr_2s", "sr12", "sr14"}

	neMap["7250_IXR"] = []string{"ixr_6", "ixr_e", "ixr_e2"}

	neMap["7450_ESS"] = []string{"ess_7", "ess_12"}

	neMap["7950_XRS"] = []string{"xrs_20", "xrs_16"}

	// Access and print data from the map

	fmt.Println("Data for 7750SR:", neMap["7750_SR"])

	fmt.Println("Data for 7250IXR:", neMap["7250_IXR"])

	fmt.Println("Data for 7450ESS:", neMap["7450_ESS"])

	fmt.Println("Data for 7950XRS:", neMap["7950_XRS"])

	// Update data in the map

	neMap["7750_SR"] = append(neMap["7750_SR"], "sr_7s")

	// Print updated data for John

	fmt.Println("Updated data for 7750SR:", neMap["7750_SR"])

	// Check if a key exists in the map

	key := "7950_XRS"

	if _, exists := neMap[key]; exists {

		fmt.Println(key, "exists in the map with data:", neMap[key])

	} else {

		fmt.Println(key, "does not exist in the map.")

	}
	node := "xrs_20e"
	if !containsElement(neMap[key], node) {
		neMap["7950_XRS"] = append(neMap["7950_XRS"], "xrs_20e")
		fmt.Println("Updated data for 7950XSR:", neMap["7950_XRS"])
	}

	//Delete depricated
	key1 := "7750_SR"
	node1 := "sr14"

	if containsElement(neMap[key1], node1) {
		removeNode(neMap, key1, node1)
		fmt.Println("Updated data for 7750SR:", neMap["7750_SR"])
	}

}
