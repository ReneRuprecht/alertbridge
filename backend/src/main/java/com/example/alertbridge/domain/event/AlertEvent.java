package com.example.alertbridge.domain.event;

import com.example.alertbridge.domain.value.AlertFingerprint;
import com.example.alertbridge.domain.value.AlertLabels;
import com.example.alertbridge.domain.value.AlertStartsAt;
import com.example.alertbridge.domain.value.AlertStatus;

public record AlertEvent(
        AlertFingerprint fingerprint, AlertLabels labels, AlertStatus status, AlertStartsAt startsAt
) {
}
