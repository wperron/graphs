package graphs

type Node struct {
	val string
}

type Relation struct {
	nodes [2]Node
}

type Graph struct {
	nodes []Node
	edges map[Node][]Node
}

func NewNode(v string) *Node {
	return &Node{v}
}

func NewRelation(n1, n2 Node) *Relation {
	return &Relation{[2]Node{n1, n2}}
}

func NewGraph(relations []Relation) *Graph {
	var g Graph
	g.edges = make(map[Node][]Node)
	for _, r := range relations {
		for _, n := range r.nodes {
			if !contains(g.nodes, n) {
				g.nodes = append(g.nodes, n)
			}
		}
		g.edges[r.nodes[0]] = append(g.edges[r.nodes[0]], r.nodes[1])
		g.edges[r.nodes[1]] = append(g.edges[r.nodes[1]], r.nodes[0])
	}
	return &g
}

func (g *Graph) Nodes() []Node {
	return g.nodes
}

func (g *Graph) Edges() map[Node][]Node {
	return g.edges
}

func (g *Graph) Distance(a, b Node) int {
	queue := []Node{a}
	visited := []Node{a}
	dists := make(map[Node]int)
	for len(queue) > 0 {
		curr := queue[0]
		adjacents := g.edges[curr]
		dist := dists[curr]
		for _, next := range adjacents {
			if contains(visited, next) {
				continue
			}
			dists[next] = dist + 1
			visited = append(visited, next)
			queue = append(queue, next)
		}
		queue = queue[1:]
	}
	return dists[b]
}

func contains(s []Node, e Node) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
