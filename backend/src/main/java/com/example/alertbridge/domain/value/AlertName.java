package com.example.alertbridge.domain.value;

public record AlertName(String value) {

    public AlertName {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException("AlertName must not be null or blank");
        }
    }
}
