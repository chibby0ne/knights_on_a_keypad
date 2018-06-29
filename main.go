package main

import (
	"flag"
	"fmt"
)

const (
    // NUMBERS are a Keypad
    NUMBERS int = 10
)

var moves [NUMBERS][]int
var start *int
var numLength *int
var help *bool
var cache map[string]int

// GetMoves returns an array containing the numbers reachable from the given
// start number position
func GetMoves(start int) []int {
	switch start {
	case 0:
		return []int{4, 6}
	case 1:
		return []int{6, 8}
	case 2:
		return []int{7, 9}
	case 3:
		return []int{4, 8}
	case 4:
		return []int{3, 9, 0}
	case 5:
		return nil
	case 6:
		return []int{1, 7, 0}
	case 7:
		return []int{2, 6}
	case 8:
		return []int{1, 3}
	case 9:
		return []int{2, 4}
	default:
		return nil
	}
}

func init() {
	start = flag.Int("startposition", 3, "Starting Number")
	numLength = flag.Int("length", 3, "Telephone number length")
	help = flag.Bool("help", false, "Show help info")
	flag.Parse()
	for i := 0; i < NUMBERS; i++ {
		moves[i] = GetMoves(i)
	}
    cache = make(map[string]int)
}

func makeKey(position, length int) string {
    return fmt.Sprintf("%v_%v", position, length)
}


// GetNumberOfPossibleNumbers returns the number of possible movies for the
// given starting position and length of the phone number
func GetNumberOfPossibleNumbers(position, length int) int {
    cacheKey := makeKey(position, length)
    val, ok := cache[cacheKey]
    if ok {
        return val
    }
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}
	length = length - 1
	num := 0
	for _, pos := range moves[position] {
		num += GetNumberOfPossibleNumbers(pos, length)
	}
    cache[cacheKey] = num
	return num
}

func main() {
	fmt.Printf("Using knight moves and a keypad\n1 2 3\n4 5 6\n7 8 9\n  0  \n")
	fmt.Printf("Number: %d as start position and a phone number of length: %d\n", *start, *numLength)
	fmt.Printf("The number of possible phone numbers are: %d\n", GetNumberOfPossibleNumbers(*start, *numLength))
}
