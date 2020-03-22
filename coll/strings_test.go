package coll

import (
	"strings"
	"testing"
)

func TestStringIndex(t *testing.T) {
	tables := []struct {
		arr      []string
		token    string
		expected int
	}{
		{[]string{"foo"}, "bar", -1},
		{[]string{"foo", "bar"}, "bar", 1},
		{[]string{"foobar"}, "bar", -1},
		{[]string{"bar", "foo", "bar"}, "bar", 0},
		{[]string{"bar"}, "bar", 0},
	}

	for _, table := range tables {
		output := StringIndex(table.arr, table.token)
		if output != table.expected {
			t.Errorf("Arr (%v), token (%v) was incorrect, got: %v, want: %v.", table.arr, table.token, output, table.expected)
		}
	}
}

func TestInclude(t *testing.T) {
	tables := []struct {
		arr      []string
		token    string
		expected bool
	}{
		{[]string{"foo"}, "bar", false},
		{[]string{"foo", "bar"}, "bar", true},
		{[]string{"foobar"}, "bar", false},
		{[]string{"bar", "foo", "bar"}, "bar", true},
		{[]string{"bar"}, "bar", true},
	}

	for _, table := range tables {
		output := Include(table.arr, table.token)
		if output != table.expected {
			t.Errorf("Arr (%v), token (%v) was incorrect, got: %v, want: %v.", table.arr, table.token, output, table.expected)
		}
	}
}

func TestAny(t *testing.T) {
	tables := []struct {
		arr      []string
		f        func(string) bool
		expected bool
	}{
		{[]string{"foo"}, func(token string) bool { return len(token) > 10 }, false},
		{[]string{"foo", "bar"}, func(token string) bool { return len(token) > 3 }, false},
		{[]string{"foobar"}, func(token string) bool { return len(token) > 3 }, true},
		{[]string{"bar", "foo", "bar", "foobar"}, func(token string) bool { return len(token) >= 4 }, true},
		{[]string{"bar"}, func(token string) bool { return len(token) < 1 }, false},
	}

	for _, table := range tables {
		output := Any(table.arr, table.f)
		if output != table.expected {
			t.Errorf("Incorrect, got: %v, want: %v.", output, table.expected)
		}
	}
}

func TestAll(t *testing.T) {
	tables := []struct {
		arr      []string
		f        func(string) bool
		expected bool
	}{
		{[]string{"foo"}, func(token string) bool { return len(token) > 10 }, false},
		{[]string{"foo", "bar"}, func(token string) bool { return len(token) > 3 }, false},
		{[]string{"foobar"}, func(token string) bool { return len(token) > 3 }, true},
		{[]string{"bar", "foo", "bar", "foobar"}, func(token string) bool { return len(token) >= 3 }, true},
		{[]string{"bar"}, func(token string) bool { return len(token) < 1 }, false},
	}

	for _, table := range tables {
		output := All(table.arr, table.f)
		if output != table.expected {
			t.Errorf("Incorrect, got: %v, want: %v.", output, table.expected)
		}
	}
}

func TestFilter(t *testing.T) {
	tables := []struct {
		arr      []string
		f        func(string) bool
		expected []string
	}{
		{[]string{"foo"}, func(token string) bool { return len(token) > 10 }, []string{""}},
		{[]string{"foo", "bar"}, func(token string) bool { return len(token) > 3 }, []string{""}},
		{[]string{"foobar"}, func(token string) bool { return len(token) > 3 }, []string{"foobar"}},
		{[]string{"bar", "foo", "bar", "foobar"}, func(token string) bool { return len(token) >= 3 }, []string{"bar", "foo", "bar", "foobar"}},
		{[]string{"bar"}, func(token string) bool { return len(token) < 1 }, []string{""}},
	}

	for _, table := range tables {
		output := Filter(table.arr, table.f)
		if !check(table.expected, output) {
			t.Errorf("Incorrect, got: %v, want: %v.", output, table.expected)
		}
	}
}

func TestMap(t *testing.T) {
	tables := []struct {
		arr      []string
		f        func(string) string
		expected []string
	}{
		{[]string{"foo"}, func(token string) string { return strings.Replace(token, "oo", "bs", -1) }, []string{"fbs"}},
		{[]string{"foo", "bar"}, func(token string) string { return strings.Replace(token, "oo", "bs", -1) }, []string{"fbs", "bar"}},
		{[]string{"foobar"}, func(token string) string { return strings.Replace(token, "oo", "bs", -1) }, []string{"fbsbar"}},
		{[]string{"bar", "foo", "bar", "foobar"}, func(token string) string { return strings.Replace(token, "oo", "bs", -1) }, []string{"bar", "fbs", "bar", "fbsbar"}},
		{[]string{"bar"}, func(token string) string { return strings.Replace(token, "oo", "bs", -1) }, []string{"bar"}},
	}

	for _, table := range tables {
		output := Map(table.arr, table.f)
		if !check(table.expected, output) {
			t.Errorf("Incorrect, got: %v, want: %v.", output, table.expected)
		}
	}
}

func TestToStringArray(t *testing.T) {
	tables := []struct {
		arr      []interface{}
		expected []string
	}{
		{[]interface{}{"foo", "bar", "foobar", "barfoo"}, []string{"", "foo", "bar", "foobar", "barfoo"}},
	}

	for _, table := range tables {
		output := ToStringArray(table.arr)
		if !check(table.expected, output) {
			t.Errorf("Incorrect, got: %v, want: %v.", output, table.expected)
		}
	}
}

func check(expected []string, output []string) bool {
	var result = true
	for _, s := range output {
		if !Include(expected, s) {
			result = false
		}
	}
	return result
}
