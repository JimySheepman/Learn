package bellman_ford

import (
	"errors"
	"math"
)

type Graph struct {
	vertices int
	edges    map[int]map[int]int
	Directed bool
}

func New(v int) *Graph {
	return &Graph{
		vertices: v,
	}
}

func (g *Graph) AddVertex(v int) {
	if g.edges == nil {
		g.edges = make(map[int]map[int]int)
	}

	// Check if vertex is present or not
	if _, ok := g.edges[v]; !ok {
		g.edges[v] = make(map[int]int)
	}
}

func (g *Graph) AddEdge(one, two int) {
	g.AddWeightedEdge(one, two, 0)
}

func (g *Graph) AddWeightedEdge(one, two, weight int) {
	g.AddVertex(one)
	g.AddVertex(two)
	g.edges[one][two] = weight
	if !g.Directed {
		g.edges[two][one] = weight
	}

}

func (g *Graph) BellmanFord(start, end int) (isReachable bool, distance int, err error) {
	INF := math.Inf(1)
	distances := make([]float64, g.vertices)

	for i := 0; i < g.vertices; i++ {
		distances[i] = INF
	}
	distances[start] = 0

	for n := 0; n < g.vertices; n++ {

		for u, adjacents := range g.edges {
			for v, weightUV := range adjacents {

				if newDistance := distances[u] + float64(weightUV); distances[v] > newDistance {
					distances[v] = newDistance
				}
			}
		}
	}

	for u, adjacents := range g.edges {
		for v, weightUV := range adjacents {
			if newDistance := distances[u] + float64(weightUV); distances[v] > newDistance {
				return false, -1, errors.New("negative weight cycle present")
			}
		}
	}

	return distances[end] != INF, int(distances[end]), nil
}
