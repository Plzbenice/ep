package ep

import (
	"testing"
)

func TestGetDateRange(t *testing.T) {
	expect1 := []string{"2000-01-01", "2000-01-02", "2000-01-03", "2000-01-04", "2000-01-05"}
	res1, err := GetDateRange("2000-01-01", "2000-01-05")
	if err != nil {
		t.Error("TestGetDateList Failed", err)
		return
	}
	if !testEq(res1, expect1) {
		t.Error("TestGetDateList Failed")
	}

	expect2 := []string{"2000-01-05", "2000-01-04", "2000-01-03", "2000-01-02", "2000-01-01"}
	res2, err := GetDateRange("2000-01-05", "2000-01-01")
	if err != nil {
		t.Error("TestGetDateList Failed", err)
		return
	}
	if !testEq(res2, expect2) {
		t.Error("TestGetDateList Failed")
	}
}

//test two []string is equal
func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
