package com.example.alertbridge.domain.port;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.value.AlertFingerprint;

import java.util.List;

public interface AlertHistoryRepository {
    void save(AlertState state);

    List<AlertEvent> findByFingerprint(AlertFingerprint fingerprint);
}
