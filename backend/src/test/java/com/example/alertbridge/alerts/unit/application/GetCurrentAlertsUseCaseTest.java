package com.example.alertbridge.alerts.unit.application;

import com.example.alertbridge.alerts.application.GetCurrentAlertsUseCase;
import com.example.alertbridge.alerts.domain.model.CurrentAlert;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateReaderPort;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.time.Instant;
import java.util.List;

import static org.assertj.core.api.AssertionsForInterfaceTypes.assertThat;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class GetCurrentAlertsUseCaseTest {

    @Mock
    private AlertCurrentStateReaderPort alertCurrentStateReaderPort;

    @InjectMocks
    private GetCurrentAlertsUseCase underTest;


    @Test
    void shouldReturnCurrentAlerts() {

        List<CurrentAlert> currentAlerts = List.of(
                new CurrentAlert(
                        new AlertFingerprint("fp1"),
                        AlertStatus.FIRING,
                        "CPUHigh",
                        AlertSeverity.CRITICAL,
                        "prod",
                        "instance-1",
                        "job-a",
                        Instant.parse("2026-01-01T00:00:00Z"),
                        Instant.parse("2026-01-01T00:01:00Z")
                ), new CurrentAlert(
                        new AlertFingerprint("fp2"),
                        AlertStatus.RESOLVED,
                        "MemoryHigh",
                        AlertSeverity.WARNING,
                        "prod",
                        "instance-2",
                        "job-b",
                        Instant.parse("2026-01-01T00:00:00Z"),
                        Instant.parse("2026-01-01T00:02:00Z")
                )
        );

        when(alertCurrentStateReaderPort.findCurrentAlerts()).thenReturn(currentAlerts);

        List<CurrentAlert> result = underTest.getCurrentAlerts();

        verify(alertCurrentStateReaderPort).findCurrentAlerts();

        assertThat(result).containsExactlyElementsOf(currentAlerts);
    }

    @Test
    void shouldReturnEmptyList_WhenNoCurrentAlertsExist() {

        when(alertCurrentStateReaderPort.findCurrentAlerts()).thenReturn(List.of());

        List<CurrentAlert> result = underTest.getCurrentAlerts();

        verify(alertCurrentStateReaderPort).findCurrentAlerts();

        assertThat(result).isEmpty();
    }
}
