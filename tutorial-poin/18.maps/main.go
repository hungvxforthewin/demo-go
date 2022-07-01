package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	/* create a map*/
	countryCapitalMap = make(map[string]string)

	/* insert key-value pairs in the map*/
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	/* test if entry is present in the map or not*/
	capital, ok := countryCapitalMap["France"]

	/* if ok is true, entry is present otherwise entry is absent*/

	if ok {
		fmt.Println("Capital of France is", capital)
	} else {

		fmt.Println("Capital of United States is not present")
	}

	// DELETE

	/* create a map*/
	countryCapitalMap_1 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("Original map")

	/* print map */
	for country := range countryCapitalMap_1 {
		fmt.Println("Capital of", country, "is", countryCapitalMap_1[country])
	}

	/* delete an entry */
	delete(countryCapitalMap_1, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("Updated map")

	/* print map */
	for country := range countryCapitalMap_1 {
		fmt.Println("Capital of", country, "is", countryCapitalMap_1[country])
	}
}
