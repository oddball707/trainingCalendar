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

func TestFloatToPace(t *testing.T) {
	testCases := []struct {
		testName string
		pace     float64
		expected string
	}{
		{
			testName: "exact minutes",
			pace:     5.0,
			expected: "5:00",
		},
		{
			testName: "minutes and seconds",
			pace:     5.5,
			expected: "5:30",
		},
		{
			testName: "minutes and seconds with rounding",
			pace:     5.79,
			expected: "5:47",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actual := FloatToPace(tc.pace)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPaceToFloat(t *testing.T) {
	testCases := []struct {
		testName string
		pace     string
		expected float64
	}{
		{
			testName: "exact minutes",
			pace:     "5:00",
			expected: 5.0,
		},
		{
			testName: "minutes and seconds",
			pace:     "5:30",
			expected: 5.5,
		},
		{
			testName: "minutes and seconds with rounding",
			pace:     "59:47",
			expected: 59.78333333333333,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actual, err := PaceToFloat(tc.pace)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSetDescription(t *testing.T) {
	testCases := []struct {
		testName     string
		desc         string
		raceLength   float64
		goalTime     string
		expectedDesc string
	}{
		{
			testName:     "no placeholder",
			desc:         "Run 5 miles",
			raceLength:   5,
			goalTime:     "25:00",
			expectedDesc: "Run 5 miles",
		},
		{
			testName:     "with placeholder",
			desc:         "Run 5 miles at <rp> pace",
			raceLength:   5,
			goalTime:     "25:00",
			expectedDesc: "Run 5 miles at 5:00 pace",
		},
		{
			testName:     "with placeholder and modifier",
			desc:         "5x600m@<112%rp>;2min recovery",
			raceLength:   3.1,
			goalTime:     "20:00",
			expectedDesc: "5x600m@5:45;2min recovery",
		},
		{
			testName:     "with double placeholder and modifier",
			desc:         "60min@<70%rp>-<65%rp>",
			raceLength:   3.1,
			goalTime:     "20:00",
			expectedDesc: "60min@9:12-9:55",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actual := SetDescription(tc.desc, tc.raceLength, tc.goalTime)
			assert.Equal(t, tc.expectedDesc, actual)
		})
	}
}

func TestParseSpeed(t *testing.T) {
	testCases := []struct {
		testName       string
		desc           string
		expectedSpeeds []int
	}{
		{
			testName:       "no placeholder",
			desc:           "Run 5 miles",
			expectedSpeeds: []int{},
		},
		{
			testName:       "with placeholder",
			desc:           "Run 5 miles at <100%rp> pace",
			expectedSpeeds: []int{100},
		},
		{
			testName:       "with placeholder and extra text",
			desc:           "Run 5 miles@<10%rp> pace for 30 minutes",
			expectedSpeeds: []int{10},
		},
		{
			testName:       "with double placeholder and modifier",
			desc:           "60min@<65%rp>-<75%rp>",
			expectedSpeeds: []int{65, 75},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			speed := ParseSpeed(tc.desc)
			assert.Equal(t, tc.expectedSpeeds, speed)
		})
	}
}
