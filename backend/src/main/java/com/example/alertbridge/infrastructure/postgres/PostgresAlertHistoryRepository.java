package com.example.alertbridge.infrastructure.postgres;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertHistoryRepository;
import com.example.alertbridge.domain.value.AlertFingerprint;
import com.example.alertbridge.infrastructure.postgres.mapper.AlertHistoryMapper;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class PostgresAlertHistoryRepository implements AlertHistoryRepository {

    private final JpaAlertHistoryRepository repository;

    public PostgresAlertHistoryRepository(JpaAlertHistoryRepository repository) {
        this.repository = repository;
    }

    @Override
    public void save(AlertState state) {
        this.repository.save(AlertHistoryMapper.toEntity(state));
    }

    @Override
    public List<AlertEvent> findByFingerprint(AlertFingerprint fingerprint) {
        return this.repository
                .findByFingerprint(fingerprint.value())
                .stream()
                .map(AlertHistoryMapper::toDomain)
                .toList();
    }

}
