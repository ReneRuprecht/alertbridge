package com.example.alertbridge.domain.port;

import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.value.AlertFingerprint;

import java.util.List;
import java.util.Optional;

public interface AlertStateRepository {
    void save(AlertState state);

    List<AlertState> findAllActive();

    Optional<AlertState> findByFingerprint(AlertFingerprint fingerprint);
}
