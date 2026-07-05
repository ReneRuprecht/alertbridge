package com.example.alertbridge.alerts.domain.model;

import com.example.alertbridge.alerts.domain.exception.InvalidAlertStartsAtException;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertLabels;
import com.example.alertbridge.alerts.domain.value.AlertStatus;

import java.time.Instant;

public class Alert {

    private final AlertFingerprint fingerprint;
    private final AlertStatus status;
    private final AlertLabels labels;
    private final Instant startsAt;
    private final Instant receivedAt;

    Alert(AlertFingerprint fingerprint,
          AlertStatus status,
          AlertLabels labels,
          Instant startsAt,
          Instant receivedAt) {
        this.fingerprint = fingerprint;
        this.status = status;
        this.labels = labels;
        this.startsAt = startsAt;
        this.receivedAt = receivedAt;
    }

    public static Alert create(AlertFingerprint fingerprint,
                               AlertStatus status,
                               AlertLabels labels,
                               Instant startsAt) {

        if (startsAt == null) throw new InvalidAlertStartsAtException("StartsAt must not be null");

        return new Alert(fingerprint, status, labels, startsAt, Instant.now());
    }

    public AlertFingerprint fingerprint() {
        return this.fingerprint;
    }

    public AlertStatus status() {
        return this.status;
    }

    public AlertLabels labels() {
        return this.labels;
    }

    public Instant startsAt() {
        return this.startsAt;
    }

    public Instant receivedAt() {
        return this.receivedAt;
    }
}
