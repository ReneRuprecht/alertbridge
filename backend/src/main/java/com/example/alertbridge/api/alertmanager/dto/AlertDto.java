package com.example.alertbridge.api.alertmanager.dto;

public record AlertDto(String fingerprint, AlertStatus status, AlertLabelsDto labels, String startsAt) {
}
