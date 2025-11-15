from datetime import datetime, timezone
from typing import Optional

from src.domain.entities.alert import Alert


def external_alert_json_to_domain(alert_json: dict) -> Alert:
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
