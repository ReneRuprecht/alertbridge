package com.example.alertbridge.domain.value;

public enum AlertStatus {
    FIRING,
    RESOLVED,
    UNKNOWN;

    public boolean isFiring() {
        return this == FIRING;
    }

    public boolean isResolved() {
        return this == RESOLVED;
    }

}
