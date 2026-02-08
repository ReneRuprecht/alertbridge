package com.example.alertbridge.infrastructure.postgres;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertHistoryRepository;
import com.example.alertbridge.domain.value.AlertInstance;
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
    public List<AlertEvent> findByAlertInstance(AlertInstance instance) {
        return this.repository
                .findByInstanceContaining(instance.value())
                .stream()
                .map(AlertHistoryMapper::toDomain)
                .toList();
    }

    @Override
    public boolean existsByAlertHash(AlertState state) {
        String hash = AlertHistoryMapper.computeHash(state);
        return this.repository.existsByAlertHash(hash);
    }

}
