package com.example.alertbridge.domain.value;

import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class AlertInstanceTest {

    @Test
    void shouldDefaultToUnknownIfNull() {
        AlertInstance instance = new AlertInstance(null);
        assertThat(instance.value()).isEqualTo("unknown");
    }

    @Test
    void shouldDefaultToUnknownIfBlank() {
        AlertInstance instance = new AlertInstance("   ");
        assertThat(instance.value()).isEqualTo("unknown");
    }

    @Test
    void shouldKeepValueIfValid() {
        AlertInstance instance = new AlertInstance("server-01");
        assertThat(instance.value()).isEqualTo("server-01");
    }

}
