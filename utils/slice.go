package utils

import "fmt"

func RemoveLast(str string) string {
	return str[:len(str)-1]
}

func SliceAddOperator(slice []int, value int) []int {
	for i := range slice {
		slice[i] += value
	}
	return slice
}

func SliceMultiplyOperator(slice []int, value int) []int {
	for i := range slice {
		slice[i] *= value
	}
	return slice
}

func SliceMultiplyTogether(slice []int) int {
	res := 1
	for _, v := range slice {
		res *= v
	}
	return res
}

func InsertBefore[T any](slice []T, at int, value T) []T {

	if at < 0 {
		//add a the start
		return append([]T{value}, slice...)
	}
	// insert before the pos
	slc := make([]T, len(slice)+1)
	copy(slc, slice[:at])
	slc[at] = value
	copy(slc[at+1:], slice[at:])
	return slc
}

func InsertAfter[T any](slice []T, at int, value T) []T {
	// insert after the pos
	slc := make([]T, len(slice)+1)
	copy(slc, slice[:at+1])
	slc[at+1] = value
	copy(slc[at+2:], slice[at+1:])
	return slc
}

func RemoveFirstSlice[T any](slice []T) []T {
	return slice[1:]
}

func RemoveAt[T any](slice []T, at int) []T {
	// fmt.Println("Removing at", at, "from", slice)
	slc := make([]T, len(slice)-1)
	copy(slc, slice[:at])
	copy(slc[at:], slice[at+1:])
	return slc
}

func MinIntIndex(slice []int) int {
	minid := 0
	for i, v := range slice {
		if v < slice[minid] {
			minid = i
		}
	}
	return minid
}

func MaxIntIndex(slice []int) int {
	maxid := 0
	for i, v := range slice {
		if v > slice[maxid] {
			maxid = i
		}
	}
	return maxid
}

func StrSliceToIntSlice(slice []string) []int {
	res := make([]int, len(slice))
	for i, v := range slice {
		res[i] = ParseInt(v)
	}
	return res
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Get the values matching the pattern
// £ = the int to get
// !!! NOT WORKING YET !!!
func GetValueInsidePattern(input, pattern string, crs *int) ([]int, error) {
	var values []int
	vId := 0
	patternPtr := 0
	for *crs < len(input) {
		fmt.Println(*crs)
		if input[*crs] == pattern[patternPtr] {
			vId++
			patternPtr++

		} else if pattern[patternPtr] == '£' && IsInt(input[*crs]) {
			v := ParseInt(input[*crs])
			if vId == 0 {
				values = append(values, v)
			} else {
				values[vId-1] = ConcatInt(values[vId-1], v)
			}
		} else {
			*crs++
			return nil, fmt.Errorf("No matching pattern")
		}
		*crs++
	}
	return values, nil
}
