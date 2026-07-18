package com.example.alertbridge.alerts.infrastructure.http;

import com.example.alertbridge.alerts.application.GetAlertHistoryByInstanceUseCase;
import com.example.alertbridge.alerts.application.GetCurrentAlertsUseCase;
import com.example.alertbridge.alerts.application.query.GetAlertHistoryByInstanceQuery;
import com.example.alertbridge.alerts.infrastructure.http.response.AlertHistoryItemResponse;
import com.example.alertbridge.alerts.infrastructure.http.response.AlertHistoryResponse;
import com.example.alertbridge.alerts.infrastructure.http.response.CurrentAlertsResponse;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/v1/alerts")
public class AlertController {

    private final GetCurrentAlertsUseCase getCurrentAlertsUseCase;
    private final GetAlertHistoryByInstanceUseCase getAlertHistoryByInstanceUseCase;

    public AlertController(GetCurrentAlertsUseCase getCurrentAlertsUseCase,
                           GetAlertHistoryByInstanceUseCase getAlertHistoryByInstanceUseCase) {
        this.getCurrentAlertsUseCase = getCurrentAlertsUseCase;
        this.getAlertHistoryByInstanceUseCase = getAlertHistoryByInstanceUseCase;
    }


    @GetMapping("/current")
    public CurrentAlertsResponse current() {

        return new CurrentAlertsResponse(getCurrentAlertsUseCase
                .getCurrentAlerts()
                .stream()
                .map(AlertHttpMapper::toAlertHistoryItemResponse)
                .toList());
    }

    @GetMapping("/history")
    public AlertHistoryResponse getHistory(@RequestParam String instance) {

        GetAlertHistoryByInstanceQuery getAlertHistoryByInstanceQuery = AlertHttpMapper.toGetAlertHistoryByInstanceQuery(
                instance);

        List<AlertHistoryItemResponse> alerts = getAlertHistoryByInstanceUseCase
                .getHistoryByInstance(getAlertHistoryByInstanceQuery)
                .stream()
                .map(AlertHttpMapper::toAlertHistoryItemResponse)
                .toList();

        return new AlertHistoryResponse(alerts);
    }

}
