package service

import (
	"fmt"
	"testing"
	"time"
	m "trainingCalendar/model"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	assert.NotNil(t, NewService())
}

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
			s := NewService()
			actual := s.prevMonday(day)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestStartDate(t *testing.T) {
	testCases := []struct {
		testName string
		day      string
		weeks    int
		expected string
	}{
		{
			testName: "sat20week",
			day:      "10/09/2021",
			weeks:    20,
			expected: "05/24/2021",
		},
		{
			testName: "sun12week",
			day:      "10/10/2021",
			weeks:    12,
			expected: "07/19/2021",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			day, _ := time.Parse(m.DateLayout, tc.day)
			expected, _ := time.Parse(m.DateLayout, tc.expected)
			s := NewService()
			actual := s.startDate(day, tc.weeks)
			fmt.Println(actual)
			assert.Equal(t, expected, actual)
		})
	}
}
