package utils

import (
	"fmt"
	"testing"
)

func TestTwoElements(t *testing.T) {
	// t.Error("no test written yer")
	list := []string{"apple", "orange"}
	result := JoinWihtCommas(list)
	expected := "apple and orange"
	if result != expected {
		t.Errorf("JoinWihtCommas(%#v) %s", list, errorString(expected, result))
	}
}

func TestThreeElements(t *testing.T) {
	// t.Error("no test written yer")
	list := []string{"apple", "orange", "bananas"}
	result := JoinWihtCommas(list)
	expected := "apple, orange, and bananas"
	if result != expected {
		// t.Error("didnot match expect value")
		t.Errorf("JoinWihtCommas(%#v) %s", list, errorString(expected, result))
	}
}

func TestOneElement(t *testing.T) {
	// t.Error("no test written yer")
	list := []string{"apple"}
	result := JoinWihtCommas(list)
	expected := "apple"
	if result != expected {
		t.Errorf("JoinWihtCommas(%#v) %s", list, errorString(expected, result))
	}
}

// You can reduce repeated code in your tests by moving it to other “helper” functions within your test file.
//  The go test command only uses functions whose names begin with Test, so as long as you name your functions anything else, you’ll be fine.

func errorString(expected string, result string) string {
	return fmt.Sprintf("expecting: %s, getting: %s", expected, result)
}

// -v flag, which stands for “verbose.” If you add it to any go test command, it will list the name and status of each test function it runs. Normally passing tests are omitted to keep the output “quiet,” but in verbose mode, go test will list even passing tests.
// go test ./utils -v

// you can add the -run option to limit the set of tests that are run. Following -run, you specify part or all of a function name, and only test functions whose name matches what you specify will be run.
// go test ./utils -v -run Elementsf
//both TestTwoElements and TestThreeElements will be run. (But not TestOneElement, because it doesn’t have an s at the end of its name.f

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {

	tests := []testData{
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "bananas"}, want: "apple, orange, and bananas"},
		testData{list: []string{"apple"}, want: "apple"},
	}

	for _, test := range tests {
		got := JoinWihtCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWihtCommas(%#v) %s", test.list, errorString(test.want, got))
		}
	}
}
