# 量潮沟通工程工具箱

沟通工程工具包，按语言拆分为 Go、Rust、Dart 和 Python 包。

## 当前发布目标

- Go：`go/v0.1.0` 与 `packages/go/v0.1.0`。
- Rust：`rust/v0.1.0`。
- Dart：`dart/v0.1.0`。
- Python：`python/v0.2.0`，包含破坏性模型迁移。

## 发布自动化

- Go 的 `go/vX.Y.Z` tag 会自动生成同一提交的 `packages/go/vX.Y.Z` 模块解析标签。
- Rust 的 `rust/vX.Y.Z` tag 会运行质量门禁并发布到 crates.io。
- Dart 的 `dart/vX.Y.Z` tag 会运行质量门禁并发布到 pub.dev。
- Python 的 `python/vX.Y.Z` GitHub Release 会触发 PyPI 发布。

Rust 和 Dart 的发布工作流需要仓库 Actions secrets：`CRATES_API_TOKEN` 和
`PUB_DEV_TOKEN`。已存在的版本可以从 GitHub Actions 手动运行对应 workflow，输入完整
release tag 进行补发。
