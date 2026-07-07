package com.example.alertbridge.alerts.integration.infrastructure.cache.redis;

import com.example.alertbridge.alerts.config.RedisTestConfiguration;
import com.example.alertbridge.alerts.domain.model.CurrentAlert;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import com.example.alertbridge.alerts.infrastructure.cache.redis.AlertCurrentStateRedis;
import com.example.alertbridge.alerts.infrastructure.cache.redis.RedisAlertCurrentStateReader;
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

import static org.assertj.core.api.Assertions.assertThat;


@Testcontainers
@ContextConfiguration(classes = RedisTestConfiguration.class)
@ExtendWith(SpringExtension.class)
public class RedisAlertCurrentStateReaderIT {

    @Container
    static RedisContainer redis = new RedisContainer("redis:8.2.3-alpine");

    @Autowired
    private RedisAlertCurrentStateReader redisAlertCurrentStateReader;

    @Autowired
    private RedisTemplate<String, AlertCurrentStateRedis> redisTemplate;

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
    void shouldReadCurrentAlertsFromRedis() {

        AlertCurrentStateRedis state = new AlertCurrentStateRedis(
                "fp1",
                "firing",
                "CPUHigh",
                "critical",
                "prod",
                "instance-1",
                "job-a",
                Instant.parse("2026-01-01T00:00:00Z"),
                Instant.parse("2026-01-01T00:01:00Z")
        );


        redisTemplate.opsForValue().set("alert:fp1", state);

        List<CurrentAlert> result = redisAlertCurrentStateReader.findCurrentAlerts();

        assertThat(result).hasSize(1);

        CurrentAlert alert = result.getFirst();

        assertThat(alert.fingerprint().value()).isEqualTo("fp1");
        assertThat(alert.status()).isEqualTo(AlertStatus.FIRING);
    }

}
