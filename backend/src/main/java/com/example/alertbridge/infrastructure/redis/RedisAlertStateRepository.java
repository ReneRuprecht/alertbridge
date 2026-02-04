package com.example.alertbridge.infrastructure.redis;

import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertStateRepository;
import com.example.alertbridge.domain.value.AlertFingerprint;
import com.example.alertbridge.infrastructure.redis.mapper.AlertStateRedisMapper;
import org.springframework.context.annotation.Profile;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.stream.StreamSupport;

@Repository
@Profile("dev")
public class RedisAlertStateRepository implements AlertStateRepository {
    private final AlertStateRedisCrudRepository repository;

    public RedisAlertStateRepository(AlertStateRedisCrudRepository repository) {
        this.repository = repository;
    }


    @Override
    public void save(AlertState state) {
        AlertStateRedis redisState = AlertStateRedisMapper.toRedis(state);
        this.repository.save(redisState);
    }

    @Override
    public List<AlertState> findAllActive() {
        return StreamSupport
                .stream(this.repository.findAll().spliterator(), false)
                .map(AlertStateRedisMapper::toDomain)
                .filter(AlertState::isActive)
                .toList();
    }

    @Override
    public Optional<AlertState> findByFingerprint(AlertFingerprint fingerprint) {
        return this.repository.findById(fingerprint.value()).map(AlertStateRedisMapper::toDomain);
    }
}
