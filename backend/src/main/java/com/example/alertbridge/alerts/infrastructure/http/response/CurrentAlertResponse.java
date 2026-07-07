package com.example.alertbridge.alerts.infrastructure.http.response;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.time.Instant;

public record CurrentAlertResponse(String fingerprint,
                                   String status,
                                   @JsonProperty("alert_name")
                                   String alertName,
                                   String severity,
                                   String environment,
                                   String instance,
                                   String job,
                                   @JsonProperty("starts_at")
                                   Instant startsAt,
                                   @JsonProperty("last_updated_at")
                                   Instant lastUpdatedAt) {
}
