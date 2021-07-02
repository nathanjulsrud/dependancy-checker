package main

import (
	"fmt"
)

func buildPath(services map[string][]string, checked [][]string, currentThread []string, checkAgainst string) [][]string {
	if len(services[checkAgainst]) == 0 {
		//if nothing more to follow append item and end process
		fmt.Println("zero")
		checked = append(checked, currentThread)
		return checked
	} else if contains(services[checkAgainst], checkAgainst) == true {
		//if self referencing append item and end process
		currentThread = append(currentThread, checkAgainst)
		checked = append(checked, currentThread)
		fmt.Println("self reference")
		return checked
	} else {
		fmt.Println("else!")
		//map level compare
		for checkNext := range services[checkAgainst] {
			currentThread = append(currentThread, services[checkAgainst][checkNext])
			//figure out if duplicate item in currentThread
			if isLoop(currentThread) == true {
				checked = append(checked, currentThread)
				return checked
			}
			//checked = buildPath(services, checked, currentThread, services[checkAgainst])
		}
	}
	return checked
}

func contains(stringSlice []string, stringFind string) bool {
	for value := range stringSlice {
		if stringSlice[value] == stringFind {
			return true
		}
	}
	return false
}

func isLoop(stringSlice []string) bool {
	for checkService := range stringSlice {
		var serviceCount int = 0
		for checkAgainst := range stringSlice {
			if stringSlice[checkService] == stringSlice[checkAgainst] {
				serviceCount++
				//if count of any stop is greater than one this is a loop
				if serviceCount > 1 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	dependMap := map[string][]string{
		"a": {"a"},
		"b": {""},
		"c": {"b", "a"},
	}

	tempSlice := []string{"a"}
	fmt.Println("result", buildPath(dependMap, make([][]string, 0), tempSlice, "a"))

}
