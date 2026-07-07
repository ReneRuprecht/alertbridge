package com.example.alertbridge.alerts.domain.value;

import com.example.alertbridge.alerts.domain.exception.InvalidAlertNameException;

public record AlertLabels(String alertName,
                          AlertSeverity severity,
                          String environment,
                          String instance,
                          String job) {

    public AlertLabels {
        if (alertName == null || alertName.isBlank())
            throw new InvalidAlertNameException("AlertName must not be null or blank");

        environment = (environment == null || environment.isBlank()) ? "unknown" : environment.trim();
        instance = (instance == null || instance.isBlank()) ? "unknown" : instance.trim();
        job = (job == null || job.isBlank()) ? "unknown" : job.trim();
    }


}
