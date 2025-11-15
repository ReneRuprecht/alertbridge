import json
from datetime import datetime, timezone

from src.domain.entities.alert import Alert


def alerts_to_postgres_params(alerts: list[Alert]) -> list[tuple]:
    return [
        (
            alert.alertname,
            alert.status,
            json.dumps(alert.labels),
            alert.fingerprint,
            alert.starts_at,
            (
                datetime.now(timezone.utc)
                if (alert.status == "resolved" and not alert.ended_at)
                else alert.ended_at
            ),
            alert.updated_at,
        )
        for alert in alerts
    ]
