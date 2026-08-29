"""
数据模型：Message、Consensus、Relation。

对应 connect-agent DRD 定义，无冗余字段。
"""

from __future__ import annotations

import uuid
from datetime import UTC, datetime
from enum import Enum

from pydantic import BaseModel, ConfigDict, Field


def _utcnow() -> datetime:
    return datetime.now(UTC)


def _new_id() -> str:
    return uuid.uuid4().hex


class MessageType(str, Enum):
    user = "user"
    agent = "agent"
    system = "system"


class Message(BaseModel):
    """对话消息，按时间排序。"""

    id: str = Field(default_factory=_new_id)
    content: str
    type: MessageType
    created_at: datetime = Field(default_factory=_utcnow)
    updated_at: datetime | None = None


class ConsensusStatus(str, Enum):
    proposed = "proposed"
    confirmed = "confirmed"
    deprecated = "deprecated"


class Consensus(BaseModel):
    """从消息中提炼出的共识。"""

    model_config = ConfigDict(extra="forbid")

    id: str = Field(default_factory=_new_id)
    title: str
    description: str = ""
    status: ConsensusStatus = ConsensusStatus.proposed
    created_at: datetime = Field(default_factory=_utcnow)
    updated_at: datetime | None = None


class ConsensusRelation(BaseModel):
    """共识之间的逻辑关联。"""

    model_config = ConfigDict(populate_by_name=True)

    id: str = Field(default_factory=_new_id)
    from_id: str = Field(alias="from")
    to: str
    relation_type: str


class ConsensusGraph(BaseModel):
    """以有向图组织多个共识及其关系。"""

    id: str = Field(default_factory=_new_id)
    name: str
    description: str = ""
    nodes: list[Consensus] = Field(default_factory=list)
    edges: list[ConsensusRelation] = Field(default_factory=list)
    created_at: datetime = Field(default_factory=_utcnow)
    updated_at: datetime | None = None


class Relation(BaseModel):
    """消息与共识之间的多对多溯源关联。"""

    id: str = Field(default_factory=_new_id)
    message_id: str
    consensus_id: str
