package server

import "testing"

type getUserIDTestCase struct {
	session  string
	expected int
}

var getUserIDTestCases = []getUserIDTestCase{
	{
		session:  "12",
		expected: 12,
	},
	{
		session:  "234234234234",
		expected: 0,
	},
	{
		session:  "2325675687993",
		expected: -1,
	},
}

func Test_getUserID(t *testing.T) {
	ConnectDatabase()
	for _, testCase := range getUserIDTestCases {
		result := getUserID(testCase.session)
		if result != testCase.expected {
			t.Errorf("Expected %v, got %v", testCase.expected, result)
		}
	}
}

type getSubjectsTestCase struct {
	userID   int
	expected []Subject
}

var getSubjectsTestCases = []getSubjectsTestCase{
	{
		userID: 1234,
		expected: []Subject{
			{
				SubjectID:          1731025092,
				SubjectName:        "Biology",
				SubjectDescription: "Subject especialy to masochistic students. Subject about human anathomy.",
			},
			{
				SubjectID:          3,
				SubjectName:        "Mathematics",
				SubjectDescription: "The basics of mathematics with extended probability theory.",
			},
			{
				SubjectID:          1,
				SubjectName:        "Calculus",
				SubjectDescription: "Calculus theory with practical examples.",
			},
		},
	},
}

func Test_getSubjects(t *testing.T) {
	ConnectDatabase()
	for _, testCase := range getSubjectsTestCases {
		result, _ := getSubjects(testCase.userID)
		if len(result) != len(testCase.expected) {
			t.Errorf("Expected %v, got %v", testCase.expected, result)
		}
		for i := 0; i < len(result); i++ {
			if (result[i].SubjectID != testCase.expected[i].SubjectID) ||
				(result[i].SubjectName != testCase.expected[i].SubjectName) ||
				(result[i].SubjectDescription != testCase.expected[i].SubjectDescription) {
				t.Errorf("Expected %v, got %v", testCase.expected, result)
			}
		}
	}
}
