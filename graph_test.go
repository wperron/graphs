package graphs

import (
	"testing"
)

/*
COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN
*/
func TestDistance(t *testing.T) {
	com := Node{"COM"}
	b := Node{"B"}
	c := Node{"C"}
	d := Node{"D"}
	e := Node{"E"}
	f := Node{"F"}
	g := Node{"G"}
	h := Node{"H"}
	i := Node{"I"}
	j := Node{"J"}
	k := Node{"K"}
	l := Node{"L"}
	san := Node{"SAN"}
	you := Node{"YOU"}

	comb := Relation{[2]Node{com, b}}
	bc := Relation{[2]Node{b, c}}
	cd := Relation{[2]Node{c, d}}
	de := Relation{[2]Node{d, e}}
	ef := Relation{[2]Node{e, f}}
	bg := Relation{[2]Node{b, g}}
	gh := Relation{[2]Node{g, h}}
	di := Relation{[2]Node{d, i}}
	ej := Relation{[2]Node{e, j}}
	jk := Relation{[2]Node{j, k}}
	kl := Relation{[2]Node{k, l}}
	kyou := Relation{[2]Node{k, you}}
	isan := Relation{[2]Node{i, san}}

	graph := NewGraph([]Relation{
		comb,
		bc,
		cd,
		de,
		ef,
		bg,
		gh,
		di,
		ej,
		jk,
		kl,
		kyou,
		isan,
	})

	dist := graph.Distance(k, i)

	if dist != 4 {
		t.Errorf("Wrong distance, expected %d, got %d", 4, dist)
	}
}
