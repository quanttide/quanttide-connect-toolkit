//! 共识领域 API 路由定义。
//!
//! 路由模板以 docs/specification/content/consensus.md 的 API 规格为准，
//! 服务端以 Method + 常量注册路由；消费方用下方构造函数生成具体路径。

/// 静态资源集合路径
pub const ROUTE_CONSENSUSES: &str = "/consensuses";
pub const ROUTE_CONSENSUS_RELATIONS: &str = "/consensus-relations";
pub const ROUTE_CONSENSUS_GRAPHS: &str = "/consensus-graphs";

/// 单资源路径（{id}）
pub const ROUTE_CONSENSUS: &str = "/consensuses/{id}";
pub const ROUTE_CONSENSUS_RELATION: &str = "/consensus-relations/{id}";
pub const ROUTE_CONSENSUS_GRAPH: &str = "/consensus-graphs/{id}";

/// 子资源路径
pub const ROUTE_CONSENSUS_RELATIONS_LIST: &str = "/consensuses/{id}/relations";
pub const ROUTE_CONSENSUS_GRAPH_NODES: &str = "/consensus-graphs/{id}/nodes";
pub const ROUTE_CONSENSUS_GRAPH_NODE: &str = "/consensus-graphs/{id}/nodes/{consensus_id}";
pub const ROUTE_CONSENSUS_GRAPH_EDGES: &str = "/consensus-graphs/{id}/edges";
pub const ROUTE_CONSENSUS_GRAPH_EDGE: &str = "/consensus-graphs/{id}/edges/{relation_id}";
pub const ROUTE_CONSENSUS_GRAPH_PATHS: &str = "/consensus-graphs/{id}/paths";

/// 将 route 模板中的 {key} 通配段替换为 value。
fn fill(route: &str, key: &str, value: &str) -> String {
    route.replace(&format!("{{{}}}", key), value)
}

/// 构造指定共识的路径。
pub fn consensus_path(id: &str) -> String {
    fill(ROUTE_CONSENSUS, "id", id)
}

/// 构造指定共识关系的路径。
pub fn consensus_relation_path(id: &str) -> String {
    fill(ROUTE_CONSENSUS_RELATION, "id", id)
}

/// 构造指定共识图的路径。
pub fn consensus_graph_path(id: &str) -> String {
    fill(ROUTE_CONSENSUS_GRAPH, "id", id)
}

/// 构造共识关系列表的路径。
pub fn consensus_relations_path(id: &str) -> String {
    fill(ROUTE_CONSENSUS_RELATIONS_LIST, "id", id)
}

/// 构造共识图节点列表的路径。
pub fn consensus_graph_nodes_path(id: &str) -> String {
    fill(ROUTE_CONSENSUS_GRAPH_NODES, "id", id)
}

/// 构造共识图指定节点的路径。
pub fn consensus_graph_node_path(id: &str, consensus_id: &str) -> String {
    fill(&fill(ROUTE_CONSENSUS_GRAPH_NODE, "id", id), "consensus_id", consensus_id)
}

/// 构造共识图边列表的路径。
pub fn consensus_graph_edges_path(id: &str) -> String {
    fill(ROUTE_CONSENSUS_GRAPH_EDGES, "id", id)
}

/// 构造共识图指定边的路径。
pub fn consensus_graph_edge_path(id: &str, relation_id: &str) -> String {
    fill(&fill(ROUTE_CONSENSUS_GRAPH_EDGE, "id", id), "relation_id", relation_id)
}

/// 构造共识图路径查询的路径。
pub fn consensus_graph_paths_path(id: &str) -> String {
    fill(ROUTE_CONSENSUS_GRAPH_PATHS, "id", id)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_consensus_path() {
        assert_eq!(consensus_path("test-id"), "/consensuses/test-id");
    }

    #[test]
    fn test_consensus_graph_node_path() {
        assert_eq!(
            consensus_graph_node_path("graph-id", "consensus-id"),
            "/consensus-graphs/graph-id/nodes/consensus-id"
        );
    }

    #[test]
    fn test_consensus_graph_edge_path() {
        assert_eq!(
            consensus_graph_edge_path("graph-id", "relation-id"),
            "/consensus-graphs/graph-id/edges/relation-id"
        );
    }
}
