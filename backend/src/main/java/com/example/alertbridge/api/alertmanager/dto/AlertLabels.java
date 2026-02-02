package com.example.alertbridge.api.alertmanager.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class AlertLabels {
    @JsonProperty("alertname")
    public String alertName;
    public String environment;
    public String instance;
    public String job;
    public AlertSeverity severity;
}
