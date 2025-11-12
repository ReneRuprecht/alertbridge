from __future__ import annotations

import json
from dataclasses import dataclass
from datetime import datetime, timezone
from typing import Any, Dict, Optional


@dataclass
class Alert:
    alertname: str
    status: str
    labels: Dict[str, str]
    fingerprint: str

    starts_at: Optional[datetime]
    ended_at: Optional[datetime]
    updated_at: Optional[datetime]

    @staticmethod
    def from_json(alert_json: dict[str, Any]) -> Alert:
        """
        Create an Alert object from Alertmanager JSON data.
        """

        labels = alert_json.get("labels", {})
        status = ""
        status_field = alert_json.get("status", "")
        if isinstance(status_field, dict):
            status = status_field.get("state", "")
            # Change status to firing,
            # Alertmanager sends 'active' on /api/v2/alerts
            if status == "active":
                status = "firing"
        elif isinstance(status_field, str):
            status = status_field

        if status == "resolved":
            ended_at = datetime.now(timezone.utc)

        def parse_datetime(dt_str: Optional[str]) -> Optional[datetime]:
            if not dt_str:
                return None
            try:
                return datetime.fromisoformat(dt_str.replace("Z", "+00:00"))
            except Exception:
                return None

        starts_at = parse_datetime(alert_json.get("startsAt"))
        updated_at = parse_datetime(
            alert_json.get("updatedAt", datetime.now(timezone.utc))
        )
        ended_at = None

        return Alert(
            alertname=labels.get("alertname", "unknown"),
            status=status,
            labels=labels,
            fingerprint=str(alert_json.get("fingerprint")),
            starts_at=starts_at,
            ended_at=ended_at,
            updated_at=updated_at,
        )

    @staticmethod
    def from_psql(record: tuple) -> Alert:
        """
        Create an Alert object from database response.
        """
        (
            _,
            alertname,
            status,
            labels_json,
            fingerprint,
            starts_at,
            ended_at,
            updated_at,
        ) = record

        labels = (
            labels_json
            if isinstance(labels_json, dict)
            else json.loads(labels_json)
        )

        def ensure_tz(dt: Optional[datetime]) -> Optional[datetime]:
            if dt is None:
                return None
            if dt.tzinfo is None:
                return dt.replace(tzinfo=timezone.utc)
            return dt

        return Alert(
            alertname=alertname,
            status=status,
            labels=labels,
            fingerprint=fingerprint,
            starts_at=ensure_tz(starts_at),
            ended_at=ensure_tz(ended_at),
            updated_at=ensure_tz(updated_at),
        )

    def to_redis_minimal_dict(self) -> dict:
        return {
            "alertname": self.alertname,
            "status": self.status,
            "fingerprint": self.fingerprint,
            "instance": self.labels.get("instance", ""),
        }
