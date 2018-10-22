package main

// Contains reports whether list contains item.
func Contains(list []string, item string) bool {
	for _, s := range list {
		if s == item {
			return true
		}
	}
	return false
}

// ContainsAll reports whether list contains all items.
func ContainsAll(list []string, items ...string) bool {
	for _, s := range items {
		if !Contains(list, s) {
			return false
		}
	}
	return true
}
