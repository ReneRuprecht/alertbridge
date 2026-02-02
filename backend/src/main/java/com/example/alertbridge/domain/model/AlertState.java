package com.example.alertbridge.domain.model;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.value.AlertFingerprint;
import com.example.alertbridge.domain.value.AlertLabels;
import com.example.alertbridge.domain.value.AlertStartsAt;
import com.example.alertbridge.domain.value.AlertStatus;

public class AlertState {
    public AlertFingerprint fingerprint;
    public AlertStatus status;
    public AlertLabels labels;
    public AlertStartsAt startsAt;

    private AlertState(AlertFingerprint fingerprint, AlertLabels labels, AlertStatus status, AlertStartsAt startsAt) {
        this.fingerprint = fingerprint;
        this.labels = labels;
        this.status = status;
        this.startsAt = startsAt;
    }

    public static AlertState fromEvent(AlertEvent event) {
        return new AlertState(event.fingerprint(), event.labels(), event.status(), event.startsAt());
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

    public boolean isActive() {
        return status.isFiring();
    }
}
