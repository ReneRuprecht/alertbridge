package com.example.alertbridge.alerts.infrastructure.cache.redis;

import com.example.alertbridge.alerts.domain.model.Alert;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateBatchWriterPort;
import com.example.alertbridge.alerts.domain.value.AlertStatus;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Repository;

import java.time.Duration;
import java.util.List;

@Repository
public class RedisAlertCurrentStateCache implements AlertCurrentStateBatchWriterPort {
    private final RedisTemplate<String, AlertCurrentState> redisTemplate;

    public RedisAlertCurrentStateCache(RedisTemplate<String, AlertCurrentState> redisTemplate) {
        this.redisTemplate = redisTemplate;
    }

    @Override
    public void saveAll(List<Alert> alerts) {

        alerts.forEach(alert -> {
            String key = String.format("alert:%s", alert.fingerprint().value());

            if (alert.status() == AlertStatus.RESOLVED) {
                redisTemplate.delete(key);
                return;
            }

            redisTemplate
                    .opsForValue()
                    .set(key, AlertCurrentStateMapper.toState(alert), Duration.ofHours(24));
        });
    }
}
