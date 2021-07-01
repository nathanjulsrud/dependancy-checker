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
		return checked
	} else if services[checkAgainst] != checkAgainst {
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

func findDependancyLoop(services map[string]string) bool {

	return false
}

func main() {
	dependMap := map[string]string{
		"aa": "c",
		"b":  "",
		"c":  "b",
	}

	fmt.Println("result", buildPath(dependMap, make([]string, 0), "aa"))
}

/*check group a
1	if a references nothing pass
2	if a references a fail
	if a references b
	  check group b
		1
		2
		3 if b references a fail
		4 if b references c
*/
