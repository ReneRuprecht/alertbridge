package com.example.alertbridge.api.alertmanager.dto;

import com.fasterxml.jackson.annotation.JsonCreator;

public enum AlertSeverityDto {
    CRITICAL, WARNING, INFO, UNKNOWN;

    @JsonCreator
    public static AlertSeverityDto fromString(String value) {
        if (value == null) return UNKNOWN;

        return switch (value.toLowerCase()) {
            case "critical" -> CRITICAL;
            case "warning" -> WARNING;
            case "info" -> INFO;
            default -> UNKNOWN;
        };
    }
}
