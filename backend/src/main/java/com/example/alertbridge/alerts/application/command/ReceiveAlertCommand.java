package com.example.alertbridge.alerts.application.command;

import java.time.Instant;

public record ReceiveAlertCommand(String fingerprint,
                                  String status,
                                  String alertName,
                                  String environment,
                                  String instance,
                                  String job,
                                  String severity,
                                  Instant startsAt) {
}
