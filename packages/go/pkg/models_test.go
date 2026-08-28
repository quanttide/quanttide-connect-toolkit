package connect_test

import (
	"encoding/json"
	"strings"
	"testing"

	connect "github.com/quanttide/quanttide-connect-toolkit/packages/go/pkg"
)

// 共识标本：与 docs/specification/content/consensus.md 示例数据同构。
const consensusTree = `{
  "id": "c0a80101-0000-0000-0000-000000000001",
  "title": "共识是沟通管理领域的核心概念",
  "description": "经过团队讨论，我们一致认为共识是沟通从分歧到统一的关键产出物，是沟通管理领域的核心概念。共识具有决策性、可追溯性、不可篡改性和血缘关联等特征，是沟通云系统的核心节点。",
  "created_at": "2026-08-28T14:30:00+08:00",
  "updated_at": "2026-08-28T14:35:00+08:00"
}`

const consensusRelationTree = `{
  "id": "r0a80101-0000-0000-0000-000000000001",
  "from": "c0a80101-0000-0000-0000-000000000001",
  "to": "c0a80101-0000-0000-0000-000000000002",
  "relation_type": "前置条件"
}`

const consensusGraphTree = `{
  "id": "g0a80101-0000-0000-0000-000000000001",
  "name": "沟通管理标准建模",
  "description": "团队讨论如何建模沟通管理标准的决策网络，包含共识、共识关系、共识图等概念的定义过程。",
  "nodes": [
    {
      "id": "c0a80101-0000-0000-0000-000000000001",
      "title": "共识是沟通管理领域的核心概念",
      "description": "经过团队讨论，我们一致认为共识是沟通从分歧到统一的关键产出物，是沟通管理领域的核心概念。共识具有决策性、可追溯性、不可篡改性和血缘关联等特征，是沟通云系统的核心节点。",
      "created_at": "2026-08-28T14:30:00+08:00",
      "updated_at": "2026-08-28T14:35:00+08:00"
    },
    {
      "id": "c0a80101-0000-0000-0000-000000000002",
      "title": "共识需要支持不可篡改性",
      "description": "共识一旦标记，内容和元数据（时间戳、参与人）不可修改，确保决策的严肃性。",
      "created_at": "2026-08-28T14:45:00+08:00"
    }
  ],
  "edges": [
    {
      "id": "r0a80101-0000-0000-0000-000000000001",
      "from": "c0a80101-0000-0000-0000-000000000001",
      "to": "c0a80101-0000-0000-0000-000000000002",
      "relation_type": "前置条件"
    }
  ],
  "created_at": "2026-08-28T14:00:00+08:00",
  "updated_at": "2026-08-28T14:00:00+08:00"
}`

// TestConsensus_UnmarshalJSON 字段与 API 规格对齐（snake_case、omitempty 不回填零值）。
func TestConsensus_UnmarshalJSON(t *testing.T) {
	var c connect.Consensus
	if err := json.Unmarshal([]byte(consensusTree), &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if c.ID != "c0a80101-0000-0000-0000-000000000001" || c.Title != "共识是沟通管理领域的核心概念" {
		t.Fatalf("unexpected consensus: %+v", c)
	}
	if c.UpdatedAt == "" {
		t.Fatalf("updated_at should not be empty")
	}
}

// TestConsensusRelation_UnmarshalJSON 字段与 API 规格对齐。
func TestConsensusRelation_UnmarshalJSON(t *testing.T) {
	var r connect.ConsensusRelation
	if err := json.Unmarshal([]byte(consensusRelationTree), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if r.ID != "r0a80101-0000-0000-0000-000000000001" || r.RelationType != "前置条件" {
		t.Fatalf("unexpected relation: %+v", r)
	}
}

// TestConsensusGraph_UnmarshalJSON 字段与 API 规格对齐。
func TestConsensusGraph_UnmarshalJSON(t *testing.T) {
	var g connect.ConsensusGraph
	if err := json.Unmarshal([]byte(consensusGraphTree), &g); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if g.ID != "g0a80101-0000-0000-0000-000000000001" || g.Name != "沟通管理标准建模" {
		t.Fatalf("unexpected graph: %+v", g)
	}
	if len(g.Nodes) != 2 || len(g.Edges) != 1 {
		t.Fatalf("unexpected nodes/edges count: nodes=%d, edges=%d", len(g.Nodes), len(g.Edges))
	}
	if g.Nodes[0].ID != "c0a80101-0000-0000-0000-000000000001" {
		t.Fatalf("unexpected first node: %+v", g.Nodes[0])
	}
	if g.Edges[0].RelationType != "前置条件" {
		t.Fatalf("unexpected edge relation type: %s", g.Edges[0].RelationType)
	}
}

// TestConsensus_OmitZeroFields 零值字段不输出（omitempty 与 API 规格一致）。
func TestConsensus_OmitZeroFields(t *testing.T) {
	c := connect.Consensus{
		ID:    "test-id",
		Title: "测试共识",
	}
	raw, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	for _, banned := range []string{"updated_at", "description"} {
		if strings.Contains(string(raw), banned) {
			t.Fatalf("字段 %s 不应出现在零值序列化结果：%s", banned, raw)
		}
	}
}

// TestParseConsensuses 解析共识列表。
func TestParseConsensuses(t *testing.T) {
	data := `[` + consensusTree + `]`
	consensuses, err := connect.ParseConsensuses([]byte(data))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(consensuses) != 1 || consensuses[0].ID != "c0a80101-0000-0000-0000-000000000001" {
		t.Fatalf("unexpected consensuses: %+v", consensuses)
	}
}
