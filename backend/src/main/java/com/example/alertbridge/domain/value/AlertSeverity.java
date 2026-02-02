package com.example.alertbridge.domain.value;

public enum AlertSeverity {
    CRITICAL,
    WARNING,
    INFO,
    UNKNOWN;

    public static AlertSeverity fromString(String value) {
        if (value == null) return UNKNOWN;
        return switch (value.toLowerCase()) {
            case "critical" -> CRITICAL;
            case "warning" -> WARNING;
            case "info" -> INFO;
            default -> UNKNOWN;
        };
    }
}
