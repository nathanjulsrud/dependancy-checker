package main

import (
	"fmt"
)

func buildPath(services map[string]string, checked []string, checkAgainst string) []string {
	if services[checkAgainst] == "" {
		//if nothing more to follow append item and end process
		checked = append(checked, checkAgainst)
		return checked
	} else if services[checkAgainst] == checkAgainst {
		//if self referencing append item and end process
		checked = append(checked, services[checkAgainst])
		checked = append(checked, services[checkAgainst])
		return checked
	} else {
		checked = append(checked, checkAgainst)
		for checkService := range checked {
			var serviceCount int = 0
			for checkAgainst := range checked {
				if checked[checkService] == checked[checkAgainst] {
					serviceCount++
				}
			}
			//if count of any stop is greater than one this is a loop
			if serviceCount > 1 {
				return checked
			}
		}
		checked = buildPath(services, checked, services[checkAgainst])
	}
	return checked
}

func findDependancyLoop(services map[string][]string) bool {
	//resultMap := make(map[string]bool)
	for compareService := range services {
		fmt.Println(compareService)
		subDependancy := services[compareService]
		fmt.Println("sub", subDependancy)
		for compare := range subDependancy {
			fmt.Println(subDependancy[compare])
		}
		/*for dependancy := range services[compareService] {
			fmt.Println("service", services[compareService])
			fmt.Println("depend", dependancy)
		}*/
	}

	return false
}

func main() {
	dependMap := map[string][]string{
		"a": {},
		"b": {"a"},
		"c": {"b", "a"},
	}

	result := findDependancyLoop(dependMap)
	fmt.Println(result)
	//fmt.Println("result", buildPath(dependMap, make([]string, 0), service))

}
