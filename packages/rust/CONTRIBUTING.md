# CONTRIBUTING.md - quanttide-connect Rust 库

仓库约定与常见任务的贡献指南。定位、结构与领域模型事实源见 [AGENTS.md](AGENTS.md)。

## 发布标签规范（Rust 包）

本仓库所有发布标签遵循统一格式 `<scope>/vX.Y.Z`，如 `rust/v0.1.0-alpha.2`、`dart/v0.1.1`。**主标签一律挂 GitHub Release**，作为人可读的版本事实源。

## 版本号以契约为准

`tests/` 的 Schema 与 Fixture 是各语言包版本的**客观依据**：一个语言包能否升版，看它是否已与当前契约体系对齐并通过契约测试。

- 已落地契约测试并全绿的包（如 `rust/`）：可按 semver 正常发版；
- 契约测试挂起、模型未对齐契约体系的包（如迁移期的 python/dart）：只允许发 `alpha`/预览级版本，不得发正式版；
- 契约变更（Schema/Fixture/路由）先行，各语言跟进后才能携带该变更进入新版本。

## 提交规范

遵循 Conventional Commits（`feat:` / `fix:` / `docs:` / `chore:` 等）；破坏性变更标 `!` 并在 body 说明迁移方式。

## 子模块协作

本仓库作为子模块挂载于 quanttide-connect 的 `packages/quanttide-connect-toolkit`：

1. 在本仓库完成修改并提交推送；
2. 回父仓库 `git add packages/quanttide-connect-toolkit && git commit` 更新引用指针并推送。

## Rust 包开发规范

### 目录结构

```
packages/rust/
├── Cargo.toml          # Rust 包配置
└── src/
    ├── lib.rs          # 库入口
    └── consensus/      # 共识子领域
        ├── mod.rs      # 模块定义
        ├── models.rs   # 数据模型定义
        ├── api.rs      # API 路由定义
        └── events.rs   # 领域事件定义
```

### 代码规范

1. **命名规范**：遵循 Rust 官方命名规范，使用蛇形命名法（snake_case）
2. **注释规范**：所有公开函数和结构体必须有文档注释，遵循 Rust Doc 规范
3. **错误处理**：使用 Rust 标准错误处理方式，返回 Result 类型
4. **测试规范**：所有公开函数必须有对应的单元测试

### 依赖管理

使用 Cargo 管理依赖，确保 `Cargo.toml` 和 `Cargo.lock` 文件保持同步。

### 契约测试

契约测试是版本发布的客观依据，必须确保：

1. 所有契约测试通过
2. 数据模型与契约定义一致
3. API 接口符合契约规范

## 发布流程

1. **代码审查**：确保代码符合规范，所有测试通过
2. **契约验证**：运行契约测试，确保与契约体系对齐
3. **版本更新**：更新 CHANGELOG.md 和 Cargo.toml 中的版本号
4. **标签创建**：创建主标签和工具链别名标签
5. **发布执行**：推送标签，创建 GitHub Release

## 常见问题

### 编译失败

检查 Cargo.toml 中的依赖版本是否正确，运行 `cargo update` 更新依赖。

### 测试失败

检查数据模型是否与契约定义一致，参考 `src/consensus/models.rs` 中的测试用例。

### 版本号不一致

确保 CHANGELOG.md、Cargo.toml 和 Git 标签中的版本号保持一致。
