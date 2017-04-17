package sanitized_anchor_name_test

import (
	"fmt"
	"testing"

	"github.com/shurcooL/sanitized_anchor_name"
)

func ExampleCreate() {
	anchorName := sanitized_anchor_name.Create("This is a header")

	fmt.Println(anchorName)

	// Output:
	// this-is-a-header
}

func ExampleCreate_two() {
	fmt.Println(sanitized_anchor_name.Create("This is a header"))
	fmt.Println(sanitized_anchor_name.Create("This is also          a header"))
	fmt.Println(sanitized_anchor_name.Create("main.go"))
	fmt.Println(sanitized_anchor_name.Create("Article 123"))
	fmt.Println(sanitized_anchor_name.Create("<- Let's try this, shall we?"))
	fmt.Printf("%q\n", sanitized_anchor_name.Create("        "))
	fmt.Println(sanitized_anchor_name.Create("Hello, 世界"))

	// Output:
	// this-is-a-header
	// this-is-also-a-header
	// main-go
	// article-123
	// let-s-try-this-shall-we
	// ""
	// hello-世界
}

// More tests. This uses a framework that is easier to extend.

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
	}

	var funcNames = []string{
		"Create",
	}

	var tests = []string{

		// FORMAT:

		// Input
		// Expected-from-Create

		"This is a header",
		"this-is-a-header",

		"This is also          a header",
		"this-is-also-a-header",

		"main.go",
		"main-go",

		"Article 123",
		"article-123",

		"<- Let's try this, shall we?",
		"let-s-try-this-shall-we",

		"        ",
		"",

		"Hello, 世界",
		"hello-世界",

		// Test the examples mentioned in the manual:
		// https://github.com/gitlabhq/gitlabhq/blob/master/doc/markdown/markdown.md#header-ids-and-links

		"This header has spaces in it",
		"this-header-has-spaces-in-it",

		"This header has a :thumbsup: in it",
		"this-header-has-a-thumbsup-in-it",

		"This header has Unicode in it: 한글",
		"this-header-has-unicode-in-it-한글",

		// Tests found in the wild:

		"httpunit",
		"httpunit",

		"Architecture overview",
		"architecture-overview",

		"http unit \"by hand\":",
		"http-unit-by-hand",

		"TOML file format:",
		"toml-file-format",

		"Basic test parameters",
		"basic-test-parameters",

		"Add a rule for a new haproxy port.",
		"add-a-rule-for-a-new-haproxy-port",

		"Oncall Tasks",
		"oncall-tasks",

		"# Header 7",
		"header-7",

		" Header 5",
		"header-5",

		" Header 6\n",
		"header-6",

		" Header 7\n",
		"header-7",

		"`-filter`",
		"filter",

		"`-header string`",
		"header-string",

		"`-ipmap string` foo",
		"ipmap-string-foo",

		"`-no10`",
		"no10",

		"`-v` and `-vv`",
		"v-and-vv",

		// More tests.

		"This is one",
		"this-is-one",

		"-This is two",
		"this-is-two",

		"--This is three",
		"this-is-three",

		"---This is four",
		"this-is-four",

		"This is -  a five",
		"this-is-a-five",

		"This is  - a six",
		"this-is-a-six",

		"-This is-   a seven",
		"this-is-a-seven",

		"--This is  - a eight",
		"this-is-a-eight",

		"---This is   -a nine",
		"this-is-a-nine",
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
