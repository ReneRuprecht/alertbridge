package com.example.alertbridge.api.alerthistory.dto.value;

import java.time.Instant;

public record AlertHistoryEventDto(String fingerprint,
                                   String alertName,
                                   String environment,
                                   String instance,
                                   String job,
                                   String severity,
                                   String status,
                                   Instant startsAt) {
}
