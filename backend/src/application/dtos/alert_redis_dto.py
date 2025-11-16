from __future__ import annotations

from pydantic import BaseModel

from src.domain.entities.alert import Alert


class AlertRedisDto(BaseModel):
    alertname: str
    status: str
    fingerprint: str
    instance: str

    @classmethod
    def from_domain(cls, alert: Alert) -> AlertRedisDto:
        return cls(
            alertname=alert.alertname,
            fingerprint=alert.fingerprint,
            status=alert.status,
            instance=alert.labels.get("instance", ""),
        )

    def to_json(self) -> str:
        return self.model_dump_json()

    @classmethod
    def from_json(cls, data: str) -> AlertRedisDto:
        return cls.model_validate_json(data)
