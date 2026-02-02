package com.example.alertbridge.api.alertmanager.dto;

import com.fasterxml.jackson.annotation.JsonCreator;

public enum AlertStatusDto {
    FIRING, RESOLVED, UNKNOWN;

    @JsonCreator
    public static AlertStatusDto fromString(String value) {
        if (value == null) return UNKNOWN;

        return switch (value.toLowerCase()) {
            case "firing" -> FIRING;
            case "resolved" -> RESOLVED;
            default -> UNKNOWN;
        };
    }

}
