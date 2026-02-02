package com.example.alertbridge.domain.value;

import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class AlertJobTest {

    @Test
    void shouldDefaultToUnknownIfNull() {
        AlertJob job = new AlertJob(null);
        assertThat(job.value()).isEqualTo("unknown");
    }

    @Test
    void shouldKeepValueIfValid() {
        AlertJob job = new AlertJob("node-exporter");
        assertThat(job.value()).isEqualTo("node-exporter");
    }

}
