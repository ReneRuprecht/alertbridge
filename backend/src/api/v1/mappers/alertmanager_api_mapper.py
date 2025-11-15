from datetime import datetime, timezone

from src.domain.entities.alert import Alert


async def alertmanager_payload_to_alerts(payload: dict) -> list[Alert]:
    alerts_json = payload.get("alerts", [])

    if not isinstance(alerts_json, list) or not alerts_json:
        return []

    alerts: list[Alert] = []
    for alert_json in alerts_json:

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

        starts_at = alert_json.get("startsAt", None)
        updated_at = alert_json.get("updatedAt", None)
        ended_at = None

        alerts.append(
            Alert(
                alertname=labels.get("alertname", "unknown"),
                status=status,
                labels=labels,
                fingerprint=str(alert_json.get("fingerprint")),
                starts_at=starts_at,
                ended_at=ended_at,
                updated_at=updated_at,
            )
        )

    return alerts
