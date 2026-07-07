package com.example.alertbridge.alerts.domain.model;

import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;

import java.time.Instant;

public record CurrentAlert(AlertFingerprint fingerprint,
                           AlertStatus status,
                           String alertName,
                           AlertSeverity severity,
                           String environment,
                           String instance,
                           String job,
                           Instant startsAt,
                           Instant lastUpdatedAt) {
}
