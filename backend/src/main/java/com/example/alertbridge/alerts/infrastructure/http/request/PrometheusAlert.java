package com.example.alertbridge.alerts.infrastructure.http.request;

import java.time.Instant;
import java.util.Map;

public record PrometheusAlert(String fingerprint,
                              String status,
                              Map<String, String> labels,
                              Instant startsAt) {
    public PrometheusAlert {
        if (labels == null) labels = Map.of();
    }
}
