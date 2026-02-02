package com.example.alertbridge.domain.value;

import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class AlertSeverityDtoTest {
    @Test
    void shouldRecognizeSeverityValues() {
        assertThat(AlertSeverity.CRITICAL.name()).isEqualTo("CRITICAL");
        assertThat(AlertSeverity.WARNING.name()).isEqualTo("WARNING");
        assertThat(AlertSeverity.INFO.name()).isEqualTo("INFO");
    }
}
