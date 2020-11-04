package partition

import (
	"fmt"
	"github.com/meirf/gopart"
)

type IntPartition interface {
	Next()
	Before([]int)*IntPartition
	Process([]int)*IntPartition
	After([]int)*IntPartition
	Do([]int)
}

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