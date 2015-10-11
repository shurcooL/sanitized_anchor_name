package sanitized_anchor_name_test

import (
	"testing"

	"github.com/shurcooL/sanitized_anchor_name"
)

type convert func(string) string

func doTestsCreate(t *testing.T, tests []string, fns []convert, fnNames []string) {

	for i := 0; i < len(tests); {
		input := tests[i]
		for j, fnName := range fnNames {
			i += 1
			expected := tests[i]
			doOneTest(t, input, expected, fnName, fns[j])
		}
		i += 1
	}
}

func doOneTest(t *testing.T, input, expected, convertName string, convertFunc convert) {
	var candidate string

	candidate = input
	actual := convertFunc(input)
	if actual != expected {
		t.Errorf("Function [%#v]\nInput    [%#v]\nExpected [%#v]\nActual   [%#v]",
			convertName, candidate, expected, actual)
	}
}

func TestCreate(t *testing.T) {

	var funcCalls = []convert{
		sanitized_anchor_name.Create,
		sanitized_anchor_name.CreateGitHub,
		sanitized_anchor_name.CreateGitLab,
	}

	var funcNames = []string{
		"Create",
		"CreateGitHub",
		"CreateGitLab",
	}

	var tests = []string{

		// FORMAT:

		// Input
		// Expected-from-Create
		// Expected-from-CreateGitHub
		// Expected-from-CreateGitLab

		"This is a header",
		"this-is-a-header",
		"this-is-a-header",
		"this-is-a-header",

		"This is also          a header",
		"this-is-also-a-header",
		"this-is-also----------a-header",
		"this-is-also-a-header",

		"main.go",
		"main-go",
		"maingo",
		"maingo",

		"Article 123",
		"article-123",
		"article-123",
		"article-123",

		"<- Let's try this, shall we?",
		"let-s-try-this-shall-we",
		"--lets-try-this-shall-we",
		"-lets-try-this-shall-we",

		"        ",
		"",
		"",
		"",

		"Hello, 世界",
		"hello-世界",
		"hello-世界",
		"hello-世界",

		// Test the examples mentioned in the manual:
		// https://github.com/gitlabhq/gitlabhq/blob/master/doc/markdown/markdown.md#header-ids-and-links

		"This header has spaces in it",
		"this-header-has-spaces-in-it",
		"this-header-has-spaces-in-it",
		"this-header-has-spaces-in-it",

		"This header has a :thumbsup: in it",
		"this-header-has-a-thumbsup-in-it",
		"this-header-has-a-thumbsup-in-it",
		"this-header-has-a-thumbsup-in-it", // FIXME: Gitlab actually generates "this-header-has-a-in-it",

		"This header has Unicode in it: 한글",
		"this-header-has-unicode-in-it-한글",
		"this-header-has-unicode-in-it-한글",
		"this-header-has-unicode-in-it-한글",

		// Tests found in the wild:

		"httpunit",
		"httpunit",
		"httpunit",
		"httpunit",

		"Architecture overview",
		"architecture-overview",
		"architecture-overview",
		"architecture-overview",

		"http unit \"by hand\":",
		"http-unit-by-hand",
		"http-unit-by-hand",
		"http-unit-by-hand",

		"TOML file format:",
		"toml-file-format",
		"toml-file-format",
		"toml-file-format",

		"Basic test parameters",
		"basic-test-parameters",
		"basic-test-parameters",
		"basic-test-parameters",

		"Add a rule for a new haproxy port.",
		"add-a-rule-for-a-new-haproxy-port",
		"add-a-rule-for-a-new-haproxy-port",
		"add-a-rule-for-a-new-haproxy-port",

		"Oncall Tasks",
		"oncall-tasks",
		"oncall-tasks",
		"oncall-tasks",

		"# Header 7",
		"header-7",
		"-header-7",
		"-header-7", // FYI: This currently can't be generated.

		" Header 5",
		"header-5",
		"header-5",
		"header-5",

		" Header 6\n",
		"header-6",
		"header-6",
		"header-6",

		" Header 7\n",
		"header-7",
		"header-7",
		"header-7",

		"`-filter`",
		"filter",
		"-filter",
		"-filter",

		"`-header string`",
		"header-string",
		"-header-string",
		"-header-string",

		"`-ipmap string` foo",
		"ipmap-string-foo",
		"-ipmap-string-foo",
		"-ipmap-string-foo",

		"`-no10`",
		"no10",
		"-no10",
		"-no10",

		"`-v` and `-vv`",
		"v-and-vv",
		"-v-and--vv",
		"-v-and-vv",

		// More tests.

		"This is one",
		"this-is-one",
		"this-is-one",
		"this-is-one",

		"-This is two",
		"this-is-two",
		"-this-is-two",
		"-this-is-two",

		"--This is three",
		"this-is-three",
		"--this-is-three",
		"-this-is-three",

		"---This is four",
		"this-is-four",
		"---this-is-four",
		"-this-is-four",

		"This is -  a five",
		"this-is-a-five",
		"this-is----a-five",
		"this-is-a-five",

		"This is  - a six",
		"this-is-a-six",
		"this-is----a-six",
		"this-is-a-six",

		"-This is-   a seven",
		"this-is-a-seven",
		"-this-is----a-seven",
		"-this-is-a-seven",

		"--This is  - a eight",
		"this-is-a-eight",
		"--this-is----a-eight",
		"-this-is-a-eight",

		"---This is   -a nine",
		"this-is-a-nine",
		"---this-is----a-nine",
		"-this-is-a-nine",
	}

	groupLen := len(funcNames) + 1
	if len(tests)%groupLen != 0 {
		t.Errorf("Number of tests (%v) should be a multiple of %v", len(tests), groupLen)
	}
	if len(funcNames) != len(funcCalls) {
		t.Errorf("Number of funcCalls and funcNames must be equal: %v != %v", len(funcNames), len(funcCalls))
	}

	doTestsCreate(t, tests, funcCalls, funcNames)

}
