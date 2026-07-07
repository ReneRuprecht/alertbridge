package com.example.alertbridge.alerts.infrastructure.http.request;

import java.util.List;

public record PrometheusPayloadRequest(List<PrometheusAlert> alerts) {
    public PrometheusPayloadRequest {
        if (alerts == null) {
            alerts = List.of();
        }
    }
}
