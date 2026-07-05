package com.example.alertbridge.alerts.unit.domain.value;

import com.example.alertbridge.alerts.domain.exception.InvalidAlertNameException;
import com.example.alertbridge.alerts.domain.value.AlertLabels;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;
import static org.assertj.core.api.AssertionsForClassTypes.assertThatThrownBy;

class AlertLabelsTest {

    @Test
    void shouldThrowException_whenAlertNameIsNull() {
        assertThatThrownBy(() -> new AlertLabels(
                null,
                AlertSeverity.CRITICAL,
                "prod",
                "i-123",
                "job-a"
        ))
                .isInstanceOf(InvalidAlertNameException.class)
                .hasMessage("AlertName must not be null or blank");
    }

    @Test
    void shouldThrowException_whenAlertNameIsBlank() {
        assertThatThrownBy(() -> new AlertLabels(
                "   ",
                AlertSeverity.CRITICAL,
                "prod",
                "i-123",
                "job-a"
        ))
                .isInstanceOf(InvalidAlertNameException.class)
                .hasMessage("AlertName must not be null or blank");
    }

    @Test
    void shouldCreateSuccessfully_whenAllValuesAreProvided() {
        AlertLabels labels = new AlertLabels(
                "CPUHigh",
                AlertSeverity.CRITICAL,
                "prod",
                "instance-1",
                "job-a"
        );

        assertThat(labels.alertName()).isEqualTo("CPUHigh");
        assertThat(labels.severity()).isEqualTo(AlertSeverity.CRITICAL);
        assertThat(labels.environment()).isEqualTo("prod");
        assertThat(labels.instance()).isEqualTo("instance-1");
        assertThat(labels.job()).isEqualTo("job-a");
    }

    @Test
    void shouldApplyDefaults_whenOptionalFieldsAreBlank() {
        AlertLabels labels = new AlertLabels("CPUHigh", AlertSeverity.WARNING, "   ", "", "    ");

        assertThat(labels.environment()).isEqualTo("unknown");
        assertThat(labels.instance()).isEqualTo("unknown");
        assertThat(labels.job()).isEqualTo("unknown");
    }

    @Test
    void shouldApplyDefaults_whenOptionalFieldsAreNull() {
        AlertLabels labels = new AlertLabels("CPUHigh", AlertSeverity.WARNING, null, null, null);

        assertThat(labels.environment()).isEqualTo("unknown");
        assertThat(labels.instance()).isEqualTo("unknown");
        assertThat(labels.job()).isEqualTo("unknown");
    }
}
