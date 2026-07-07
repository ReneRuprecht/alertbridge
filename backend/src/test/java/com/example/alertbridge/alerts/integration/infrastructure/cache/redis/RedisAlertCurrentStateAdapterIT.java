package com.example.alertbridge.alerts.integration.infrastructure.cache.redis;

import com.example.alertbridge.alerts.config.RedisTestConfiguration;
import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.value.AlertFingerprint;
import com.example.alertbridge.alerts.domain.value.AlertLabels;
import com.example.alertbridge.alerts.domain.value.AlertSeverity;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import com.example.alertbridge.alerts.infrastructure.cache.redis.AlertCurrentState;
import com.example.alertbridge.alerts.infrastructure.cache.redis.RedisAlertCurrentStateAdapter;
import com.redis.testcontainers.RedisContainer;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.DynamicPropertyRegistry;
import org.springframework.test.context.DynamicPropertySource;
import org.springframework.test.context.junit.jupiter.SpringExtension;
import org.testcontainers.junit.jupiter.Container;
import org.testcontainers.junit.jupiter.Testcontainers;

import java.time.Instant;
import java.util.List;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@Testcontainers
@ContextConfiguration(classes = RedisTestConfiguration.class)
@ExtendWith(SpringExtension.class)
public class RedisAlertCurrentStateAdapterIT {

    @Container
    static RedisContainer redis = new RedisContainer("redis:8.2.3-alpine");

    @Autowired
    private RedisAlertCurrentStateAdapter redisAlertCurrentStateAdapter;

    @Autowired
    private RedisTemplate<String, AlertCurrentState> redisTemplate;

    @DynamicPropertySource
    static void configureProperties(DynamicPropertyRegistry registry) {
        registry.add("spring.data.redis.host", redis::getHost);
        registry.add("spring.data.redis.port", redis::getFirstMappedPort);
    }

    @Test
    void redisContainerShouldBeRunning() {
        assertThat(redis.isRunning()).isTrue();
    }

    @Test
    void shouldSaveCurrentState_whenAlertIsFiring() {

        Alert alert = testAlert("fp1", "FIRING");

        redisAlertCurrentStateAdapter.saveAll(List.of(alert));

        AlertCurrentState stored = redisTemplate.opsForValue().get("alert:fp1");

        assertThat(stored).isNotNull();
        assertThat(stored.fingerprint()).isEqualTo("fp1");
        assertThat(stored.status()).isEqualTo("FIRING");
        assertThat(stored.alertName()).isEqualTo("CPUHigh");
    }


    @Test
    void shouldDeleteCurrentState_whenAlertIsResolved() {

        redisAlertCurrentStateAdapter.saveAll(List.of(testAlert("fp1", "FIRING")));
        assertThat(redisTemplate.hasKey("alert:fp1")).isTrue();

        redisAlertCurrentStateAdapter.saveAll(List.of(testAlert("fp1", "RESOLVED")));
        assertThat(redisTemplate.hasKey("alert:fp1")).isFalse();
    }


    private Alert testAlert(String fingerprint, String status) {

        return Alert.create(
                new AlertFingerprint(fingerprint),
                AlertStatus.of(status),
                new AlertLabels("CPUHigh", AlertSeverity.CRITICAL, "prod", "i-1", "node-exporter"),
                Instant.parse("2026-01-01T00:00:00Z")
        );
    }
}
