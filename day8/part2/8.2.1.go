package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	x int
	y int
	z int
}

// src & tgt are indices in nodes slice
type Edge struct {
	source int
	target int
	length int
}

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
		source := Node{x: x, y: y, z: z}
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

	// sort edges by length
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].length < edges[j].length
	})

	// setup initial isolated circuits
	circuits := map[int]map[int]bool{}
	for i := range nodes {
		circuits[i] = map[int]bool{i: true}
	}
	println("Initial circuits:", len(circuits))

	finalEdge := -1
	for e := 0; len(circuits) > 1; e++ {
		edge := edges[e]
		src, tgt := -1, -1
		for i, c := range circuits {
			if c[edge.source] {
				src = i
			}
			if c[edge.target] {
				tgt = i
			}
		}
		if src != tgt {
			println("Circuit", src, "- size", len(circuits[src]), "was merged into Circuit", tgt, "- size", len(circuits[tgt]))
			for x := range circuits[src] {
				circuits[tgt][x] = true
			}
			delete(circuits, src)
			finalEdge = e
		}
	}

	src := nodes[edges[finalEdge].source].x
	tgt := nodes[edges[finalEdge].target].x
	fmt.Println("Last edge:", src, "->", tgt, "after processing", finalEdge, "edges")
	fmt.Println("Product:", src*tgt)
}
