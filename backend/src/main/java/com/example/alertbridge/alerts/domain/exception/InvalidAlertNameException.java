package com.example.alertbridge.alerts.domain.exception;

public class InvalidAlertNameException extends RuntimeException {
    public InvalidAlertNameException(String message) {
        super(message);
    }
}
