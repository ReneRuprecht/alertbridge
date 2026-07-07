package com.example.alertbridge.alerts.infrastructure.cache.redis;

import java.time.Instant;

public record AlertCurrentStateRedis(String fingerprint,
                                     String status,
                                     String alertName,
                                     String severity,
                                     String environment,
                                     String instance,
                                     String job,
                                     Instant startsAt,
                                     Instant lastUpdatedAt) {
}
