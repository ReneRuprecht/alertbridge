package com.example.alertbridge.domain.value;

import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class AlertEnvironmentTest {

    @Test
    void shouldDefaultToUnknownIfNull() {
        AlertEnvironment env = new AlertEnvironment(null);
        assertThat(env.value()).isEqualTo("unknown");
    }

    @Test
    void shouldKeepValueIfNotBlank() {
        AlertEnvironment env = new AlertEnvironment("dev");
        assertThat(env.value()).isEqualTo("dev");
    }
}
