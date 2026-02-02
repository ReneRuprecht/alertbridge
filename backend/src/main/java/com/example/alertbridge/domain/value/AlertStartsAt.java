package com.example.alertbridge.domain.value;

import java.time.Instant;

public record AlertStartsAt(Instant value) {

    public AlertStartsAt {
        if (value == null) {
            throw new IllegalArgumentException("AlertStartsAt must not be null or blank");
        }
    }
}
