package com.example.alertbridge.alerts.infrastructure.http.response;

import java.util.List;

public record CurrentAlertsResponse(List<CurrentAlertResponse> alerts) {
}
