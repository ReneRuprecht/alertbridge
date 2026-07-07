package com.example.alertbridge.alerts.infrastructure.cache.redis;

import com.example.alertbridge.alerts.domain.model.Alert;

import java.time.Instant;

public class AlertCurrentStateMapper {
    public static AlertCurrentState toState(Alert alert) {

        return new AlertCurrentState(
                alert.fingerprint().value(),
                alert.status().name(),
                alert.labels().alertName(),
                alert.labels().severity().name(),
                alert.labels().environment(),
                alert.labels().instance(),
                alert.labels().job(),
                alert.startsAt(),
                Instant.now()
        );
    }
}
