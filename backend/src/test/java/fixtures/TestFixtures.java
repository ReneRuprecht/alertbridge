package fixtures;

import com.example.alertbridge.api.alertmanager.dto.AlertDto;
import com.example.alertbridge.api.alertmanager.dto.AlertLabelsDto;
import com.example.alertbridge.api.alertmanager.dto.AlertSeverityDto;
import com.example.alertbridge.api.alertmanager.dto.AlertStatusDto;
import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.value.*;

import java.time.Instant;

public class TestFixtures {

    public static AlertFingerprint fingerprint(String id) {
        return new AlertFingerprint(id);
    }

    public static AlertLabels defaultLabels() {
        return new AlertLabels(
                new AlertName("CPUHigh"),
                new AlertEnvironment("prod"),
                new AlertInstance("server-01"),
                new AlertJob("node-exporter"),
                AlertSeverity.CRITICAL
        );
    }

    public static AlertLabels defaultLabels(String instance) {
        return new AlertLabels(
                new AlertName("CPUHigh"),
                new AlertEnvironment("prod"),
                new AlertInstance(instance),
                new AlertJob("node-exporter"),
                AlertSeverity.CRITICAL
        );
    }


    public static AlertStartsAt alertStartsAt(Instant instant) {
        return new AlertStartsAt(instant);
    }

    public static AlertReceivedAt alertReceivedAt(Instant instant) {
        return new AlertReceivedAt(instant);
    }

    public static AlertEvent firingEvent(String fingerprint) {
        return new AlertEvent(
                fingerprint(fingerprint),
                defaultLabels(),
                AlertStatus.FIRING,
                alertStartsAt(Instant.parse("2026-02-02T10:00:00Z")),
                alertReceivedAt(Instant.parse("2026-02-02T11:00:00Z"))
        );
    }

    public static AlertEvent resolvedEvent(String fingerprint) {
        return new AlertEvent(
                fingerprint(fingerprint),
                defaultLabels(),
                AlertStatus.RESOLVED,
                alertStartsAt(Instant.parse("2026-02-02T10:00:00Z")),
                alertReceivedAt(Instant.parse("2026-02-02T11:00:00Z"))
        );
    }

    public static AlertEvent firingEvent(String fingerprint, String instance) {
        return new AlertEvent(
                fingerprint(fingerprint),
                defaultLabels(instance),
                AlertStatus.FIRING,
                alertStartsAt(Instant.parse("2026-02-02T10:00:00Z")),
                alertReceivedAt(Instant.parse("2026-02-02T11:00:00Z"))
        );
    }

    public static AlertEvent resolvedEvent(String fingerprint, String instance) {
        return new AlertEvent(
                fingerprint(fingerprint),
                defaultLabels(instance),
                AlertStatus.RESOLVED,
                alertStartsAt(Instant.parse("2026-02-02T10:00:00Z")),
                alertReceivedAt(Instant.parse("2026-02-02T11:00:00Z"))
        );
    }


    public static AlertDto alertDto(String fingerprint, AlertStatusDto status, String environment) {
        return new AlertDto(
                fingerprint,
                status,
                new AlertLabelsDto(
                        "CPUHigh",
                        environment,
                        "server-01",
                        "node-exporter",
                        AlertSeverityDto.CRITICAL
                ),
                Instant.parse("2026-02-02T10:00:00Z").toString()
        );
    }

    public static AlertState firingAlert(String fingerprint) {
        return AlertState.fromEvent(firingEvent(fingerprint));
    }

    public static AlertState resolvedAlert(String fingerprint) {
        return AlertState.fromEvent(resolvedEvent(fingerprint));
    }

}
