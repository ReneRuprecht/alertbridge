from __future__ import annotations

from dataclasses import dataclass, field
from typing import Any, Dict, Optional


@dataclass
class Alert:
    alertname: str
    status: str
    labels: Dict[str, str] = field(default_factory=dict)
    fingerprint: Optional[str] = None

    @staticmethod
    def from_json(alert_json: dict[str, Any]) -> Alert:
        labels = alert_json.get("labels", {})
        status = alert_json["status"]["state"]

        return Alert(
            alertname=labels.get("alertname", "unknown"),
            status=status,
            labels=labels,
            fingerprint=alert_json.get("fingerprint"),
        )
