# Changelog

## [v0.1.0] - 2026-08-29

### Added

- 发布正式版 Go 工具包，提供 `pkg/consensus` 共识模型、API 路由和领域事件定义。

### Changed

- 以 `packages/go/v0.1.0` 作为 Go 工具链解析标签，`go/v0.1.0` 作为主发布标签。

## [v0.1.0-alpha.2] - 2026-08-28

### Changed

- 重构包结构：从单一 pkg 包改为按子领域组织（pkg/consensus）
- 更新导入路径：`github.com/quanttide/quanttide-connect-toolkit/packages/go/pkg` → `github.com/quanttide/quanttide-connect-toolkit/packages/go/pkg/consensus`

## [v0.1.0-alpha.1] - 2026-08-28

### Added

- 初始化 Go 包结构
- 增加数据模型定义：Consensus、ConsensusRelation、ConsensusGraph
- 增加 API 路由常量和构造函数
- 增加领域事件类型定义
- 增加模型测试
- 增加 README.md 和 CONTRIBUTING.md

### Changed

- 无

### Fixed

- 无

### Removed

- 无
