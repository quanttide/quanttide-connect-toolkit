// Package connect 提供沟通管理领域模型。
//
// 模型以 docs/specification/content/consensus.md 的领域模型定义为准，
// JSON 标签与 API 规格字段一一对应，供各应用与工具复用。
package connect

import "encoding/json"

// Message 是消息，沟通的基本载体。
type Message struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Role      string `json:"role"`                // "user" / "agent" / "system"
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Consensus 是共识，团队成员通过讨论达成的一致决策或结论。
type Consensus struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// ConsensusRelation 是共识关系，建立共识之间的逻辑关联。
type ConsensusRelation struct {
	ID           string `json:"id"`
	From         string `json:"from"`
	To           string `json:"to"`
	RelationType string `json:"relation_type"` // "前置条件" / "支持" / "反对" / "补充"
}

// ConsensusGraph 是共识图，多个共识以 DAG 结构组织的集合。
type ConsensusGraph struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description,omitempty"`
	Nodes       []*Consensus         `json:"nodes"`
	Edges       []*ConsensusRelation `json:"edges"`
	CreatedAt   string               `json:"created_at"`
	UpdatedAt   string               `json:"updated_at,omitempty"`
}

// ParseMessages 从 JSON 数据解析消息列表。
func ParseMessages(data []byte) ([]*Message, error) {
	var out []*Message
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ParseConsensuses 从 JSON 数据解析共识列表。
func ParseConsensuses(data []byte) ([]*Consensus, error) {
	var out []*Consensus
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ParseConsensusRelations 从 JSON 数据解析共识关系列表。
func ParseConsensusRelations(data []byte) ([]*ConsensusRelation, error) {
	var out []*ConsensusRelation
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ParseConsensusGraphs 从 JSON 数据解析共识图列表。
func ParseConsensusGraphs(data []byte) ([]*ConsensusGraph, error) {
	var out []*ConsensusGraph
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}
