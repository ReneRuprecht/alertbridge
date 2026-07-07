package com.example.alertbridge.alerts.infrastructure.cache.redis;

import com.example.alertbridge.alerts.domain.model.CurrentAlert;
import com.example.alertbridge.alerts.domain.ports.AlertCurrentStateReaderPort;
import org.springframework.data.redis.core.RedisTemplate;

import java.util.List;
import java.util.Objects;
import java.util.Set;

public class RedisAlertCurrentStateReader implements AlertCurrentStateReaderPort {

    private final RedisTemplate<String, AlertCurrentStateRedis> redisTemplate;

    public RedisAlertCurrentStateReader(RedisTemplate<String, AlertCurrentStateRedis> redisTemplate) {
        this.redisTemplate = redisTemplate;
    }

    @Override
    public List<CurrentAlert> findCurrentAlerts() {
        Set<String> keys = redisTemplate.keys("alert:*");

        if (keys == null || keys.isEmpty()) {
            return List.of();
        }

        return keys
                .stream()
                .map(key -> redisTemplate.opsForValue().get(key))
                .filter(Objects::nonNull)
                .map(AlertCurrentStateMapper::toDomain)
                .toList();
    }
}
