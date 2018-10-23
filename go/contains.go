package main

// contains reports whether list contains item.
func contains(list []string, item string) bool {
	for _, s := range list {
		if s == item {
			return true
		}
	}
	return false
}

// containsAll reports whether list contains all items.
func containsAll(list []string, items []string) bool {
	for _, s := range items {
		if !contains(list, s) {
			return false
		}
	}
	return true
}
