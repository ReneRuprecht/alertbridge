package com.example.alertbridge.alerts.domain.value;

import com.example.alertbridge.alerts.domain.exception.InvalidAlertSeverityException;

public enum AlertSeverity {

    CRITICAL("critical"), WARNING("warning"), INFO("info");

    private final String value;

    AlertSeverity(String value) {
        this.value = value;
    }

    public static AlertSeverity of(String value) {
        if (value == null || value.isBlank()) {
            throw new InvalidAlertSeverityException("Severity must not be null or blank");
        }

        String cleanValue = value.trim().toLowerCase();

        for (AlertSeverity severity : values()) {
            if (severity.value.equals(cleanValue)) {
                return severity;
            }
        }
        throw new InvalidAlertSeverityException("Invalid Severity: " + value);
    }

    public String value() {
        return this.value;
    }


}
