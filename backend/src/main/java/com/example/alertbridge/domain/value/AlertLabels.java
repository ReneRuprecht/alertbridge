package com.example.alertbridge.domain.value;

public record AlertLabels(
        AlertName alertName,
        AlertEnvironment environment,
        AlertInstance instance,
        AlertJob job,
        AlertSeverity severity
) {
}
