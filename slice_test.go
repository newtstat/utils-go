package utils_test

import (
	"testing"

	"github.com/newtstat/utils-go"
)

func TestIntSliceContains(t *testing.T) {
	testIntSlice := []int{0, 1, 2}
	t.Run("TestIntSliceContains/true", func(t *testing.T) {
		if !utils.Slice.ContainsInt(testIntSlice, 0) {
			t.Fail()
		}
	})
	t.Run("TestIntSliceContains/false", func(t *testing.T) {
		if utils.Slice.ContainsInt(testIntSlice, 3) {
			t.Fail()
		}
	})
}

func TestStringSliceContains(t *testing.T) {
	testStringSlice := []string{"0", "1", "2"}
	t.Run("TestStringSliceContains/true", func(t *testing.T) {
		if !utils.Slice.ContainsString(testStringSlice, "0") {
			t.Fail()
		}
	})
	t.Run("TestStringSliceContains/false", func(t *testing.T) {
		if utils.Slice.ContainsString(testStringSlice, "3") {
			t.Fail()
		}
	})
}
