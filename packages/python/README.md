# quanttide-connect (Python)

沟通工程工具包 Python SDK。提供人机沟通的核心数据模型和共识管理能力，供 `apps/` 中的各应用复用。

## 安装

```bash
uv add quanttide-connect
```

## v0.2.0 迁移

- `Message.role` / `Role` 已改为 `Message.type` / `MessageType`。
- `Consensus.content` 已改为 `Consensus.title` 和 `Consensus.description`。
- `Consensus.status` 仍是服务层生命周期扩展字段；跨语言基础模型以 `title`、`description` 和时间戳为准。
- `ConsensusService.propose("标题", "描述")` 替代旧的 `propose("内容")`。

## 模块

- `quanttide_connect.models` — Message、Consensus、Relation 数据模型
- `quanttide_connect.events` — 领域事件定义
- `quanttide_connect.repository` — 仓储接口（Repository Protocol）
- `quanttide_connect.services` — 应用服务
  - `services.message` — `MessageService`（发送、编辑消息）
  - `services.consensus` — `ConsensusService`（提议、确认、废弃共识）
  - `services.relation` — `RelationService`（关联、解除关联）
