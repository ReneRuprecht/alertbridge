package com.example.alertbridge.alerts.domain.exception;

public class InvalidAlertStatusException extends RuntimeException {
    public InvalidAlertStatusException(String message) {
        super(message);
    }
}
