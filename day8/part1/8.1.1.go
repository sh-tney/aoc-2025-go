package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// adj is list of indices in nodes slice
type Node struct {
	x   int
	y   int
	z   int
	adj []int
}

// src & tgt are indices in nodes slice
type Edge struct {
	source int
	target int
	length int
}

const MAX_EDGES = 1000

func v1() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic("file error")
	}
	defer file.Close()

	nodes := []Node{}
	edges := []Edge{}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		y, err := strconv.Atoi(coords[1])
		z, err := strconv.Atoi(coords[2])
		if err != nil {
			panic("parsing error")
		}
		source := Node{x: x, y: y, z: z, adj: []int{}}
		//fmt.Println("Node:", source)

		s := len(nodes)
		for t, target := range nodes {
			dx := (source.x - target.x) * (source.x - target.x)
			dy := (source.y - target.y) * (source.y - target.y)
			dz := (source.z - target.z) * (source.z - target.z)
			d := dx + dy + dz
			edge := Edge{source: s, target: t, length: d}
			edges = append(edges, edge)
			//fmt.Println("Edge:", edge)
		}

		nodes = append(nodes, source)
	}
	fmt.Println("Nodes:", len(nodes))
	fmt.Println("Edges:", len(edges))

	// sort edges by length, get shortest MAX_EDGES
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].length < edges[j].length
	})
	edges = edges[:MAX_EDGES]
	fmt.Println("Shortest", MAX_EDGES, "edges")

	// build adjacency list
	for _, edge := range edges {
		nodes[edge.source].adj = append(nodes[edge.source].adj, edge.target)
		nodes[edge.target].adj = append(nodes[edge.target].adj, edge.source)
	}

	// build circuits
	circuits := [][]int{}
	visited := map[int]bool{}
	for len(visited) != len(nodes) {
		// get unvisited node
		i := 0
		for visited[i] {
			i++
		}

		circuit := []int{}
		toVisit := []int{i}
		for len(toVisit) > 0 {
			current := toVisit[0]
			toVisit = toVisit[1:]
			if visited[current] {
				continue
			}

			visited[current] = true
			circuit = append(circuit, current)
			for _, adj := range nodes[current].adj {
				if !visited[adj] {
					toVisit = append(toVisit, adj)
				}
			}
		}

		circuits = append(circuits, circuit)
		fmt.Println("Circuit:", circuit)
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})
	fmt.Println("Top 3 Lengths:", len(circuits[0]), len(circuits[1]), len(circuits[2]))
	fmt.Println("Product:", len(circuits[0])*len(circuits[1])*len(circuits[2]))
}
