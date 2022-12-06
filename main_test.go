package main

import (
	"testing"
	"time"
)

func TestCalc(t *testing.T) {
	dateLayout := "2006-01-02 15:04"

	tests := []struct {
		name             string
		waitUntil        string
		relativeTo       string
		expectedTime     string
		expectedDuration time.Duration
	}{
		// h
		{
			name:             "wu 9",
			waitUntil:        "9",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-07 09:00",
			expectedDuration: 23*time.Hour + 45*time.Minute,
		},
		{
			name:             "wu 10",
			waitUntil:        "10",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-06 10:00",
			expectedDuration: 45 * time.Minute,
		},
		{
			name:             "wu 16",
			waitUntil:        "16",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-06 16:00",
			expectedDuration: 6*time.Hour + 45*time.Minute,
		},
		{
			name:             "wu 24",
			waitUntil:        "24",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-07 00:00",
			expectedDuration: 14*time.Hour + 45*time.Minute,
		},
		{
			name:             "wu 25",
			waitUntil:        "25",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-07 01:00",
			expectedDuration: 15*time.Hour + 45*time.Minute,
		},
		{
			name:             "wu -2",
			waitUntil:        "-2",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-06 22:00",
			expectedDuration: 12*time.Hour + 45*time.Minute,
		},
		// h:m
		{
			name:             "wu 9:10",
			waitUntil:        "9:10",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-07 09:10",
			expectedDuration: 23*time.Hour + 55*time.Minute,
		},
		{
			name:             "wu 9:20",
			waitUntil:        "9:20",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-06 09:20",
			expectedDuration: 5 * time.Minute,
		},
		{
			name:             "wu 16:30",
			waitUntil:        "16:30",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-06 16:30",
			expectedDuration: 7*time.Hour + 15*time.Minute,
		},
		{
			name:             "wu 26:90",
			waitUntil:        "26:90",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-07 03:30",
			expectedDuration: 18*time.Hour + 15*time.Minute,
		},
		{
			name:             "wu 24:-5",
			waitUntil:        "24:-5",
			relativeTo:       "2022-12-06 09:15",
			expectedTime:     "2022-12-06 23:55",
			expectedDuration: 14*time.Hour + 40*time.Minute,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			relativeTo, err := time.Parse(dateLayout, test.relativeTo)
			if err != nil {
				t.Error(err)
			}

			actualTime, actualDuration := calc(test.waitUntil, relativeTo)
			if test.expectedTime != actualTime.Format(dateLayout) {
				t.Errorf("Calculated time does not match:\n%v\nbut got\n%v", test.expectedTime, actualTime)
			}
			if test.expectedDuration != actualDuration {
				t.Errorf("Calculated duration does not match:\n%v\nbut got\n%v", test.expectedDuration, actualDuration)
			}
		})
	}
}
