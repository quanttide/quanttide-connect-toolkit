package consensus

// 领域事件类型常量（唯一事实源：docs/specification/content/consensus.md 领域事件）。
const (
	// 共识相关事件。
	EventConsensusCreated = "ConsensusCreated"
	EventConsensusUpdated = "ConsensusUpdated"
	EventConsensusDeleted = "ConsensusDeleted"

	// 共识关系相关事件。
	EventConsensusRelationCreated = "ConsensusRelationCreated"
	EventConsensusRelationDeleted = "ConsensusRelationDeleted"

	// 共识图节点相关事件。
	EventConsensusGraphNodeAdded   = "ConsensusGraphNodeAdded"
	EventConsensusGraphNodeRemoved = "ConsensusGraphNodeRemoved"

	// 共识图边相关事件。
	EventConsensusGraphEdgeAdded   = "ConsensusGraphEdgeAdded"
	EventConsensusGraphEdgeRemoved = "ConsensusGraphEdgeRemoved"
)

// Event 是领域事件的基础结构。
type Event struct {
	EventID   string `json:"event_id"`
	EventType string `json:"event_type"`
	Timestamp string `json:"timestamp"`
}

// ConsensusCreatedData 是 ConsensusCreated 事件的数据。
type ConsensusCreatedData struct {
	ConsensusID string `json:"consensus_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// ConsensusCreatedEvent 是 ConsensusCreated 事件。
type ConsensusCreatedEvent struct {
	Event
	Data ConsensusCreatedData `json:"data"`
}

// ConsensusUpdatedData 是 ConsensusUpdated 事件的数据。
type ConsensusUpdatedData struct {
	ConsensusID string `json:"consensus_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
}

// ConsensusUpdatedEvent 是 ConsensusUpdated 事件。
type ConsensusUpdatedEvent struct {
	Event
	Data ConsensusUpdatedData `json:"data"`
}

// ConsensusDeletedData 是 ConsensusDeleted 事件的数据。
type ConsensusDeletedData struct {
	ConsensusID string `json:"consensus_id"`
}

// ConsensusDeletedEvent 是 ConsensusDeleted 事件。
type ConsensusDeletedEvent struct {
	Event
	Data ConsensusDeletedData `json:"data"`
}

// ConsensusRelationCreatedData 是 ConsensusRelationCreated 事件的数据。
type ConsensusRelationCreatedData struct {
	RelationID   string `json:"relation_id"`
	From         string `json:"from"`
	To           string `json:"to"`
	RelationType string `json:"relation_type"`
}

// ConsensusRelationCreatedEvent 是 ConsensusRelationCreated 事件。
type ConsensusRelationCreatedEvent struct {
	Event
	Data ConsensusRelationCreatedData `json:"data"`
}

// ConsensusRelationDeletedData 是 ConsensusRelationDeleted 事件的数据。
type ConsensusRelationDeletedData struct {
	RelationID string `json:"relation_id"`
}

// ConsensusRelationDeletedEvent 是 ConsensusRelationDeleted 事件。
type ConsensusRelationDeletedEvent struct {
	Event
	Data ConsensusRelationDeletedData `json:"data"`
}

// ConsensusGraphNodeAddedData 是 ConsensusGraphNodeAdded 事件的数据。
type ConsensusGraphNodeAddedData struct {
	GraphID     string `json:"graph_id"`
	ConsensusID string `json:"consensus_id"`
}

// ConsensusGraphNodeAddedEvent 是 ConsensusGraphNodeAdded 事件。
type ConsensusGraphNodeAddedEvent struct {
	Event
	Data ConsensusGraphNodeAddedData `json:"data"`
}

// ConsensusGraphNodeRemovedData 是 ConsensusGraphNodeRemoved 事件的数据。
type ConsensusGraphNodeRemovedData struct {
	GraphID     string `json:"graph_id"`
	ConsensusID string `json:"consensus_id"`
}

// ConsensusGraphNodeRemovedEvent 是 ConsensusGraphNodeRemoved 事件。
type ConsensusGraphNodeRemovedEvent struct {
	Event
	Data ConsensusGraphNodeRemovedData `json:"data"`
}

// ConsensusGraphEdgeAddedData 是 ConsensusGraphEdgeAdded 事件的数据。
type ConsensusGraphEdgeAddedData struct {
	GraphID    string `json:"graph_id"`
	RelationID string `json:"relation_id"`
}

// ConsensusGraphEdgeAddedEvent 是 ConsensusGraphEdgeAdded 事件。
type ConsensusGraphEdgeAddedEvent struct {
	Event
	Data ConsensusGraphEdgeAddedData `json:"data"`
}

// ConsensusGraphEdgeRemovedData 是 ConsensusGraphEdgeRemoved 事件的数据。
type ConsensusGraphEdgeRemovedData struct {
	GraphID    string `json:"graph_id"`
	RelationID string `json:"relation_id"`
}

// ConsensusGraphEdgeRemovedEvent 是 ConsensusGraphEdgeRemoved 事件。
type ConsensusGraphEdgeRemovedEvent struct {
	Event
	Data ConsensusGraphEdgeRemovedData `json:"data"`
}
