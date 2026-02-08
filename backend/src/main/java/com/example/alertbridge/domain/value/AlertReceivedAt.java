package com.example.alertbridge.domain.value;

import java.time.Instant;

public record AlertReceivedAt(Instant value) {

    public AlertReceivedAt {
        if (value == null) {
            throw new IllegalArgumentException("AlertStartsAt must not be null or blank");
        }
    }
}
