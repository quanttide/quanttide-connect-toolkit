# AGENTS.md - quanttide-connect Rust 库

## 项目结构

```
packages/rust/
├── AGENTS.md           # 本文件
├── CHANGELOG.md        # 版本变更记录
├── CONTRIBUTING.md     # 贡献指南
├── README.md           # 项目说明
├── Cargo.toml          # Rust 包配置
└── src/
    ├── lib.rs          # 库入口
    └── consensus/      # 共识子领域
        ├── mod.rs      # 模块定义
        ├── models.rs   # 数据模型定义
        ├── api.rs      # API 路由定义
        └── events.rs   # 领域事件定义
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
cargo test
```

运行文档测试：

```bash
cargo test --doc
```

## 发布

发布流程参考 [CONTRIBUTING.md](CONTRIBUTING.md) 的发布流程部分。
