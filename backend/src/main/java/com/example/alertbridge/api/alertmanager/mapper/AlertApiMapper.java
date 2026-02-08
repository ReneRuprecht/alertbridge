package com.example.alertbridge.api.alertmanager.mapper;

import com.example.alertbridge.api.alertmanager.dto.AlertDto;
import com.example.alertbridge.api.alertmanager.dto.AlertLabelsDto;
import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.value.*;
import org.springframework.stereotype.Component;

import java.time.Instant;


@Component
public class AlertApiMapper {

    public AlertEvent toEvent(AlertDto dto) {
        AlertFingerprint fingerprint = new AlertFingerprint(dto.fingerprint());
        AlertStatus status = mapStatus(dto.status().toString());
        AlertLabels labels = toAlertLabels(dto.labels());
        AlertStartsAt startsAt = new AlertStartsAt(Instant.parse(dto.startsAt()));
        AlertReceivedAt receivedAt = new AlertReceivedAt(Instant.now());

        return new AlertEvent(fingerprint, labels, status, startsAt, receivedAt);
    }

    private AlertLabels toAlertLabels(AlertLabelsDto dto) {
        return new AlertLabels(
                new AlertName(dto.alertName()),
                new AlertEnvironment(dto.environment()),
                new AlertInstance(dto.instance()),
                new AlertJob(dto.job()),
                AlertSeverity.fromString(dto.severity().toString())
        );
    }

    private AlertStatus mapStatus(String alertStatus) {
        return switch (alertStatus.toLowerCase()) {
            case "firing" -> AlertStatus.FIRING;
            case "resolved" -> AlertStatus.RESOLVED;
            default -> AlertStatus.UNKNOWN;
        };
    }
}
