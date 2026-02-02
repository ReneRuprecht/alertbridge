package com.example.alertbridge.api.alertmanager.dto;

import com.fasterxml.jackson.annotation.JsonCreator;

public enum AlertStatus {
    FIRING, RESOLVED, UNKNOWN;

    @JsonCreator
    public static AlertStatus fromString(String value) {
        if (value == null) return UNKNOWN;

        return switch (value.toLowerCase()) {
            case "firing" -> FIRING;
            case "resolved" -> RESOLVED;
            default -> UNKNOWN;
        };
    }

}
