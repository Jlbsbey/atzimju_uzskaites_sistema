package server

import (
	"testing"
)

type HashPasswordTestCase struct {
	password string
	salt     string
	expected string
}

var hashPasswordTestCases = []HashPasswordTestCase{
	{
		password: "myVerySimplePasswordHere",
		salt:     "dadaf67409828e4ee168e1b667d9c013e524124576265c10231c3e18e66fa78c",
		expected: "cf7babea87fe5c2d5a18f7607eea5ee1ccb36202b2fcbc09e5469eea1dd09110",
	},
	{
		password: "Hfjkd4.fio(*fdskfek",
		salt:     "80a2f915ad8abe0be868acda547f7a641cb7a0701385607695dd7e6de44233d8",
		expected: "80f012f2ce1e26f78932287c6c20eb9c71a94598a8b083fccfdfcc99fef80552",
	},
	{
		password: "admin",
		salt:     "3f04fdfe7e5b87f178dcd90751effbe592e720b304358913d2228bdd0d8143a3",
		expected: "62161205713f1d8315d2844029d0388577cec6b57dd1d82eac4ab57435b95054",
	},
	{
		password: "&jJF.d(jj39fKJfd93..f",
		salt:     "3f04fdfe7e5b87f178dcd90751effbe592e720b304358913d2228bdd0d8143a3",
		expected: "22874ba6ef41da0fda2072ce3bd3d734b7fe42e08006da67085a50ae684117ee",
	},
	{
		password: "dontLookAt773mostUsedPassword",
		salt:     "3f04fdfe7e5b87f178dcd90751effbe592e720b304358913d2228bdd0d8143a3",
		expected: "4b2b9285cf357bfb8d3387011f701abe02a448144cd04a659216fb7cca57b705",
	},
}

func Test_hashPassword(t *testing.T) {
	for _, testCase := range hashPasswordTestCases {
		var got = hashPassword(testCase.password, testCase.salt)
		if got != testCase.expected {
			t.Errorf("got %q, wanted %q", got, testCase.expected)
		}
	}
}

type TryLoginTestCase struct {
	username           string
	hashedPassword     string
	salt               string
	expectedIsLoggedIn bool
	expectedUserId     int
}

var tryLoginTestCases = []TryLoginTestCase{
	{
		username:           "admin",
		hashedPassword:     "f51eb0b2f465c2ea64e8ff3b6ba96d6ff0e299c6645c367d6bde892527b6c493",
		expectedIsLoggedIn: true,
		expectedUserId:     12,
	},
}

func Test_tryLogin(t *testing.T) {
	for _, testCase := range tryLoginTestCases {
		var gotIsLoggedIn, gotUserId = tryLogin(testCase.username, testCase.hashedPassword)
		if (gotIsLoggedIn != testCase.expectedIsLoggedIn) &&
			(gotUserId != testCase.expectedUserId) {
			t.Errorf(
				"got %t and %q, expected %t and %q",
				gotIsLoggedIn, gotUserId, testCase.expectedIsLoggedIn, testCase.expectedUserId,
			)
		}
	}

}
