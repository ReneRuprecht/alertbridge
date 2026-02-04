package com.example.alertbridge.infrastructure.redis.mapper;

import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.infrastructure.redis.AlertStateRedis;
import org.springframework.stereotype.Component;

@Component
public class AlertStateRedisMapper {
    public static AlertStateRedis toRedis(AlertState state) {
        return new AlertStateRedis(state.fingerprint().value(), state);
    }

    public static AlertState toDomain(AlertStateRedis redis) {
        return redis.getState();
    }
}
