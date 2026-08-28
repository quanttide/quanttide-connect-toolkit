//! 共识领域模型定义。
//!
//! 模型以 docs/specification/content/consensus.md 的领域模型定义为准，
//! JSON 标签与 API 规格字段一一对应，供各应用与工具复用。

use serde::{Deserialize, Serialize};

/// 共识，团队成员通过讨论达成的一致决策或结论。
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Consensus {
    /// 共识的唯一标识符
    pub id: String,
    /// 共识的标题，简要描述决策内容
    pub title: String,
    /// 共识的详细描述
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    /// 共识创建时间
    pub created_at: String,
    /// 共识最后更新时间
    #[serde(skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<String>,
}

/// 共识关系，建立共识之间的逻辑关联。
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusRelation {
    /// 关系的唯一标识符
    pub id: String,
    /// 关系起点的共识 ID
    pub from: String,
    /// 关系终点的共识 ID
    pub to: String,
    /// 关系类型（如：支持、反对、补充、前置条件等）
    pub relation_type: String,
}

/// 共识图，多个共识以 DAG 结构组织的集合。
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusGraph {
    /// 共识图的唯一标识符
    pub id: String,
    /// 共识图的名称
    pub name: String,
    /// 共识图的描述
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    /// 图中的共识节点列表
    pub nodes: Vec<Consensus>,
    /// 图中的关系边列表
    pub edges: Vec<ConsensusRelation>,
    /// 共识图创建时间
    pub created_at: String,
    /// 共识图最后更新时间
    #[serde(skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<String>,
}

/// 从 JSON 数据解析共识。
pub fn parse_consensus(data: &[u8]) -> Result<Consensus, serde_json::Error> {
    serde_json::from_slice(data)
}

/// 从 JSON 数据解析共识列表。
pub fn parse_consensuses(data: &[u8]) -> Result<Vec<Consensus>, serde_json::Error> {
    serde_json::from_slice(data)
}

/// 从 JSON 数据解析共识关系。
pub fn parse_consensus_relation(data: &[u8]) -> Result<ConsensusRelation, serde_json::Error> {
    serde_json::from_slice(data)
}

/// 从 JSON 数据解析共识关系列表。
pub fn parse_consensus_relations(data: &[u8]) -> Result<Vec<ConsensusRelation>, serde_json::Error> {
    serde_json::from_slice(data)
}

/// 从 JSON 数据解析共识图。
pub fn parse_consensus_graph(data: &[u8]) -> Result<ConsensusGraph, serde_json::Error> {
    serde_json::from_slice(data)
}

/// 从 JSON 数据解析共识图列表。
pub fn parse_consensus_graphs(data: &[u8]) -> Result<Vec<ConsensusGraph>, serde_json::Error> {
    serde_json::from_slice(data)
}

#[cfg(test)]
mod tests {
    use super::*;

    const CONSENSUS_JSON: &str = r#"{
        "id": "c0a80101-0000-0000-0000-000000000001",
        "title": "共识是沟通管理领域的核心概念",
        "description": "经过团队讨论，我们一致认为共识是沟通从分歧到统一的关键产出物。",
        "created_at": "2026-08-28T14:30:00+08:00",
        "updated_at": "2026-08-28T14:35:00+08:00"
    }"#;

    const CONSENSUS_RELATION_JSON: &str = r#"{
        "id": "r0a80101-0000-0000-0000-000000000001",
        "from": "c0a80101-0000-0000-0000-000000000001",
        "to": "c0a80101-0000-0000-0000-000000000002",
        "relation_type": "前置条件"
    }"#;

    #[test]
    fn test_parse_consensus() {
        let consensus = parse_consensus(CONSENSUS_JSON.as_bytes()).unwrap();
        assert_eq!(consensus.id, "c0a80101-0000-0000-0000-000000000001");
        assert_eq!(consensus.title, "共识是沟通管理领域的核心概念");
        assert!(consensus.description.is_some());
        assert!(consensus.updated_at.is_some());
    }

    #[test]
    fn test_parse_consensus_relation() {
        let relation = parse_consensus_relation(CONSENSUS_RELATION_JSON.as_bytes()).unwrap();
        assert_eq!(relation.id, "r0a80101-0000-0000-0000-000000000001");
        assert_eq!(relation.relation_type, "前置条件");
    }

    #[test]
    fn test_consensus_omit_none_fields() {
        let consensus = Consensus {
            id: "test-id".to_string(),
            title: "测试共识".to_string(),
            description: None,
            created_at: "2026-08-28T14:30:00+08:00".to_string(),
            updated_at: None,
        };
        let json = serde_json::to_string(&consensus).unwrap();
        assert!(!json.contains("description"));
        assert!(!json.contains("updated_at"));
    }
}
