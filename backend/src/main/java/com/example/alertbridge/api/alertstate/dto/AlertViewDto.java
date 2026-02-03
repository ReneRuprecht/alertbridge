package com.example.alertbridge.api.alertstate.dto;

import java.time.Instant;

public record AlertViewDto(String fingerprint,
                           String alertName,
                           String environment,
                           String instance,
                           String job,
                           String severity,
                           String status,
                           Instant startsAt) {
}
