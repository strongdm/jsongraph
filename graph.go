package jsongraph

import "encoding/json"

// Metadata is free form information that can be associated with a graph, node, or edge.
type Metadata map[string]string

// Graph is a set of nodes and edges. Graph is a json.Marshaler and a json.Unmarshaler.
type Graph struct {
	*base
	Label     string   `json:"label,omitempty"`
	GraphType string   `json:"graph_type,omitempty"`
	Directed  bool     `json:"directed,omitempty"`
	Metadata  Metadata `json:"metadata,omitempty"`
	Nodes     []Node   `json:"nodes"`
	Edges     []Edge   `json:"edges"`
}

// Node is a point on a graph. Node is a json.Marshaler and a json.Unmarshaler.
type Node struct {
	*base
	ID       string   `json:"id"`
	Label    string   `json:"label,omitempty"`
	NodeType string   `json:"node_type,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}

// Edge is a connection between two nodes. Edge is a json.Marshaler and a json.Unmarshaler.
type Edge struct {
	*base
	SourceID string   `json:"source_id"`
	TargetID string   `json:"target_id"`
	Relation string   `json:"relation,omitempty"`
	Directed bool     `json:"directed,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}

type base struct{}

// MarshalJSON builds a JSON representation of the struct.
func (b *base) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

// Unmarshal builds a JSON representation of the struct.
func (b *base) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, b)
}
