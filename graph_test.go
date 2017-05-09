package jsongraph

import (
	"strings"
	"testing"
)

func TestMarshalling(t *testing.T) {
	graph := Graph{
		Label: "Test Graph",
	}

	node1 := Node{ID: "Node #1"}
	node2 := Node{ID: "Node #2"}
	node3 := Node{ID: "Node #3"}
	node4 := Node{ID: "Node #4"}

	graph.Nodes = append(graph.Nodes, node1)
	graph.Nodes = append(graph.Nodes, node2)
	graph.Nodes = append(graph.Nodes, node3)
	graph.Nodes = append(graph.Nodes, node4)

	edge12 := Edge{SourceID: node1.ID, TargetID: node2.ID}
	edge23 := Edge{SourceID: node2.ID, TargetID: node3.ID}
	edge24 := Edge{SourceID: node2.ID, TargetID: node4.ID}

	graph.Edges = append(graph.Edges, edge12)
	graph.Edges = append(graph.Edges, edge23)
	graph.Edges = append(graph.Edges, edge24)

	b, err := graph.ToJSON()
	if err != nil {
		t.Log("Marshalling error")
		t.Fail()
	}

	expected := `
{
	"graph": {
		"label": "Test Graph",
		"directed": false,
		"nodes": [
			{
				"id": "Node #1"
			},
			{
				"id": "Node #2"
			},
			{
				"id": "Node #3"
			},
			{
				"id": "Node #4"
			}
		],
		"edges": [
			{
				"source": "Node #1",
				"target": "Node #2",
				"directed": false
			},
			{
				"source": "Node #2",
				"target": "Node #3",
				"directed": false
			},
			{
				"source": "Node #2",
				"target": "Node #4",
				"directed": false
			}
		]
	}
}`

	if string(b) != strings.TrimSpace(expected) {
		t.Logf("JSON was not what was expected.\n\nExpected: %s\nActual: %s", expected, string(b))
		t.Fail()
	}
}
