package com.example.alertbridge.alerts.infrastructure.http.response;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.time.Instant;

public record AlertHistoryItemResponse(String fingerprint,
                                       String status,
                                       @JsonProperty("alert_name")
                                       String alertName,
                                       String severity,
                                       String environment,
                                       String instance,
                                       String job,
                                       @JsonProperty("starts_at")
                                       Instant startsAt,
                                       @JsonProperty("received_at")
                                       Instant receivedAt) {
}
