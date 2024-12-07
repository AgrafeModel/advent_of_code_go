package utils

import (
	"math"
	"strconv"
)

type StrNumber string

var (
	ONE   = StrNumber("one")
	TWO   = StrNumber("two")
	THREE = StrNumber("three")
	FOUR  = StrNumber("four")
	FIVE  = StrNumber("five")
	SIX   = StrNumber("six")
	SEVEN = StrNumber("seven")
	EIGHT = StrNumber("eight")
	NINE  = StrNumber("nine")
)

var Numbers = []StrNumber{ONE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE}

func (s StrNumber) String() string {
	return string(s)
}

func (s StrNumber) Number() int {
	switch s {
	case ONE:
		return 1
	case TWO:
		return 2
	case THREE:
		return 3
	case FOUR:
		return 4
	case FIVE:
		return 5
	case SIX:
		return 6
	case SEVEN:
		return 7
	case EIGHT:
		return 8
	case NINE:
		return 9
	}
	return 0
}

type Letters interface {
	string | byte | rune
}

// V1 between V2 and V3
func Between(v1, v2, v3 int) bool {
	return v1 >= v2 && v1 <= v3
}

func Dist(v1, v2 int) int {
	v1abs := math.Abs(float64(v1))
	v2abs := math.Abs(float64(v2))
	if v1abs > v2abs {
		return int(v1abs) - int(v2abs)
	}
	return int(v2abs) - int(v1abs)
}

func DirectionSign(v1, v2 int) int {
	if v1 > v2 {
		return 1
	}
	return -1
}

func ConcatInt(v ...int) int {
	//transform []int to []string
	var str []string
	for _, i := range v {
		str = append(str, strconv.Itoa(i))
	}

	//concatenate all strings
	var res string
	for _, s := range str {
		res += s
	}

	//parse int
	val, err := strconv.ParseInt(res, 10, 64)
	HandleErr(err)
	return int(val)
}

func ParseInt[T Letters](value T) int {
	str := string(value)
	val, err := strconv.ParseInt(str, 10, 64)
	HandleErr(err)
	return int(val)
}

func ParseSliceInt[T Letters](value []T) []int {
	var res []int
	for _, v := range value {
		res = append(res, ParseInt(v))
	}
	return res
}

func IsInt(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
