package com.example.alertbridge.alerts.infrastructure.http;

import com.example.alertbridge.alerts.application.ReceiveAlertsUseCase;
import com.example.alertbridge.alerts.application.command.ReceiveAlertsCommand;
import com.example.alertbridge.alerts.infrastructure.http.request.PrometheusPayloadRequest;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/alerts")
public class PrometheusWebhookController {

    private final ReceiveAlertsUseCase receiveAlertsUseCase;

    public PrometheusWebhookController(ReceiveAlertsUseCase receiveAlertsUseCase) {
        this.receiveAlertsUseCase = receiveAlertsUseCase;
    }

    @PostMapping("/webhook")
    public ResponseEntity<Void> handleWebhook(@RequestBody PrometheusPayloadRequest request) {

        ReceiveAlertsCommand command = AlertHttpMapper.toCommand(request);
        receiveAlertsUseCase.receive(command);

        return ResponseEntity.ok().build();
    }

}
