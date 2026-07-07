package com.example.alertbridge.alerts.integration.infrastructure.http;

import com.example.alertbridge.alerts.application.GetCurrentAlertsUseCase;
import com.example.alertbridge.alerts.domain.model.CurrentAlert;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import com.example.alertbridge.alerts.infrastructure.http.AlertController;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.webmvc.test.autoconfigure.WebMvcTest;
import org.springframework.http.MediaType;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.test.web.servlet.MockMvc;

import java.time.Instant;
import java.util.List;

import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@WebMvcTest(AlertController.class)
class AlertControllerTest {

    @MockitoBean
    GetCurrentAlertsUseCase getCurrentAlertsUseCase;

    @Autowired
    MockMvc mockMvc;

    @Test
    void shouldReturnCurrentAlerts() throws Exception {

        CurrentAlert alert = new CurrentAlert(
                new AlertFingerprint("fp1"),
                AlertStatus.FIRING,
                "CPUHigh",
                AlertSeverity.CRITICAL,
                "prod",
                "server-1",
                "job-a",
                Instant.parse("2026-01-01T00:00:00Z"),
                Instant.parse("2026-01-01T00:01:00Z")
        );

        when(getCurrentAlertsUseCase.getCurrentAlerts()).thenReturn(List.of(alert));


        mockMvc
                .perform(get("/api/v1/alerts/current"))
                .andExpect(status().isOk())
                .andExpect(content().contentType(MediaType.APPLICATION_JSON))
                .andExpect(jsonPath("$.alerts").isArray())
                .andExpect(jsonPath("$.alerts.length()").value(1))
                .andExpect(jsonPath("$.alerts[0].fingerprint").value("fp1"))
                .andExpect(jsonPath("$.alerts[0].status").value("FIRING"))
                .andExpect(jsonPath("$.alerts[0].alert_name").value("CPUHigh"))
                .andExpect(jsonPath("$.alerts[0].severity").value("CRITICAL"))
                .andExpect(jsonPath("$.alerts[0].environment").value("prod"))
                .andExpect(jsonPath("$.alerts[0].instance").value("server-1"))
                .andExpect(jsonPath("$.alerts[0].job").value("job-a"))
                .andExpect(jsonPath("$.alerts[0].starts_at").value("2026-01-01T00:00:00Z"))
                .andExpect(jsonPath("$.alerts[0].last_updated_at").value("2026-01-01T00:01:00Z"));

        verify(getCurrentAlertsUseCase).getCurrentAlerts();
    }


    @Test
    void shouldReturnEmptyAlerts_whenNoCurrentAlertsExist() throws Exception {

        when(getCurrentAlertsUseCase.getCurrentAlerts()).thenReturn(List.of());


        mockMvc
                .perform(get("/api/v1/alerts/current"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.alerts").isArray())
                .andExpect(jsonPath("$.alerts").isEmpty());
    }
}
