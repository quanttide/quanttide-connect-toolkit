# Changelog

## [0.2.0] - 2026-08-29

### Added

- 增加 `ConsensusRelation` 与 `ConsensusGraph` 模型，对齐工程标准 v0.1.1。

### Changed

- 破坏性变更：`MessageService.send` 使用 `MessageType` 与 `type` 字段，不再使用旧 `Role`。
- 破坏性变更：`Consensus` 使用 `title` 和 `description` 字段，不再使用旧 `content`。
- `ConsensusService.propose` 改为 `propose(title, description="", related_message_ids=None)`。
- 共识事件载荷改为发布 `title` 和 `description`。

## [0.1.0] - 2026-05-17

- 初始化 Python SDK
- 数据模型：Message、Consensus、Relation
- 领域事件系统（EventBus）
- 仓储接口（Repository Protocol）
- 应用服务：MessageService、ConsensusService、RelationService
