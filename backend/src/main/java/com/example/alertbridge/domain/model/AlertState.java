package com.example.alertbridge.domain.model;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.value.*;

public class AlertState {
    public AlertFingerprint fingerprint;
    public AlertStatus status;
    public AlertLabels labels;
    public AlertStartsAt startsAt;
    public AlertReceivedAt receivedAt;

    private AlertState(AlertFingerprint fingerprint,
                       AlertLabels labels,
                       AlertStatus status,
                       AlertStartsAt startsAt,
                       AlertReceivedAt receivedAt) {
        this.fingerprint = fingerprint;
        this.labels = labels;
        this.status = status;
        this.startsAt = startsAt;
        this.receivedAt = receivedAt;
    }

    public static AlertState fromEvent(AlertEvent event) {
        return new AlertState(
                event.fingerprint(),
                event.labels(),
                event.status(),
                event.startsAt(),
                event.receivedAt()
        );
    }

    public void apply(AlertEvent event) {
        if (!event.fingerprint().equals(this.fingerprint)) {
            throw new IllegalArgumentException("Cannot apply event for different alert");
        }
        this.status = event.status();
    }

    public AlertFingerprint fingerprint() {
        return fingerprint;
    }

    public AlertLabels labels() {
        return labels;
    }

    public AlertStatus status() {
        return status;
    }

    public AlertStartsAt startsAt() {
        return startsAt;
    }

    public AlertReceivedAt receivedAt() {
        return receivedAt;
    }

    public boolean isActive() {
        return status.isFiring();
    }
}
