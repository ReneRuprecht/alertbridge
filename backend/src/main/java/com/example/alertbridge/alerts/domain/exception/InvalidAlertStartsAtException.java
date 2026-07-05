package com.example.alertbridge.alerts.domain.exception;

public class InvalidAlertStartsAtException extends RuntimeException {
    public InvalidAlertStartsAtException(String message) {
        super(message);
    }
}
