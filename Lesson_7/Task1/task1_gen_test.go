// CODE GENERATED AUTOMATICALLY
// THIS FILE SHOULD NOT BE EDITED BY HAND
package main

import (
	"testing"
)

type TestStruct struct {
	m        map[string]interface{}
	expected MyStruct
}

func TestSetStructValuesSuccess(t *testing.T) {
	tests := []TestStruct{

		TestStruct{
			m: map[string]interface{}{
				"I": 8752868801788864235,
				"Y": false,
			},
			expected: MyStruct{
				8752868801788864235,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 4682783784757851203,
				"Y": false,
			},
			expected: MyStruct{
				4682783784757851203,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 8302132053310701207,
				"Y": false,
			},
			expected: MyStruct{
				8302132053310701207,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 2695359068405206371,
				"Y": true,
			},
			expected: MyStruct{
				2695359068405206371,
				true,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 4111710841820132141,
				"Y": true,
			},
			expected: MyStruct{
				4111710841820132141,
				true,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 2479346653986449934,
				"Y": false,
			},
			expected: MyStruct{
				2479346653986449934,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 6125847403189666898,
				"Y": true,
			},
			expected: MyStruct{
				6125847403189666898,
				true,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 555161461305148190,
				"Y": false,
			},
			expected: MyStruct{
				555161461305148190,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 7542536434660668278,
				"Y": false,
			},
			expected: MyStruct{
				7542536434660668278,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 6215469264505523913,
				"Y": false,
			},
			expected: MyStruct{
				6215469264505523913,
				false,
			},
		},
	}

	ts := NewMyStruct()
	for _, v := range tests {
		err := SetStructValues(ts, v.m)
		if err != nil {
			t.Errorf("Ошибка «%v»\n", err)
		}

		if ts.I != v.expected.I {
			t.Errorf("Ожидалось: «%v». Пришло: «%v»\n", v.expected.I, ts.I)
		}
		if ts.Y != v.expected.Y {
			t.Errorf("Ожидалось: «%v». Пришло: «%v»\n", v.expected.Y, ts.Y)
		}
	}
}

func TestSetStructValuesIntFailed(t *testing.T) {
	tests := []TestStruct{

		TestStruct{
			m: map[string]interface{}{
				"I": 8752868801788864236,
				"Y": false,
			},
			expected: MyStruct{
				8752868801788864235,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 4682783784757851204,
				"Y": false,
			},
			expected: MyStruct{
				4682783784757851203,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 8302132053310701208,
				"Y": false,
			},
			expected: MyStruct{
				8302132053310701207,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 2479346653986449935,
				"Y": false,
			},
			expected: MyStruct{
				2479346653986449934,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 555161461305148191,
				"Y": false,
			},
			expected: MyStruct{
				555161461305148190,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 7542536434660668279,
				"Y": false,
			},
			expected: MyStruct{
				7542536434660668278,
				false,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 6215469264505523914,
				"Y": false,
			},
			expected: MyStruct{
				6215469264505523913,
				false,
			},
		},
	}

	ts := NewMyStruct()
	for _, v := range tests {
		err := SetStructValues(ts, v.m)
		if err != nil {
			t.Errorf("Ошибка «%v»\n", err)
		}
		if ts.I == v.expected.I {
			t.Errorf("Ожидалось: «%v». Пришло: «%v»", v.expected.I, ts.I)
		}
	}
}

func TestSetStructValuesBoolFailed(t *testing.T) {
	tests := []TestStruct{

		TestStruct{
			m: map[string]interface{}{
				"I": 2695359068405206371,
				"Y": false,
			},
			expected: MyStruct{
				2695359068405206371,
				true,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 4111710841820132141,
				"Y": false,
			},
			expected: MyStruct{
				4111710841820132141,
				true,
			},
		},

		TestStruct{
			m: map[string]interface{}{
				"I": 6125847403189666898,
				"Y": false,
			},
			expected: MyStruct{
				6125847403189666898,
				true,
			},
		},
	}

	ts := NewMyStruct()
	for _, v := range tests {
		err := SetStructValues(ts, v.m)
		if err != nil {
			t.Errorf("Ошибка «%v»\n", err)
		}
		if ts.Y == v.expected.Y {
			t.Errorf("Ожидалось: «%v». Пришло: «%v»", v.expected.Y, ts.Y)
		}
	}
}
