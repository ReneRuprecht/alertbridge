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
    ends_at: Optional[datetime]

    @staticmethod
    def from_json(alert_json: dict[str, Any]) -> Alert:
        """
        Create an Alert object from Alertmanager JSON data.
        """

        labels = alert_json.get("labels", {})
        status = alert_json["status"]["state"]

        def parse_datetime(dt_str: Optional[str]) -> Optional[datetime]:
            if not dt_str:
                return None
            try:
                return datetime.fromisoformat(dt_str.replace("Z", "+00:00"))
            except Exception:
                return None

        starts_at = parse_datetime(alert_json.get("startsAt"))
        ends_at = parse_datetime(alert_json.get("endsAt"))

        return Alert(
            alertname=labels.get("alertname", "unknown"),
            status=status,
            labels=labels,
            fingerprint=str(alert_json.get("fingerprint")),
            starts_at=starts_at,
            ends_at=ends_at,
        )

    @staticmethod
    def from_psql(record: tuple) -> Alert:
        """
        Create an Alert object from database response.
        """
        _, alertname, status, labels_json, fingerprint, starts_at, ends_at = record

        labels = labels_json if isinstance(labels_json, dict) else json.loads(labels_json)

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
            ends_at=ensure_tz(ends_at),
        )
