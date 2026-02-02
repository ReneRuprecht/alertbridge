package com.example.alertbridge.domain.value;

public record AlertEnvironment(String value) {

    public AlertEnvironment(String value) {
        this.value = (value == null || value.isBlank()) ? "unknown" : value;
    }
}