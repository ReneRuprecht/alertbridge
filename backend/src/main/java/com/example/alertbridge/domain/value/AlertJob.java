package com.example.alertbridge.domain.value;

public record AlertJob(String value) {

    public AlertJob(String value) {
        this.value = (value == null || value.isBlank()) ? "unknown" : value;
    }

}
