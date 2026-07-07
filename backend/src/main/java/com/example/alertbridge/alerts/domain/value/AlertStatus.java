package com.example.alertbridge.alerts.domain.value;

import com.example.alertbridge.alerts.domain.exception.InvalidAlertStatusException;

public enum AlertStatus {

    FIRING("firing"), RESOLVED("resolved");

    private final String value;

    AlertStatus(String value) {
        this.value = value;
    }

    public static AlertStatus of(String value) {
        if (value == null || value.isBlank()) {
            throw new InvalidAlertStatusException("Status must not be null or blank");
        }

        String cleanValue = value.trim().toLowerCase();

        for (AlertStatus status : values()) {
            if (status.value.equals(cleanValue)) {
                return status;
            }
        }

        throw new InvalidAlertStatusException("Invalid Status: " + value);
    }

}
