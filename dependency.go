package main

import (
	"fmt"
	"sort"
)

func buildPath(services map[string][]string, checked [][]string, currentThread []string, checkAgainst string) [][]string {
	if len(services[checkAgainst]) == 0 {
		//if nothing more to follow append item and end process
		checked = append(checked, currentThread)
		return checked
	} else if contains(services[checkAgainst], checkAgainst) == true {
		//if self referencing append item and end process
		currentThread = append(currentThread, checkAgainst)
		checked = append(checked, currentThread)
		return checked
	} else {
		//map level compare
		for checkNext := range services[checkAgainst] {
			tempThread := append(currentThread, services[checkAgainst][checkNext])
			//figure out if duplicate item in currentThread
			if isLoop(tempThread) == true {
				checked = append(checked, tempThread)
				return checked
			}
			tempMap := buildPath(services, checked, tempThread, services[checkAgainst][checkNext])
			if len(tempMap) > 0 {
				checked = append(checked, tempMap[len(tempMap)-1])
			}
		}
		return checked
	}
}

func checkDependency(dependMap map[string][]string) []string {
	badServices := make([]string, 0)
	for key := range dependMap {
		tempSlice := []string{key}
		results := buildPath(dependMap, make([][]string, 0), tempSlice, key)
		if len(results) > 0 {
			for path := range results {
				if isLoop(results[path]) == true {
					for addItems := range results[path] {
						if contains(badServices, results[path][addItems]) == false {
							badServices = append(badServices, results[path][addItems])
						}
					}
				}
			}
		}
	}
	sort.Strings(badServices)
	return badServices
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
		"a": {},
		"b": {"a"},
		"c": {"b", "a", "d"},
		"d": {"a"},
	}
	results := checkDependency(dependMap)
	fmt.Println(results)
}
