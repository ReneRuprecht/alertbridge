package com.example.alertbridge.alerts.integration.infrastructure.http;

import com.example.alertbridge.alerts.application.ReceiveAlertsUseCase;
import com.example.alertbridge.alerts.application.command.ReceiveAlertCommand;
import com.example.alertbridge.alerts.application.command.ReceiveAlertsCommand;
import com.example.alertbridge.alerts.infrastructure.http.PrometheusWebhookController;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.webmvc.test.autoconfigure.WebMvcTest;
import org.springframework.http.MediaType;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.test.web.servlet.MockMvc;

import static org.assertj.core.api.AssertionsForInterfaceTypes.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.verify;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(PrometheusWebhookController.class)
public class PrometheusWebhookControllerTest {

    @MockitoBean
    private ReceiveAlertsUseCase receiveAlertsUseCase;

    @Autowired
    private MockMvc mockMvc;

    private static String validJson() {
        return """
                {
                  "alerts": [
                    {
                      "fingerprint": "fp1",
                      "status": "firing",
                      "labels": {
                        "alertname": "CPUHigh",
                        "environment": "prod",
                        "instance": "i-1",
                        "job": "job-a",
                        "severity": "critical"
                      },
                      "startsAt": "2026-01-01T00:00:00Z"
                    }
                  ]
                }
                """;
    }

    @Test
    void shouldForwardRequestToUseCase() throws Exception {


        mockMvc
                .perform(post("/api/v1/alerts/webhook")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(validJson()))
                .andExpect(status().isOk());

        verify(receiveAlertsUseCase).receive(any(ReceiveAlertsCommand.class));
    }

    @Test
    void shouldMapAndPassCorrectCommand() throws Exception {

        ArgumentCaptor<ReceiveAlertsCommand> captor = ArgumentCaptor.forClass(ReceiveAlertsCommand.class);

        mockMvc
                .perform(post("/api/v1/alerts/webhook")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(validJson()))
                .andExpect(status().isOk());

        verify(receiveAlertsUseCase).receive(captor.capture());

        assertThat(captor.getValue().alerts())
                .hasSize(1)
                .extracting(ReceiveAlertCommand::fingerprint)
                .containsExactly("fp1");
    }

    @Test
    void shouldStillReturnOk_whenAlertsAreEmpty() throws Exception {

        String json = """
                { "alerts": [] }
                """;

        mockMvc
                .perform(post("/api/v1/alerts/webhook")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(json))
                .andExpect(status().isOk());

        verify(receiveAlertsUseCase).receive(any());
    }

}
