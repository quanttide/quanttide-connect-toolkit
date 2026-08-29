# Changelog

## [Unreleased]

- 增加 Go 模块解析别名、Rust crates.io 和 Dart pub.dev 发布工作流，支持 tag 自动发布与手动补发。

## [2026-08-29]

- Go 包准备发布 v0.1.0，保留 `go/v0.1.0` 与 `packages/go/v0.1.0` 双标签规则。
- Rust 包准备发布 v0.1.0，并移除误入版本库的 `packages/rust/target/` 构建产物。
- Dart 包补齐共识模型与路由，准备发布 v0.1.0。
- Python 包升级到 v0.2.0，完成 `role -> type` 与 `content -> title/description` 破坏性迁移。
