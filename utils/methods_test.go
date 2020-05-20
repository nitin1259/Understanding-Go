package utils

import "testing"

func TestTwoElements(t *testing.T)  {
	// t.Error("no test written yer")
	result:= JoinWihtCommas([]string{"apple", "orange"})
	if result != "apple and orange"{
		t.Error("didnot match expect value")
	}
}

func TestThreeElements(t *testing.T)  {
	// t.Error("no test written yer")
	list := []string{"apple", "orange", "bananas"}
	result:= JoinWihtCommas(list)
	expected:= "apple, orange, and bananas"
	if result != expected{
		// t.Error("didnot match expect value")
		t.Errorf("JoinWihtCommas(%#v) wants %s get %s",list,expected, result )
	}
}