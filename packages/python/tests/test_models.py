"""测试数据模型。"""

from datetime import datetime

import pytest
from pydantic import ValidationError

from quanttide_connect.models import (
    Consensus,
    ConsensusGraph,
    ConsensusRelation,
    ConsensusStatus,
    Message,
    MessageType,
    Relation,
)


class TestMessage:
    def test_create(self) -> None:
        msg = Message(content="你好", type=MessageType.user)
        assert msg.content == "你好"
        assert msg.type == MessageType.user
        assert isinstance(msg.id, str)
        assert len(msg.id) == 32
        assert isinstance(msg.created_at, datetime)

    def test_updated_at_none_by_default(self) -> None:
        msg = Message(content="test", type=MessageType.system)
        assert msg.updated_at is None


class TestConsensus:
    def test_create_proposed(self) -> None:
        c = Consensus(title="测试共识", description="共识的详细说明")
        assert c.title == "测试共识"
        assert c.description == "共识的详细说明"
        assert c.status == ConsensusStatus.proposed

    def test_status_enum(self) -> None:
        c = Consensus(title="test", status=ConsensusStatus.confirmed)
        assert c.status == ConsensusStatus.confirmed

    def test_deprecated(self) -> None:
        c = Consensus(title="test", status=ConsensusStatus.deprecated)
        assert c.status == ConsensusStatus.deprecated

    def test_content_field_is_not_supported(self) -> None:
        with pytest.raises(ValidationError):
            Consensus(title="新字段", content="旧字段")  # type: ignore[call-arg]


class TestRelation:
    def test_create(self) -> None:
        r = Relation(message_id="msg1", consensus_id="con1")
        assert r.message_id == "msg1"

    def test_unique_id(self) -> None:
        r1 = Relation(message_id="a", consensus_id="b")
        r2 = Relation(message_id="a", consensus_id="b")
        assert r1.id != r2.id


class TestConsensusRelation:
    def test_serializes_from_with_api_field_name(self) -> None:
        relation = ConsensusRelation(from_id="c1", to="c2", relation_type="支持")
        assert relation.model_dump(by_alias=True)["from"] == "c1"


class TestConsensusGraph:
    def test_create(self) -> None:
        graph = ConsensusGraph(
            name="团队决策",
            nodes=[Consensus(title="采用 SQLite")],
            edges=[],
        )
        assert graph.nodes[0].title == "采用 SQLite"
