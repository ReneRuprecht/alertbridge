import json

from src.domain.entities.alert import Alert


def row_to_alert(record: tuple) -> Alert:
    (
        _id,
        alertname,
        status,
        labels_json,
        fingerprint,
        starts_at,
        ended_at,
        updated_at,
    ) = record

    try:
        labels = (
            labels_json
            if isinstance(labels_json, dict)
            else json.loads(labels_json)
        )
    except json.JSONDecodeError:
        labels = {}

    return Alert(
        alertname=alertname,
        status=status,
        labels=labels,
        fingerprint=fingerprint,
        starts_at=starts_at,
        ended_at=ended_at,
        updated_at=updated_at,
    )
