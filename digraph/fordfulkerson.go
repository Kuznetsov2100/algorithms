package digraph

import (
	"fmt"
	"math"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
)

type FordFulkerson struct {
	v      int
	marked []bool
	edgeTo []*FlowEdge
	value  float64
}

func NewFordFulkerson(G *FlowNetwork, s, t int) *FordFulkerson {
	ford := &FordFulkerson{v: G.V()}
	ford.validate(s)
	ford.validate(t)
	if s == t {
		panic("source equals sink")
	}
	if !ford.isFeasible(G, s, t) {
		panic("initial flow is infeasible")
	}

	ford.value = ford.excess(G, t)
	for ford.hasAugmentingPath(G, s, t) {
		bottle := math.Inf(1)
		for v := t; v != s; v = ford.edgeTo[v].Other(v) {
			bottle = math.Min(bottle, ford.edgeTo[v].ResidualCapacityTo(v))
		}

		for v := t; v != s; v = ford.edgeTo[v].Other(v) {
			ford.edgeTo[v].AddResidualFlowTo(v, bottle)
		}
		ford.value += bottle
	}

	return ford
}

func (ford *FordFulkerson) Value() float64 {
	return ford.value
}

func (ford *FordFulkerson) InCut(v int) bool {
	ford.validate(v)
	return ford.marked[v]
}

func (ford *FordFulkerson) hasAugmentingPath(G *FlowNetwork, s, t int) bool {
	ford.edgeTo = make([]*FlowEdge, G.V())
	ford.marked = make([]bool, G.V())

	queue := arrayqueue.New()
	queue.Enqueue(s)
	ford.marked[s] = true
	for !queue.IsEmpty() && !ford.marked[t] {
		val, _ := queue.Dequeue()
		v := val.(int)
		for _, e := range G.Adj(v) {
			w := e.Other(v)
			if e.ResidualCapacityTo(w) > 0 {
				if !ford.marked[w] {
					ford.edgeTo[w] = e
					ford.marked[w] = true
					queue.Enqueue(w)
				}
			}
		}
	}
	return ford.marked[t]
}

func (ford *FordFulkerson) excess(G *FlowNetwork, v int) float64 {
	excess := 0.0
	for _, e := range G.Adj(v) {
		if v == e.From() {
			excess -= e.Flow()
		} else {
			excess += e.Flow()
		}
	}
	return excess
}

func (ford *FordFulkerson) isFeasible(G *FlowNetwork, s, t int) bool {
	EPSILON := 1e-11
	for v := 0; v < G.V(); v++ {
		for _, e := range G.Adj(v) {
			if e.Flow() < -EPSILON || e.Flow() > e.Capacity()+EPSILON {
				fmt.Println("Edge does not satisfy capacity constraint: ", e)
				return false
			}
		}
	}

	if math.Abs(ford.value+ford.excess(G, s)) > EPSILON {
		fmt.Println("excess at source = ", ford.excess(G, s))
		fmt.Println("Max flow     = ", ford.value)
		return false
	}

	if math.Abs(ford.value-ford.excess(G, t)) > EPSILON {
		fmt.Println("excess at sink = ", ford.excess(G, t))
		fmt.Println("Max flow     = ", ford.value)
		return false
	}

	for v := 0; v < G.V(); v++ {
		if v == s || v == t {
			continue
		}
		if math.Abs(ford.excess(G, v)) > EPSILON {
			fmt.Println("Net flow out of ", v, " doesn't equal zero")
			return false
		}
	}
	return true
}

func (ford *FordFulkerson) validate(v int) {
	if v < 0 || v >= ford.v {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", ford.v-1))
	}
}
