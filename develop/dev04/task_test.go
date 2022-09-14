package main

import (
	"reflect"
	"testing"
)

func TestGetAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    *[]string
		expected *map[string][]string
	}{
		{
			"common case",
			&[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			&map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"тяпка", "пятак", "пятка"},
			},
		},
		{
			"common case with uppercase letters",
			&[]string{"ПяТак", "пятКа", "тЯпка", "ЛИСТОК", "СЛИТок", "столИК"},
			&map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"тяпка", "пятак", "пятка"},
			},
		},
		{
			"no anagrams with empty output",
			&[]string{"пятак", "листок"},
			&map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)
		reflect.DeepEqual(*tt.expected, *GetAnagrams(tt.input))
	}
}
