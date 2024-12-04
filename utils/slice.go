package utils

import "fmt"

func RemoveLast(str string) string {
	return str[:len(str)-1]
}

func SliceMultiply(slice []int) int {
	res := 1
	for _, v := range slice {
		res *= v
	}
	return res
}

func SliceSum(slice []int) int {
	res := 0
	for _, v := range slice {
		res += v
	}
	return res
}

func SliceSumFloat(slice []float64) float64 {
	res := 0.0
	for _, v := range slice {
		res += v
	}
	return res
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