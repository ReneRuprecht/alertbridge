package com.example.alertbridge.domain.value;

import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class AlertStatusDtoTest {
    @Test
    void shouldRecognizeFiringAndResolved() {
        assertThat(AlertStatus.FIRING.isFiring()).isTrue();
        assertThat(AlertStatus.FIRING.isResolved()).isFalse();

        assertThat(AlertStatus.RESOLVED.isFiring()).isFalse();
        assertThat(AlertStatus.RESOLVED.isResolved()).isTrue();
    }
}
