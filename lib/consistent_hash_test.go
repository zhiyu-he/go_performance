package lib

import (
	"github.com/serialx/hashring"
	"testing"
)

var (
	ip_list = []string{
		"10.8.69.212",
		"10.8.70.51",
		"10.8.72.89",
		"10.8.72.98",
	}
	ring = hashring.New(ip_list)
	ring2 = hashring.NewWithVNodeSize(ip_list, 400)
	ring3 = hashring.NewWithVNodeSize(ip_list, 800)
)


func Check(r *hashring.HashRing, t *testing.T) {
	sortedKeys := r.GetSortedKeys()
	min := uint32(0)
	max := ^uint32(0)
	t.Logf("%v [%v %v] %v %v\n", len(sortedKeys), sortedKeys[0], sortedKeys[len(sortedKeys)-1], min, max)

	expectedBoundary := max / uint32(len(sortedKeys))

	var maxDistance, totalDistance uint32
	for i, rightBoundary := range sortedKeys {
		var diff uint32
		if i == 0 {
			diff = uint32(rightBoundary) - min
			goto end
		}
		if i == len(sortedKeys) - 1 {
			diff = max - uint32(sortedKeys[len(sortedKeys) - 1])
			goto end
		}
		diff = uint32(sortedKeys[i] - sortedKeys[i - 1])
		end:
			totalDistance += diff
			if diff > maxDistance {
				maxDistance = diff
			}
	}
	actualBoundary := totalDistance/uint32(len(sortedKeys))
	t.Logf("(%d, %d) max: %d times: %d\n", expectedBoundary, actualBoundary, maxDistance, maxDistance/expectedBoundary)
}


func TestHashRingRange(t *testing.T) {
	Check(ring, t)
	Check(ring2, t)
	Check(ring3, t)
}
