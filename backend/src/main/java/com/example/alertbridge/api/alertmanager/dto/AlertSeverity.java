package com.example.alertbridge.api.alertmanager.dto;

import com.fasterxml.jackson.annotation.JsonCreator;

public enum AlertSeverity {
    CRITICAL, WARNING, INFO, UNKNOWN;

    @JsonCreator
    public static AlertSeverity fromValue(String value) {
        if (value == null) return UNKNOWN;

        return switch (value.toLowerCase()) {
            case "critical" -> CRITICAL;
            case "warning" -> WARNING;
            case "info" -> INFO;
            default -> UNKNOWN;
        };
    }
}
