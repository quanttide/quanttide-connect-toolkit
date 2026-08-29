/// 量潮沟通工程 Dart SDK。
library;

const String routeConsensuses = '/consensuses';

/// 构造共识详情路径。
String consensusPath(String id) => '$routeConsensuses/$id';

/// 共识领域模型。
class Consensus {
  const Consensus({
    required this.id,
    required this.title,
    this.description = '',
    required this.createdAt,
    this.updatedAt,
  });

  factory Consensus.fromJson(Map<String, dynamic> json) {
    return Consensus(
      id: json['id'] as String,
      title: json['title'] as String,
      description: json['description'] as String? ?? '',
      createdAt: DateTime.parse(json['created_at'] as String),
      updatedAt: json['updated_at'] == null
          ? null
          : DateTime.parse(json['updated_at'] as String),
    );
  }

  final String id;
  final String title;
  final String description;
  final DateTime createdAt;
  final DateTime? updatedAt;

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'title': title,
      'description': description,
      'created_at': createdAt.toIso8601String(),
      if (updatedAt != null) 'updated_at': updatedAt!.toIso8601String(),
    };
  }
}

/// 共识之间的逻辑关联。
class ConsensusRelation {
  const ConsensusRelation({
    required this.id,
    required this.from,
    required this.to,
    required this.relationType,
  });

  factory ConsensusRelation.fromJson(Map<String, dynamic> json) {
    return ConsensusRelation(
      id: json['id'] as String,
      from: json['from'] as String,
      to: json['to'] as String,
      relationType: json['relation_type'] as String,
    );
  }

  final String id;
  final String from;
  final String to;
  final String relationType;

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'from': from,
      'to': to,
      'relation_type': relationType,
    };
  }
}

/// 由共识节点和关系边组成的共识图。
class ConsensusGraph {
  const ConsensusGraph({
    required this.id,
    required this.name,
    this.description = '',
    this.nodes = const [],
    this.edges = const [],
    required this.createdAt,
    this.updatedAt,
  });

  final String id;
  final String name;
  final String description;
  final List<Consensus> nodes;
  final List<ConsensusRelation> edges;
  final DateTime createdAt;
  final DateTime? updatedAt;
}
