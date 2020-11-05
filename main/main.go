package main

import (
	"github.com/slham/toolbelt/partition"
)

const smallLength = 10
const mediumLength = 100
const largeLength = 1000

func intArrOfN(num int) []int {
	out := make([]int, num)
	for i := 0; i < num; i++ {
		out[i] = i
	}
	return out
}

func main() {
	small := intArrOfN(smallLength)
	medium := intArrOfN(mediumLength)
	large := intArrOfN(largeLength)

	p := partition.PercentPartition{
		Left:       0,
		Right:      0,
		Segment:    1,
	}
	s := partition.SimplePartition{Partitions: 4}

	p.Partitions = 25
	p.Length = smallLength
	p.Time(small, "percent small")
	s.Time(small, "simple small")

	p.Partitions = 250
	p.Length = mediumLength
	p.Time(medium, "percent medium")
	s.Time(medium, "simple medium")

	p.Partitions = 2500
	p.Length = largeLength
	p.Time(large, "percent large")
	s.Time(large, "simple large")
}
