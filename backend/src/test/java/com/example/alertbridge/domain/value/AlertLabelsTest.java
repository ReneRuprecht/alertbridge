package com.example.alertbridge.domain.value;

import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

class AlertLabelsTest {

    @Test
    void shouldSetDefaultsForOptionalLabels() {
        AlertLabels labels = new AlertLabels(new AlertName("CPUHigh"), new AlertEnvironment(null), new AlertInstance(null), new AlertJob(null), AlertSeverity.fromString("critical"));

        assertThat(labels.environment().value()).isEqualTo("unknown");
        assertThat(labels.instance().value()).isEqualTo("unknown");
        assertThat(labels.job().value()).isEqualTo("unknown");
    }
}
