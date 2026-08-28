# quanttide-connect Rust 库

沟通管理领域 Rust 语言工具包，提供数据模型、API 路由和领域事件定义。

## 安装

在 `Cargo.toml` 中添加：

```toml
[dependencies]
quanttide-connect = "0.1.0"
```

## 使用

### 数据模型

```rust
use quanttide_connect::consensus::Consensus;

fn main() {
    let consensus_json = r#"{
        "id": "c0a80101-0000-0000-0000-000000000001",
        "title": "共识是沟通管理领域的核心概念",
        "description": "经过团队讨论，我们一致认为共识是沟通从分歧到统一的关键产出物。",
        "created_at": "2026-08-28T14:30:00+08:00"
    }"#;

    let consensus = quanttide_connect::consensus::parse_consensus(consensus_json.as_bytes()).unwrap();
    println!("共识ID: {}", consensus.id);
    println!("共识标题: {}", consensus.title);
}
```

### API 路由

```rust
use quanttide_connect::consensus::api;

fn main() {
    let consensus_id = "c0a80101-0000-0000-0000-000000000001";
    let path = api::consensus_path(consensus_id);
    println!("共识详情路径: {}", path);

    let graph_id = "g0a80101-0000-0000-0000-000000000001";
    let node_path = api::consensus_graph_node_path(graph_id, consensus_id);
    println!("共识图节点路径: {}", node_path);
}
```

### 领域事件

```rust
use quanttide_connect::consensus::events::{Event, ConsensusCreatedEvent, ConsensusCreatedData, EVENT_CONSENSUS_CREATED};

fn main() {
    let event = ConsensusCreatedEvent {
        event: Event {
            event_id: "e0a80101-0000-0000-0000-000000000001".to_string(),
            event_type: EVENT_CONSENSUS_CREATED.to_string(),
            timestamp: "2026-08-28T14:30:00+08:00".to_string(),
        },
        data: ConsensusCreatedData {
            consensus_id: "c0a80101-0000-0000-0000-000000000001".to_string(),
            title: "共识是沟通管理领域的核心概念".to_string(),
            description: "经过团队讨论，我们一致认为共识是沟通从分歧到统一的关键产出物。".to_string(),
            created_at: "2026-08-28T14:30:00+08:00".to_string(),
        },
    };

    let json = serde_json::to_string(&event).unwrap();
    println!("事件JSON: {}", json);
}
```

## 数据模型

本包提供以下数据模型：

- **Consensus**：共识，团队成员通过讨论达成的一致决策或结论
- **ConsensusRelation**：共识关系，建立共识之间的逻辑关联
- **ConsensusGraph**：共识图，多个共识以 DAG 结构组织的集合

所有模型定义以 `docs/specification/content/consensus.md` 的领域模型定义为准。

## API 路由

本包提供以下 API 路由常量和构造函数：

- **静态资源集合路径**：`/consensuses`、`/consensus-relations`、`/consensus-graphs`
- **单资源路径**：`/consensuses/{id}`、`/consensus-relations/{id}`、`/consensus-graphs/{id}`
- **子资源路径**：`/consensuses/{id}/relations`、`/consensus-graphs/{id}/nodes` 等

所有路由定义以 `docs/specification/content/consensus.md` 的 API 规格为准。

## 领域事件

本包提供以下领域事件类型：

- **共识相关事件**：`ConsensusCreated`、`ConsensusUpdated`、`ConsensusDeleted`
- **共识关系相关事件**：`ConsensusRelationCreated`、`ConsensusRelationDeleted`
- **共识图节点相关事件**：`ConsensusGraphNodeAdded`、`ConsensusGraphNodeRemoved`
- **共识图边相关事件**：`ConsensusGraphEdgeAdded`、`ConsensusGraphEdgeRemoved`

所有事件定义以 `docs/specification/content/consensus.md` 的领域事件为准。

## 版本号

本包遵循语义化版本规范，版本号以契约为准：

- **alpha**：预览版本，契约可能变化
- **beta**：候选版本，契约基本稳定
- **正式版**：稳定版本，契约不变

## 贡献

请参考 [CONTRIBUTING.md](CONTRIBUTING.md) 了解贡献指南。
