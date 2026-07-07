package com.example.alertbridge.alerts.infrastructure.http;

import com.example.alertbridge.alerts.application.GetCurrentAlertsUseCase;
import com.example.alertbridge.alerts.infrastructure.http.response.CurrentAlertsResponse;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/alerts")
public class AlertController {

    private final GetCurrentAlertsUseCase getCurrentAlertsUseCase;

    public AlertController(GetCurrentAlertsUseCase getCurrentAlertsUseCase) {
        this.getCurrentAlertsUseCase = getCurrentAlertsUseCase;
    }


    @GetMapping("/current")
    public CurrentAlertsResponse current() {

        return new CurrentAlertsResponse(getCurrentAlertsUseCase
                .getCurrentAlerts()
                .stream()
                .map(AlertHttpMapper::toResponse)
                .toList());
    }
}
