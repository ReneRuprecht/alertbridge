package com.example.alertbridge.alerts.unit.application;

import com.example.alertbridge.alerts.application.GetAlertHistoryByInstanceUseCase;
import com.example.alertbridge.alerts.application.query.GetAlertHistoryByInstanceQuery;
import com.example.alertbridge.alerts.domain.model.AlertHistorySnapshot;
import com.example.alertbridge.alerts.domain.ports.AlertHistoryReaderPort;
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

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class GetAlertHistoryByInstanceUseCaseTest {

    @InjectMocks
    GetAlertHistoryByInstanceUseCase underTest;
    @Mock
    private AlertHistoryReaderPort alertHistoryReaderPort;

    @Test
    void shouldReturnHistoryForInstance() {
        GetAlertHistoryByInstanceQuery query = new GetAlertHistoryByInstanceQuery("backend-01");

        List<AlertHistorySnapshot> expected = List.of(new AlertHistorySnapshot(
                new AlertFingerprint("fp123"),
                AlertStatus.FIRING,
                "HighCPU",
                AlertSeverity.CRITICAL,
                "prod",
                "backend-01",
                "backend-exporter",
                Instant.parse("2026-07-08T10:00:00Z"),
                Instant.parse("2026-07-08T10:05:00Z")
        ));

        when(alertHistoryReaderPort.findByInstance("backend-01")).thenReturn(expected);

        List<AlertHistorySnapshot> result = underTest.getHistoryByInstance(query);

        assertThat(result).isEqualTo(expected);

        verify(alertHistoryReaderPort).findByInstance("backend-01");
    }

    @Test
    void shouldReturnEmptyHistoryWhenNoAlertsExist() {
        GetAlertHistoryByInstanceQuery query = new GetAlertHistoryByInstanceQuery("backend-01");

        when(alertHistoryReaderPort.findByInstance("backend-01")).thenReturn(List.of());

        List<AlertHistorySnapshot> result = underTest.getHistoryByInstance(query);

        assertThat(result).isEmpty();

        verify(alertHistoryReaderPort).findByInstance("backend-01");
    }

}
