package com.example.alertbridge.api.alertmanager.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public record AlertLabelsDto(@JsonProperty("alertname") String alertName,
                             String environment,
                             String instance,
                             String job,
                             AlertSeverityDto severity) {
}
