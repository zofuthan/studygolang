package util_test

import (
	"testing"
	. "util"
)

type model struct {
	Id   int
	Name string
}

func TestModels2Intslice(t *testing.T) {
	models := []*model{
		&model{12, "polaris"},
		&model{13, "xuxinhua"},
	}

	actualResult := Models2Intslice(models, "Id")
	expectResult := []int{12, 13}

	if !sliceIsEqual(actualResult, expectResult) {
		t.Fatalf("expect:%v, actual:%v", expectResult, actualResult)
	}
}

func sliceIsEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}

	return true
}
