package jsongraph

// Metadata is free form information that can be associated with a graph, node, or edge.
type Metadata map[string]string

// Graph is a set of nodes and edges. Graph is a json.Marshaler and a json.Unmarshaler.
type Graph struct {
	Name      string   `json:"name,omitempty"`
	GraphType string   `json:"graph_type,omitempty"`
	Directed  bool     `json:"directed"`
	Metadata  Metadata `json:"metadata,omitempty"`
	Nodes     []Node   `json:"nodes"`
	Edges     []Edge   `json:"edges"`
}

// Node is a point on a graph. Node is a json.Marshaler and a json.Unmarshaler.
type Node struct {
	ID       string   `json:"id"`
	Label    string   `json:"label,omitempty"`
	NodeType string   `json:"node_type,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}

// Edge is a connection between two nodes. Edge is a json.Marshaler and a json.Unmarshaler.
type Edge struct {
	SourceID string   `json:"source_id"`
	TargetID string   `json:"target_id"`
	Directed bool     `json:"directed"`
	Relation string   `json:"relation,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}
