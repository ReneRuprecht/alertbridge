package com.example.alertbridge.api.alertmanager.dto;

public record AlertDto(String fingerprint, AlertStatusDto status, AlertLabelsDto labels, String startsAt) {
}
