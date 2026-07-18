package com.example.alertbridge.alerts.infrastructure.http.response;

import java.util.List;

public record AlertHistoryResponse(List<AlertHistoryItemResponse> alerts) {
}
