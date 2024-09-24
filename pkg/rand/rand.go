package rand

import "math/rand/v2"

// nolint:revive
// nolint:gosec
func RandIntForInterval(min, max int) int {
	// nolint:gosec
	return rand.IntN(max+1-min) + min
}
