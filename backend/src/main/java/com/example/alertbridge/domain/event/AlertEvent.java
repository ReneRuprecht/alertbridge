package com.example.alertbridge.domain.event;

import com.example.alertbridge.domain.value.*;

public record AlertEvent(AlertFingerprint fingerprint,
                         AlertLabels labels,
                         AlertStatus status,
                         AlertStartsAt startsAt,
                         AlertReceivedAt receivedAt
) {
}
