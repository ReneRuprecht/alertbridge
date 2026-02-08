package com.example.alertbridge.api.alerthistory.dto.value;

import java.time.Instant;

public record AlertHistoryInfoDto(
        String alertName,
        String job,
        String severity,
        String status,
        Instant startsAt) {
}
