package com.example.alertbridge.domain.value;

public record AlertFingerprint(String value) {

    public AlertFingerprint {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException("Fingerprint must not be null or blank");
        }
    }
}
