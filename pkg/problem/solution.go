package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	startInterval int
	endInterval int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	intervalList := make([]Interval, 0)
	//Iterate for STDIN from cmd until exit
	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			inputInterval := strings.Split(scanner.Text(), " ")
			start, _ := strconv.Atoi(inputInterval[0])
			end, _ := strconv.Atoi(inputInterval[1])
			intervalList = append(intervalList, Interval{startInterval: start, endInterval: end})
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort itemList by start point of each element ie - ["3 6", "2 6"] -> ["2 6", "3, 6"]
	//Sorting inputList slice while keeping the original order of equal starting point same
	//TO DO: Throw error when split string is empty
	sort.SliceStable(intervalList, func(i, j int) bool {
		return intervalList[i].startInterval < intervalList[j].startInterval
	})

	intervalStack := make([]Interval, 0)
	// Push inputInterval items into intervalStack
	intervalStack = append(intervalStack, intervalList[0])
	for _, value := range intervalList[1:] {
		if(overlapOrNot(value, intervalStack[len(intervalStack)-1]) == false){
			intervalStack = append(intervalStack, value)
		} else if(overlapOrNot(value, intervalStack[len(intervalStack)-1]) == true && compareEnd(value, intervalStack[len(intervalStack)-1]) == true) {
			intervalStack[len(intervalStack)-1].endInterval = value.endInterval
		}
	}

	// Sorted list of merged intervals
	fmt.Println(intervalStack)

	// Iterate over the intervalStack and output the final list of intervals that are not covers by input intervals.
	for index, element := range intervalStack[:len(intervalStack)-1] {
		fmt.Println(element.endInterval, " ", intervalStack[index+1].startInterval)
	}

}

// Check if last element in stack overlapped to current element in intervalList inside the for loop
func overlapOrNot(listItem, stackItem Interval) bool {
	return listItem.startInterval <= stackItem.endInterval
}

// Compare endInterval in stack and current intervalList
func compareEnd(listItem, stackItem Interval) bool {
	return listItem.endInterval > stackItem.endInterval
}
