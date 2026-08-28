# AGENTS.md - quanttide-connect-toolkit Go 包

## 项目结构

```
packages/go/
├── AGENTS.md           # 本文件
├── CHANGELOG.md        # 版本变更记录
├── CONTRIBUTING.md     # 贡献指南
├── README.md           # 项目说明
├── go.mod              # Go 模块定义
└── pkg/                # 源代码
    ├── models.go       # 数据模型定义
    ├── api.go          # API 路由定义
    ├── events.go       # 领域事件定义
    └── models_test.go  # 模型测试
```

## 事实源

- **领域模型**：`docs/specification/content/consensus.md` 的领域模型定义
- **API 规格**：`docs/specification/content/consensus.md` 的 API 规格
- **领域事件**：`docs/specification/content/consensus.md` 的领域事件

## 提交约定

遵循 Conventional Commits（`feat:` / `fix:` / `docs:` / `chore:` 等）；破坏性变更标 `!` 并在 body 说明迁移方式。

## 测试

运行测试：

```bash
go test ./pkg/...
```

运行契约测试：

```bash
go test ./pkg/ -run TestContract
```

## 发布

发布流程参考 [CONTRIBUTING.md](CONTRIBUTING.md) 的发布流程部分。
