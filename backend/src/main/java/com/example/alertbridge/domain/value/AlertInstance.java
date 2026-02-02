package com.example.alertbridge.domain.value;

public record AlertInstance(String value) {
    public AlertInstance(String value) {
        this.value = (value == null || value.isBlank()) ? "unknown" : value;
    }
}
