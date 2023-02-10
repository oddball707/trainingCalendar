package service

import (
	"fmt"
	"testing"
	"time"
	m "github.com/oddball707/trainingCalendar/model"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	assert.NotNil(t, NewService())
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
