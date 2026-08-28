# quanttide-connect-toolkit Go 包

沟通管理领域 Go 语言工具包，提供数据模型、API 路由和领域事件定义。

## 安装

```bash
go get github.com/quanttide/quanttide-connect-toolkit/packages/go/pkg
```

## 使用

### 数据模型

```go
package main

import (
	"encoding/json"
	"fmt"

	connect "github.com/quanttide/quanttide-connect-toolkit/packages/go/pkg"
)

func main() {
	// 解析共识
	consensusJSON := `{
		"id": "c0a80101-0000-0000-0000-000000000001",
		"title": "共识是沟通管理领域的核心概念",
		"description": "经过团队讨论，我们一致认为共识是沟通从分歧到统一的关键产出物。",
		"created_at": "2026-08-28T14:30:00+08:00"
	}`
	
	var c connect.Consensus
	if err := json.Unmarshal([]byte(consensusJSON), &c); err != nil {
		panic(err)
	}
	
	fmt.Printf("共识ID: %s\n", c.ID)
	fmt.Printf("共识标题: %s\n", c.Title)
}
```

### API 路由

```go
package main

import (
	"fmt"

	connect "github.com/quanttide/quanttide-connect-toolkit/packages/go/pkg"
)

func main() {
	// 构造API路径
	consensusID := "c0a80101-0000-0000-0000-000000000001"
	path := connect.ConsensusPath(consensusID)
	fmt.Printf("共识详情路径: %s\n", path)
	
	// 构造共识图节点路径
	graphID := "g0a80101-0000-0000-0000-000000000001"
	nodePath := connect.ConsensusGraphNodePath(graphID, consensusID)
	fmt.Printf("共识图节点路径: %s\n", nodePath)
}
```

### 领域事件

```go
package main

import (
	"encoding/json"
	"fmt"

	connect "github.com/quanttide/quanttide-connect-toolkit/packages/go/pkg"
)

func main() {
	// 创建共识事件
	event := connect.ConsensusCreatedEvent{
		Event: connect.Event{
			EventID:   "e0a80101-0000-0000-0000-000000000001",
			EventType: connect.EventConsensusCreated,
			Timestamp: "2026-08-28T14:30:00+08:00",
		},
		Data: connect.ConsensusCreatedData{
			ConsensusID: "c0a80101-0000-0000-0000-000000000001",
			Title:       "共识是沟通管理领域的核心概念",
			Description: "经过团队讨论，我们一致认为共识是沟通从分歧到统一的关键产出物。",
			CreatedAt:   "2026-08-28T14:30:00+08:00",
		},
	}
	
	jsonData, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("事件JSON: %s\n", jsonData)
}
```

## 数据模型

本包提供以下数据模型：

- **Message**：消息，沟通的基本载体
- **Consensus**：共识，团队成员通过讨论达成的一致决策或结论
- **ConsensusRelation**：共识关系，建立共识之间的逻辑关联
- **ConsensusGraph**：共识图，多个共识以 DAG 结构组织的集合

所有模型定义以 `docs/specification/content/consensus.md` 的领域模型定义为准。

## API 路由

本包提供以下 API 路由常量和构造函数：

- **静态资源集合路径**：`/messages`、`/consensuses`、`/consensus-relations`、`/consensus-graphs`
- **单资源路径**：`/messages/{id}`、`/consensuses/{id}`、`/consensus-relations/{id}`、`/consensus-graphs/{id}`
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
