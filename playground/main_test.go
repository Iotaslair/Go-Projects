package main

import "testing"

func TestLinkedIn(t *testing.T) {
	testCases := []struct {
		inputEmail     string
		expectedOutput bool
	}{
		{
			"william@linkedin.com",
			true,
		},
		{
			"william@google.com",
			false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.inputEmail, func(t *testing.T) {
			actualOutput := IsLinkedInEmployee(testCase.inputEmail)
			if actualOutput != testCase.expectedOutput {
				t.Error("Doesn't match")
			}
		})
	}
}
