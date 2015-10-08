package sanitized_anchor_name_test

import (
	"fmt"

	"github.com/shurcooL/sanitized_anchor_name"
)

func ExampleCreate() {
	anchorName := sanitized_anchor_name.Create("This is a header")

	fmt.Println(anchorName)

	// Output:
	// this-is-a-header
}

func ExampleCreate2() {
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
	// maingo
	// article-123
	// -lets-try-this-shall-we
	// ""
	// hello-世界
}

func ExampleCreateGitlab() {

	// Test the examples mentioned in the manual:
	// https://github.com/gitlabhq/gitlabhq/blob/master/doc/markdown/markdown.md#header-ids-and-links

	var tests = []string{
		"This header has spaces in it",
		//"This header has a :thumbsup: in it",
		"This header has Unicode in it: 한글",
		"This header has spaces in it",
		//"This header has spaces in it",
	}

	for _, t := range tests {
		fmt.Println(sanitized_anchor_name.Create(t))
	}

	// Output:
	// this-header-has-spaces-in-it
	// this-header-has-unicode-in-it-한글
	// this-header-has-spaces-in-it
}

func ExampleCreate3() {

	// Test odd cases found in wiki documents found in the wild.

	var tests = []string{
		"httpunit",
		"Architecture overview",
		"http unit \"by hand\":",
		"TOML file format:",
		"Basic test parameters",
		"IP variables",
		"Add a rule for a new haproxy port.",
		"Oncall Tasks",
	}

	for _, t := range tests {
		fmt.Println(sanitized_anchor_name.Create(t))
	}

	// Output:
	// httpunit
	// architecture-overview
	// http-unit-by-hand
	// toml-file-format
	// basic-test-parameters
	// ip-variables
	// add-a-rule-for-a-new-haproxy-port
	// oncall-tasks
}

func ExampleCreate4() {

	// Edge cases related to dashes and code.

	var tests = []string{
		"# Header 7",
		" Header 5",
		" Header 6\n",
		" Header 7\n",
		"`-filter`",
		"`-header string`",
		"`-ipmap string` foo",
		"`-no10`",
		"`-v` and `-vv`",
	}

	for _, t := range tests {
		fmt.Println(sanitized_anchor_name.Create(t))
	}

	// Output:
	// -header-7
	// header-5
	// header-6
	// header-7
	// -filter
	// -header-string
	// -ipmap-string-foo
	// -no10
	// -v-and-vv
}

func ExampleCreate5() {

	// Edge cases related to dashes and spaces.

	var tests = []string{
		"This is one",
		"-This is two",
		"--This is three",
		"---This is four",
		"This is -  a five",
		"This is  - a six",
		"-This is-   a seven",
		"--This is  - a eight",
		"---This is   -a nine",
	}

	for _, t := range tests {
		fmt.Println(sanitized_anchor_name.Create(t))
	}

	// Output:
	// this-is-one
	// -this-is-two
	// -this-is-three
	// -this-is-four
	// this-is-a-five
	// this-is-a-six
	// -this-is-a-seven
	// -this-is-a-eight
	// -this-is-a-nine
}
