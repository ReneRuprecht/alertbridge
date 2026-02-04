package com.example.alertbridge.infrastructure.repository;

import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.port.AlertStateRepository;
import com.example.alertbridge.domain.value.AlertFingerprint;
import org.springframework.context.annotation.Profile;
import org.springframework.stereotype.Repository;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;

@Repository
@Profile("test")
public class InMemoryAlertStateRepository implements AlertStateRepository {
    private final Map<AlertFingerprint, AlertState> store = new HashMap<>();

    @Override
    public void save(AlertState state) {
        store.put(state.fingerprint(), state);
    }

    @Override
    public List<AlertState> findAllActive() {
        return store.values().stream().filter(AlertState::isActive).toList();
    }

    @Override
    public Optional<AlertState> findByFingerprint(AlertFingerprint fingerprint) {
        return Optional.ofNullable(store.get(fingerprint));
    }
}
