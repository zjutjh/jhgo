package utils

// Dedup removes duplicate elements from a slice while maintaining the order of the first occurrence.
// It returns a new slice containing only unique elements.
func Dedup[K comparable](s []K) []K {
	seen := make(map[K]struct{}, len(s))
	placeIdx := 0
	for idx, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			if placeIdx != idx {
				s[placeIdx] = v
			}
			placeIdx++
		}
	}
	if placeIdx < len(s) {
		s = s[:placeIdx]
	}
	return s
}
