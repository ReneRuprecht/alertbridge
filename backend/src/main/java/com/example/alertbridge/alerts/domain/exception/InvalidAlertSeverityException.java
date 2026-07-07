package com.example.alertbridge.alerts.domain.exception;

public class InvalidAlertSeverityException extends RuntimeException {
    public InvalidAlertSeverityException(String message) {
        super(message);
    }
}
