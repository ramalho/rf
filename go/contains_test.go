package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		list []string
		item string
		want bool
		name string
	}{
		{[]string{}, "B", false, "empty list"},
		{[]string{"A"}, "B", false, "not found"},
		{[]string{"B"}, "B", true, "found"},
		{[]string{"A", "B"}, "B", true, "found last"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := contains(tc.list, tc.item)
			if got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	testCases := []struct {
		list  []string
		items []string
		want  bool
		name  string
	}{
		{[]string{}, []string{"B"}, false, "1 item, empty list"},
		{[]string{"A"}, []string{"B"}, false, "1 item, not found"},
		{[]string{"B"}, []string{"B"}, true, "1 item, found"},
		{[]string{"A", "B"}, []string{"B"}, true, "1 item, found last"},
		{[]string{}, []string{"B", "A"}, false, "2 items, empty list"},
		{[]string{"A"}, []string{"B", "A"}, false, "2 items, B not found"},
		{[]string{"A"}, []string{"B", "A"}, false, "2 items, A not found"},
		{[]string{"A", "B"}, []string{"B", "A"}, true, "2 items, found"},
		{[]string{"A", "B", "C"}, []string{"B", "A"}, true, "2 items, found with extra"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := containsAll(tc.list, tc.items)
			if got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
