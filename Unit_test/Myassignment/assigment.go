package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {

	if x > math.MaxUint32-y {
		return x + y, true
	}

	return x + y, false
}

func CeilNumber(f float64) float64 {
	num1 := int(f) // Convert float to int

	kesir := f - float64(num1) //Find value after comma

	if int(100*kesir) == 0 {
		return float64(num1)
	} else if int(100*kesir) > 0 && int(100*kesir) <= 25 {
		return float64(num1) + 0.25
	} else if int(100*kesir) > 25 && int(100*kesir) <= 50 {
		return float64(num1) + 0.50
	} else if int(100*kesir) > 50 && int(100*kesir) <= 75 {
		return float64(num1) + 0.75
	} else {
		return float64(num1) + 1.0
	}
}

func AlphabetSoup(s string) string {
	w := strings.Split(s, "")
	sort.Strings(w)

	/*char := []rune(s) //Convert String to array

	sort.Slice(char, func(i, j int) bool {
		return char[j] > char[i]
	})
	str := string(char[:]) //Convert slice to string*/

	return strings.Join(w, "")
}

func StringMask(s string, n uint) string {
	len := len(s)     // in order to convert slice , length has been taken .
	ulen := uint(len) // Convert string length to uint

	w := []rune(s) // string to slice

	if n >= ulen { // Controlling n paramter is bigger than or equal length of string
		NewStr := String_Mask_MakeStar(ulen)
		return NewStr
	} else {
		NewStr := string(w[0:n]) + String_Mask_MakeStar(ulen-n) // Merging stars and the rest of string
		return NewStr
	}

}
func String_Mask_MakeStar(v uint) string {
	var str string // ****

	if v == 0 { // Controlling empty string
		return "*"
	}

	var i uint
	for i = 0; i < v; i++ { // Creating **
		str += "*"
	}

	return str
}

func WordSplit(arr [2]string) string {
	stArray := strings.Split(arr[1], ",")
	var str1, str2, newStr string

	len := len(arr[0])
	String_Chars := strings.Split(arr[0], "")

	for i := 0; i < len; i++ {
		newStr += String_Chars[i]
		if c := WordSplitControlling(newStr, stArray); c {

			if str1 == "" {
				str1 = newStr
			} else {
				str2 = newStr
			}
			newStr = ""
		}
	}
	if str1 == "" || str2 == "" {
		return "not possible"
	}
	newStr = str1 + "," + str2

	return newStr
}
func WordSplitControlling(s string, s_arr []string) bool {
	for _, v := range s_arr {
		if v == s {
			return true
		}
	}

	return false
}

func VariadicSet(i ...interface{}) []interface{} {
	var ResultI []interface{}
	ResultI = append(ResultI, i[0])

	for t := 1; t < len(i); t++ {
		if c := VariadicSetControl(i[t], ResultI); !c {
			ResultI = append(ResultI, i[t])
		}
	}

	return ResultI
}

func VariadicSetControl(i interface{}, t []interface{}) bool {

	for _, v := range t {
		if v == i {
			return true
		}
	}

	return false
}
