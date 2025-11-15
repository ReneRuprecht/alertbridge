from datetime import datetime, timezone

import pytest
from pydantic import ValidationError

from src.domain.entities.alert import Alert


def test_alert_creation_valid():
    alert = Alert(
        alertname="HighCPU",
        status="firing",
        labels={"instance": "server1"},
        fingerprint="abc123",
        starts_at=datetime.now(timezone.utc),
        ended_at=None,
        updated_at=datetime.now(timezone.utc),
    )
    assert alert.alertname == "HighCPU"
    assert alert.status == "firing"
    assert alert.labels["instance"] == "server1"
    assert alert.ended_at is None


def test_alert_optional_dates_none():
    alert = Alert(
        alertname="LowDisk",
        status="resolved",
        labels={},
        fingerprint="abc123",
    )
    assert alert.starts_at is None
    assert alert.ended_at is None
    assert alert.updated_at is None


def test_alert_invalid_labels_type():
    with pytest.raises(ValidationError):
        Alert(
            alertname="InstanceDown",
            status="firing",
            labels="not-a-dict",
            fingerprint="abc123",
            starts_at=None,
            ended_at=None,
            updated_at=None,
        )


def test_alert_missing_required_field():
    with pytest.raises(ValidationError):
        Alert(status="firing", labels={}, fingerprint="abc123")
