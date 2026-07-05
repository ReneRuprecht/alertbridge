package com.example.alertbridge.alerts.unit.domain.value;

import com.example.alertbridge.alerts.domain.exception.InvalidAlertSeverityException;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.AssertionsForClassTypes.assertThatThrownBy;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class AlertSeverityTest {

    @Test
    void of_shouldReturnCritical_whenInputIsCritical() {
        AlertSeverity result = AlertSeverity.of("critical");

        assertEquals(AlertSeverity.CRITICAL, result);
    }

    @Test
    void of_shouldReturnWarning_whenInputIsWarning() {
        AlertSeverity result = AlertSeverity.of("warning");

        assertEquals(AlertSeverity.WARNING, result);
    }

    @Test
    void of_shouldReturnInfo_whenInputIsInfo() {
        AlertSeverity result = AlertSeverity.of("info");

        assertEquals(AlertSeverity.INFO, result);
    }

    @Test
    void of_shouldTrimInput() {
        assertEquals(AlertSeverity.CRITICAL, AlertSeverity.of("  critical  "));
    }

    @Test
    void of_shouldThrowException_whenInputIsInvalid() {

        assertThatThrownBy(() -> AlertSeverity.of("severe"))
                .isInstanceOf(InvalidAlertSeverityException.class)
                .hasMessageContaining("Invalid Severity");
    }

    @Test
    void of_shouldThrowException_whenInputIsBlank() {

        assertThatThrownBy(() -> AlertSeverity.of(" "))
                .isInstanceOf(InvalidAlertSeverityException.class)
                .hasMessageContaining("Severity must not be null or blank");
    }

    @Test
    void of_shouldThrowException_whenInputIsNull() {

        assertThatThrownBy(() -> AlertSeverity.of(null))
                .isInstanceOf(InvalidAlertSeverityException.class)
                .hasMessageContaining("Severity must not be null or blank");
    }

    @Test
    void value_shouldReturnRawStringValue() {
        assertEquals("critical", AlertSeverity.CRITICAL.value());
        assertEquals("warning", AlertSeverity.WARNING.value());
        assertEquals("info", AlertSeverity.INFO.value());
    }


}
