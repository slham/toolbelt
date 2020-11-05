package partition

import (
	"fmt"
	"github.com/meirf/gopart"
	"time"
)

type PercentPartition struct {
	Left       int
	Right      int
	Segment    int
	Partitions int
	Length     int
}

func (p *PercentPartition) Next() {
	p.Right = p.Length * p.Segment / p.Partitions
}

func (p *PercentPartition) Before(them []int)*PercentPartition {
	return p
}

func (p *PercentPartition) Process(them []int)*PercentPartition {
	fmt.Println(them)
	return p
}

func (p *PercentPartition) After(them []int)*PercentPartition {
	return p
}

func (p *PercentPartition) Do(them []int) {
	p.Next()
	for p.Right != p.Length {
		p.Before(them[p.Left:p.Right]).Process(them[p.Left:p.Right]).After(them[p.Left:p.Right])
		p.Left = p.Right
		p.Segment ++
		p.Next()
	}
	p.Before(them[p.Left:p.Right]).Process(them[p.Left:p.Right]).After(them[p.Left:p.Right])
}

func (p *PercentPartition) Time(arr []int, desc string) {
	start := time.Now()
	p.Do(arr)
	fmt.Printf("%s took %v\n", desc, time.Since(start))
}

type SimplePartition struct {
	Partitions int
}

func (p *SimplePartition) Next() {}

func (p *SimplePartition) Before(them []int) *SimplePartition {
	return p
}
func (p *SimplePartition) Process(them []int) *SimplePartition {
	fmt.Println(them)
	return p
}
func (p *SimplePartition) After(them []int) *SimplePartition {
	return p
}

func (p *SimplePartition) Do(them []int) {
	for idxRange := range gopart.Partition(len(them), p.Partitions) {
		segment := them[idxRange.Low:idxRange.High]
		p.Before(segment).Process(segment).After(segment)
	}
}

func (p *SimplePartition) Time(arr []int, desc string) {
	start := time.Now()
	p.Do(arr)
	fmt.Printf("%s took %v\n", desc, time.Since(start))
}
