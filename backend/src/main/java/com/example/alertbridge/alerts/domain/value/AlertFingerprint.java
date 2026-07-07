package com.example.alertbridge.alerts.domain.value;

import com.example.alertbridge.alerts.domain.exception.InvalidAlertFingerprintException;

public record AlertFingerprint(String value) {

    public AlertFingerprint {
        if (value == null || value.isBlank()) {
            throw new InvalidAlertFingerprintException("Fingerprint must not be null or blank");
        }
    }
}
