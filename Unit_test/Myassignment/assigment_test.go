package assignment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	testCases := []struct {
		first_value, second_value, sum uint32
		MyResult                       bool
	}{
		{
			first_value: 1, second_value: 1, sum: 2, MyResult: false,
		},
		{
			first_value: 4294967290, second_value: 6, sum: 0, MyResult: true,
		},
	}

	for i, tC := range testCases {
		fsum, fResult := AddUint32(testCases[i].first_value, testCases[i].second_value)
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tC.sum, fsum)
			assert.Equal(t, tC.MyResult, fResult)
		})
	}
}

func TestCeilNumber(t *testing.T) {
	testCases := []struct {
		value, MyResult float64
	}{
		{
			value: 42.42, MyResult: 42.50,
		},
		{
			value: 42.99, MyResult: 43,
		},
		{
			value: 42.55, MyResult: 42.75,
		},
		{
			value: 43.13, MyResult: 43.25,
		},
		{
			value: 42.75, MyResult: 42.75,
		},
	}

	for i, tC := range testCases {
		fResult := CeilNumber(testCases[i].value)
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tC.MyResult, fResult)
		})
	}
}

func TestAlphabetSoup(t *testing.T) {
	testCases := []struct {
		value, MyResult string
	}{
		{
			value: "hello", MyResult: "ehllo",
		},
		{
			value: "bac", MyResult: "abc",
		},
		{
			value: "ab", MyResult: "ab",
		},
	}

	for i, tC := range testCases {
		fResult := AlphabetSoup(testCases[i].value)
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tC.MyResult, fResult)
		})
	}
}

func TestStringMask(t *testing.T) {
	testCases := []struct {
		str, MyResult string
		value         uint
	}{
		{
			str: "a", value: 1, MyResult: "*",
		},
		{
			str: "string", value: 3, MyResult: "str***",
		},
		{
			str: "string", value: 11, MyResult: "******",
		},
		{
			str: "", value: 5, MyResult: "*",
		},
	}

	for i, tC := range testCases {
		fResult := StringMask(testCases[i].str, testCases[i].value)
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tC.MyResult, fResult)
		})
	}
}

func TestString_Mask_MakeStar(t *testing.T) {
	tstring := String_Mask_MakeStar(3)
	assert.Equal(t, "***", tstring)
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"

	testCases := []struct {
		MyResult     string
		first_String string
	}{
		{first_String: "hellocat", MyResult: "hello,cat"},
		{first_String: "catbat", MyResult: "cat,bat"},
		{first_String: "notcat", MyResult: "not possible"},
		{first_String: "bootcamprocks!", MyResult: "not possible"},
		{first_String: "", MyResult: "not possible"},
	}
	var strArray [2]string
	strArray[1] = words
	for _, tC := range testCases {
		strArray[0] = tC.first_String
		fResult := WordSplit(strArray)
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tC.MyResult, fResult)
		})
	}
}

func TestWordSplitControlling(t *testing.T) {
	strArray := []string{"Istanbul", "Ankara", "Mersin", "Adana"}
	result := WordSplitControlling("Mersin", strArray)
	assert.Equal(t, true, result)
}

func TestVariadicSet(t *testing.T) {

	testCases := []struct {
		value    []interface{}
		MyResult []interface{}
	}{
		{
			value:    []interface{}{"bootcamp", "rocks!", "really", "rocks!"},
			MyResult: []interface{}{"bootcamp", "rocks!", "really"},
		},
		{
			value:    []interface{}{1, uint32(1), "first", 2, uint32(2), "second!"},
			MyResult: []interface{}{1, uint32(1), "first", 2, uint32(2), "second!"},
		},
		{
			value:    []interface{}{4, 2, 5, 4, 2, 4},
			MyResult: []interface{}{4, 2, 5},
		},
	}

	for i, tC := range testCases {
		fResult := VariadicSet(testCases[i].value...)
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tC.MyResult, fResult)
		})
	}
}
