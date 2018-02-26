package main

import (
	"testing"
)

func TestGetNumberOfPossibleNumbers(t *testing.T) {
	tables := []struct {
		position int
		length   int
		result   int
	}{
		{4, 3, 6},
		{1, 3, 5},
		{2, 3, 4},
		{0, 2, 2},
	}

	for _, table := range tables {
		actualResult := GetNumberOfPossibleNumbers(table.position, table.length)
		if table.result != actualResult {
			t.Errorf("getNumberOfPossibleNumbers(%d, %d) was incorrect. Got %d, want %d\n", table.position, table.length, actualResult, table.result)
		}
	}
}

func TestGetMoves(t *testing.T) {
	tables := []struct {
		num    int
		result []int
	}{
		{0, []int{4, 6}},
		{1, []int{6, 8}},
		{2, []int{7, 9}},
		{3, []int{4, 8}},
		{4, []int{3, 9, 0}},
		{5, nil},
		{6, []int{1, 7, 0}},
		{7, []int{2, 6}},
		{8, []int{1, 3}},
		{9, []int{2, 4}},
	}

	for _, table := range tables {
		actualResult := GetMoves(table.num)
		if equals(table.result, actualResult) == false {
			t.Errorf("getMoves(%d) was incorrect. Got %v, want %v\n", table.num, actualResult, table.result)
		}
	}
}

func equals(sliceA, sliceB []int) bool {
	if sliceA == nil && sliceB == nil {
		return true
	}
	if sliceA == nil || sliceB == nil {
		return false
	}
	if len(sliceA) != len(sliceB) {
		return false
	}
	for i := range sliceA {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}
	return true
}
