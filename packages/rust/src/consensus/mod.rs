//! 共识子领域模块。
//!
//! 提供共识、共识关系、共识图的领域模型、API路由和领域事件定义。

pub mod api;
pub mod events;
pub mod models;

pub use api::*;
pub use events::*;
pub use models::*;
