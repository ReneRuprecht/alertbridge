package com.example.alertbridge.alerts.domain.exception;

public class InvalidAlertFingerprintException extends RuntimeException {
    public InvalidAlertFingerprintException(String message) {
        super(message);
    }
}
