//! 共识领域事件定义。
//!
//! 事件类型以 docs/specification/content/consensus.md 的领域事件为准。

use serde::{Deserialize, Serialize};

/// 领域事件类型常量
pub const EVENT_CONSENSUS_CREATED: &str = "ConsensusCreated";
pub const EVENT_CONSENSUS_UPDATED: &str = "ConsensusUpdated";
pub const EVENT_CONSENSUS_DELETED: &str = "ConsensusDeleted";
pub const EVENT_CONSENSUS_RELATION_CREATED: &str = "ConsensusRelationCreated";
pub const EVENT_CONSENSUS_RELATION_DELETED: &str = "ConsensusRelationDeleted";
pub const EVENT_CONSENSUS_GRAPH_NODE_ADDED: &str = "ConsensusGraphNodeAdded";
pub const EVENT_CONSENSUS_GRAPH_NODE_REMOVED: &str = "ConsensusGraphNodeRemoved";
pub const EVENT_CONSENSUS_GRAPH_EDGE_ADDED: &str = "ConsensusGraphEdgeAdded";
pub const EVENT_CONSENSUS_GRAPH_EDGE_REMOVED: &str = "ConsensusGraphEdgeRemoved";

/// 领域事件的基础结构
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Event {
    /// 事件ID
    pub event_id: String,
    /// 事件类型
    pub event_type: String,
    /// 事件时间戳
    pub timestamp: String,
}

/// ConsensusCreated 事件的数据
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusCreatedData {
    /// 共识ID
    pub consensus_id: String,
    /// 共识标题
    pub title: String,
    /// 共识描述
    pub description: String,
    /// 共识创建时间
    pub created_at: String,
}

/// ConsensusCreated 事件
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusCreatedEvent {
    #[serde(flatten)]
    pub event: Event,
    pub data: ConsensusCreatedData,
}

/// ConsensusUpdated 事件的数据
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusUpdatedData {
    /// 共识ID
    pub consensus_id: String,
    /// 共识标题
    pub title: String,
    /// 共识描述
    pub description: String,
    /// 共识更新时间
    pub updated_at: String,
}

/// ConsensusUpdated 事件
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusUpdatedEvent {
    #[serde(flatten)]
    pub event: Event,
    pub data: ConsensusUpdatedData,
}

/// ConsensusDeleted 事件的数据
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusDeletedData {
    /// 共识ID
    pub consensus_id: String,
}

/// ConsensusDeleted 事件
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusDeletedEvent {
    #[serde(flatten)]
    pub event: Event,
    pub data: ConsensusDeletedData,
}

/// ConsensusRelationCreated 事件的数据
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusRelationCreatedData {
    /// 关系ID
    pub relation_id: String,
    /// 关系起点的共识 ID
    pub from: String,
    /// 关系终点的共识 ID
    pub to: String,
    /// 关系类型
    pub relation_type: String,
}

/// ConsensusRelationCreated 事件
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusRelationCreatedEvent {
    #[serde(flatten)]
    pub event: Event,
    pub data: ConsensusRelationCreatedData,
}

/// ConsensusRelationDeleted 事件的数据
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusRelationDeletedData {
    /// 关系ID
    pub relation_id: String,
}

/// ConsensusRelationDeleted 事件
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusRelationDeletedEvent {
    #[serde(flatten)]
    pub event: Event,
    pub data: ConsensusRelationDeletedData,
}

/// ConsensusGraphNodeAdded 事件的数据
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusGraphNodeAddedData {
    /// 共识图ID
    pub graph_id: String,
    /// 共识ID
    pub consensus_id: String,
}

/// ConsensusGraphNodeAdded 事件
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusGraphNodeAddedEvent {
    #[serde(flatten)]
    pub event: Event,
    pub data: ConsensusGraphNodeAddedData,
}

/// ConsensusGraphNodeRemoved 事件的数据
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusGraphNodeRemovedData {
    /// 共识图ID
    pub graph_id: String,
    /// 共识ID
    pub consensus_id: String,
}

/// ConsensusGraphNodeRemoved 事件
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusGraphNodeRemovedEvent {
    #[serde(flatten)]
    pub event: Event,
    pub data: ConsensusGraphNodeRemovedData,
}

/// ConsensusGraphEdgeAdded 事件的数据
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusGraphEdgeAddedData {
    /// 共识图ID
    pub graph_id: String,
    /// 关系ID
    pub relation_id: String,
}

/// ConsensusGraphEdgeAdded 事件
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusGraphEdgeAddedEvent {
    #[serde(flatten)]
    pub event: Event,
    pub data: ConsensusGraphEdgeAddedData,
}

/// ConsensusGraphEdgeRemoved 事件的数据
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusGraphEdgeRemovedData {
    /// 共识图ID
    pub graph_id: String,
    /// 关系ID
    pub relation_id: String,
}

/// ConsensusGraphEdgeRemoved 事件
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConsensusGraphEdgeRemovedEvent {
    #[serde(flatten)]
    pub event: Event,
    pub data: ConsensusGraphEdgeRemovedData,
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_consensus_created_event() {
        let event = ConsensusCreatedEvent {
            event: Event {
                event_id: "e0a80101-0000-0000-0000-000000000001".to_string(),
                event_type: EVENT_CONSENSUS_CREATED.to_string(),
                timestamp: "2026-08-28T14:30:00+08:00".to_string(),
            },
            data: ConsensusCreatedData {
                consensus_id: "c0a80101-0000-0000-0000-000000000001".to_string(),
                title: "共识是沟通管理领域的核心概念".to_string(),
                description: "经过团队讨论，我们一致认为共识是沟通从分歧到统一的关键产出物。"
                    .to_string(),
                created_at: "2026-08-28T14:30:00+08:00".to_string(),
            },
        };
        let json = serde_json::to_string(&event).unwrap();
        assert!(json.contains("ConsensusCreated"));
        assert!(json.contains("共识是沟通管理领域的核心概念"));
    }
}
