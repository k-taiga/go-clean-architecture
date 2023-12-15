package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3

	if result != expected {
		t.Errorf("Add(1, 2) = %v, want %v", result, expected)
	}
}

type divideTestCase struct {
	a, b           int
	expectedResult int
	expectErr      bool
}

func TestDivide(t *testing.T) {
	// テーブルドリブンテスト
	testCases := []divideTestCase{
		{a: 6, b: 3, expectedResult: 2, expectErr: false},
		{a: 6, b: 0, expectedResult: 0, expectErr: true},
	}

	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)
		if tc.expectErr {
			assert.Error(t, err, "an expected error occurred: %v", err)
		} else {
			assert.NoError(t, err, "an expected error occurred: %v", err)
			assert.Equal(t, tc.expectedResult, result, "they should be equal")
		}
	}
}
