package com.example.alertbridge.alerts.unit.domain.value;

import com.example.alertbridge.alerts.domain.exception.InvalidAlertStatusException;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;
import static org.assertj.core.api.AssertionsForClassTypes.assertThatThrownBy;

public class AlertStatusTest {

    @Test
    void of_shouldThrowException_whenValueIsNull() {
        assertThatThrownBy(() -> AlertStatus.of(null))
                .isInstanceOf(InvalidAlertStatusException.class)
                .hasMessage("Status must not be null or blank");
    }

    @Test
    void of_shouldThrowException_whenValueIsBlank() {
        assertThatThrownBy(() -> AlertStatus.of("   "))
                .isInstanceOf(InvalidAlertStatusException.class)
                .hasMessage("Status must not be null or blank");
    }

    @Test
    void of_shouldThrowException_whenValueIsInvalid() {
        assertThatThrownBy(() -> AlertStatus.of("pending"))
                .isInstanceOf(InvalidAlertStatusException.class)
                .hasMessage("Invalid Status: pending");
    }

    @Test
    void of_shouldReturnFiring_whenValueIsValid() {
        assertThat(AlertStatus.of("firing")).isEqualTo(AlertStatus.FIRING);
    }

    @Test
    void of_shouldReturnResolved_whenValueIsValid() {
        assertThat(AlertStatus.of("resolved")).isEqualTo(AlertStatus.RESOLVED);
    }


}
