package main

// contains reports whether list contains item.
func contains[K comparable](list []K, item K) bool {
	for _, s := range list {
		if s == item {
			return true
		}
	}
	return false
}

// containsAll reports whether list contains all items.
func containsAll[K comparable](list []K, items []K) bool {
	for _, s := range items {
		if !contains(list, s) {
			return false
		}
	}
	return true
}
