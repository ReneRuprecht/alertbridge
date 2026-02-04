package com.example.alertbridge.infrastructure.redis;

import com.example.alertbridge.domain.model.AlertState;
import com.redis.testcontainers.RedisContainer;
import fixtures.TestFixtures;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.data.redis.test.autoconfigure.DataRedisTest;
import org.springframework.context.annotation.Import;
import org.springframework.test.context.ActiveProfiles;
import org.testcontainers.junit.jupiter.Container;
import org.testcontainers.junit.jupiter.Testcontainers;

import static org.assertj.core.api.Assertions.assertThat;

@DataRedisTest
@Testcontainers
@Import(RedisAlertStateRepository.class)
@ActiveProfiles("redis-it")
public class RedisAlertStateRepositoryIT {

    @Container
    static RedisContainer redis = new RedisContainer("redis:8.2.3-alpine").withExposedPorts(6379);

    @Autowired
    RedisAlertStateRepository repository;

    @Test
    void redisContainerShouldBeRunning() {
        assertThat(redis.isRunning()).isTrue();

    }

    @Test
    void shouldPersistAndLoadAlertState() {
        AlertState state = TestFixtures.firingAlert("fp-1");

        repository.save(state);

        AlertState stateFromRedis = repository.findByFingerprint(state.fingerprint()).orElseThrow();

        assertThat(stateFromRedis.status()).isEqualTo(state.status());
        assertThat(stateFromRedis.labels()).isEqualTo(state.labels());
    }


}
