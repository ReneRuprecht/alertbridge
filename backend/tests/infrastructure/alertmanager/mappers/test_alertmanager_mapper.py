from datetime import datetime, timezone

from src.infrastructure.alertmanager.mappers.alertmanager_mapper import (
    external_alert_json_to_domain,
)

sample_payload = {
    "receiver": "webhook-handler",
    "status": "firing",
    "alerts": [
        {
            "status": "firing",
            "labels": {
                "alertname": "InstanceDown",
                "env": "prod",
                "environment": "monitoring",
                "instance": "loadbalancer:9100",
                "job": "node_exporter",
                "severity": "critical",
            },
            "annotations": {},
            "startsAt": "2025-11-15T21:11:26.872Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "fingerprint": "abc123",
        },
        {
            "status": "resolved",
            "labels": {
                "alertname": "InstanceDown",
                "env": "prod",
                "environment": "monitoring",
                "instance": "monitoring:9100",
                "job": "node_exporter",
                "severity": "critical",
            },
            "annotations": {},
            "startsAt": "2025-11-16T00:38:56.872Z",
            "endsAt": "2025-11-16T00:44:56.872Z",
            "fingerprint": "abc123",
        },
    ],
}


def test_external_alert_json_to_domain_status():
    alerts_json = sample_payload["alerts"]
    alerts = [external_alert_json_to_domain(a) for a in alerts_json]

    assert len(alerts) == 2

    alert1 = alerts[0]
    assert alert1.status == "firing"
    assert alert1.alertname == "InstanceDown"
    assert alert1.fingerprint == "abc123"
    assert alert1.labels["instance"] == "loadbalancer:9100"
    assert isinstance(alert1.starts_at, datetime)
    assert alert1.starts_at.tzinfo == timezone.utc
    assert isinstance(alert1.ended_at, datetime)
    assert alert1.ended_at.tzinfo == timezone.utc

    alert2 = alerts[1]
    assert alert2.status == "resolved"
    assert alert2.fingerprint == "abc123"
    assert alert2.labels["instance"] == "monitoring:9100"
    assert isinstance(alert2.starts_at, datetime)
    assert isinstance(alert2.ended_at, datetime)


def test_external_alert_json_to_domain_missing_fields():
    payload = {
        "receiver": "webhook-handler",
        "status": "firing",
        "alerts": [
            {"status": "firing", "labels": {}, "fingerprint": "abc123"}
        ],
    }
    alert = external_alert_json_to_domain(alert_json=payload)
    assert alert.alertname == "unknown"
    assert alert.labels == {}
    assert alert.starts_at is None
    assert alert.ended_at is None
    assert isinstance(alert.updated_at, datetime)


def test_external_alert_json_to_domain_invalid_dates():
    payload = {
        "receiver": "webhook-handler",
        "status": "firing",
        "alerts": [
            {
                "status": "firing",
                "labels": {
                    "alertname": "InstanceDown",
                },
                "startsAt": "invalid",
                "endsAt": "invalid",
                "fingerprint": "abc123",
            },
        ],
    }
    alert = external_alert_json_to_domain(alert_json=payload)
    assert alert.starts_at is None
    assert alert.ended_at is None


def test_external_alert_json_to_domain_status_active():
    payload = {
        "receiver": "webhook-handler",
        "status": "firing",
        "alerts": [
            {
                "status": "active",
                "labels": {
                    "alertname": "InstanceDown",
                },
                "startsAt": "2025-11-16T00:38:56.872Z",
                "endsAt": "2025-11-16T00:44:56.872Z",
                "fingerprint": "abc123",
            },
        ],
    }
    alert = external_alert_json_to_domain(alert_json=payload)
    assert alert.status == "firing"
