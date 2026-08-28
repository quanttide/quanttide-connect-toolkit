package consensus

import "strings"

// 沟通管理域 API 路由模板（唯一事实源：docs/specification/content/consensus.md API 规格）。
//
// Route* 常量含 {参数} 通配段，段名即 API 规格定义的路径变量名，
// 服务端以 Method + 常量注册路由；消费方用下方构造函数生成具体路径。
// 构造函数与常量在同一文件内推导，保证永不漂移。
const (
	// 静态资源集合路径。
	RouteMessages          = "/messages"
	RouteConsensuses       = "/consensuses"
	RouteConsensusRelations = "/consensus-relations"
	RouteConsensusGraphs   = "/consensus-graphs"

	// 单资源路径（{id}）。
	RouteMessage          = "/messages/{id}"
	RouteConsensus        = "/consensuses/{id}"
	RouteConsensusRelation = "/consensus-relations/{id}"
	RouteConsensusGraph   = "/consensus-graphs/{id}"

	// 子资源路径。
	RouteConsensusRelationsList  = "/consensuses/{id}/relations"
	RouteConsensusGraphNodes     = "/consensus-graphs/{id}/nodes"
	RouteConsensusGraphNode      = "/consensus-graphs/{id}/nodes/{consensus_id}"
	RouteConsensusGraphEdges     = "/consensus-graphs/{id}/edges"
	RouteConsensusGraphEdge      = "/consensus-graphs/{id}/edges/{relation_id}"
	RouteConsensusGraphPaths     = "/consensus-graphs/{id}/paths"
)

// fill 将 route 模板中的 {key} 通配段替换为 value。
func fill(route, key, value string) string {
	return strings.ReplaceAll(route, "{"+key+"}", value)
}

// MessagePath 构造指定消息的路径。
func MessagePath(id string) string { return fill(RouteMessage, "id", id) }

// ConsensusPath 构造指定共识的路径。
func ConsensusPath(id string) string { return fill(RouteConsensus, "id", id) }

// ConsensusRelationPath 构造指定共识关系的路径。
func ConsensusRelationPath(id string) string { return fill(RouteConsensusRelation, "id", id) }

// ConsensusGraphPath 构造指定共识图的路径。
func ConsensusGraphPath(id string) string { return fill(RouteConsensusGraph, "id", id) }

// ConsensusRelationsPath 构造共识关系列表的路径。
func ConsensusRelationsPath(id string) string { return fill(RouteConsensusRelationsList, "id", id) }

// ConsensusGraphNodesPath 构造共识图节点列表的路径。
func ConsensusGraphNodesPath(id string) string { return fill(RouteConsensusGraphNodes, "id", id) }

// ConsensusGraphNodePath 构造共识图指定节点的路径。
func ConsensusGraphNodePath(id, consensusID string) string {
	return fill(fill(RouteConsensusGraphNode, "id", id), "consensus_id", consensusID)
}

// ConsensusGraphEdgesPath 构造共识图边列表的路径。
func ConsensusGraphEdgesPath(id string) string { return fill(RouteConsensusGraphEdges, "id", id) }

// ConsensusGraphEdgePath 构造共识图指定边的路径。
func ConsensusGraphEdgePath(id, relationID string) string {
	return fill(fill(RouteConsensusGraphEdge, "id", id), "relation_id", relationID)
}

// ConsensusGraphPathsPath 构造共识图路径查询的路径。
func ConsensusGraphPathsPath(id string) string { return fill(RouteConsensusGraphPaths, "id", id) }
