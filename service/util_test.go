package service

import (
	"testing"
	"time"
	m "github.com/oddball707/trainingCalendar/model"

	"github.com/stretchr/testify/assert"
)

func TestPrevMonday(t *testing.T) {
	testCases := []struct {
		testName string
		day      string
		expected string
	}{
		{
			testName: "monday",
			day:      "10/04/2021",
			expected: "10/04/2021",
		},
		{
			testName: "tuesday",
			day:      "10/05/2021",
			expected: "10/04/2021",
		},
		{
			testName: "wednesday",
			day:      "10/06/2021",
			expected: "10/04/2021",
		},
		{
			testName: "thursday",
			day:      "10/07/2021",
			expected: "10/04/2021",
		},
		{
			testName: "friday",
			day:      "10/08/2021",
			expected: "10/04/2021",
		},
		{
			testName: "saturday",
			day:      "10/09/2021",
			expected: "10/04/2021",
		},
		{
			testName: "sunday",
			day:      "10/10/2021",
			expected: "10/04/2021",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			day, _ := time.Parse(m.DateLayout, tc.day)
			expected, _ := time.Parse(m.DateLayout, tc.expected)
			actual := PrevMonday(day)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestNextMonday(t *testing.T) {
	testCases := []struct {
		testName string
		day      string
		expected string
	}{
		{
			testName: "monday",
			day:      "07/11/2022",
			expected: "07/18/2022",
		},
		{
			testName: "tuesday",
			day:      "07/12/2022",
			expected: "07/18/2022",
		},
		{
			testName: "wednesday",
			day:      "07/13/2022",
			expected: "07/18/2022",
		},
		{
			testName: "thursday",
			day:      "07/14/2022",
			expected: "07/18/2022",
		},
		{
			testName: "friday",
			day:      "07/15/2022",
			expected: "07/18/2022",
		},
		{
			testName: "saturday",
			day:      "07/16/2022",
			expected: "07/18/2022",
		},
		{
			testName: "sunday",
			day:      "07/17/2022",
			expected: "07/18/2022",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			day, _ := time.Parse(m.DateLayout, tc.day)
			expected, _ := time.Parse(m.DateLayout, tc.expected)
			actual := NextMonday(day)
			assert.Equal(t, expected, actual)
		})
	}
}