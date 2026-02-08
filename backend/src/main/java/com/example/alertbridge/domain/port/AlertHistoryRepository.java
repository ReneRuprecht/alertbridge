package com.example.alertbridge.domain.port;

import com.example.alertbridge.domain.event.AlertEvent;
import com.example.alertbridge.domain.model.AlertState;
import com.example.alertbridge.domain.value.AlertInstance;

import java.util.List;

public interface AlertHistoryRepository {
    void save(AlertState state);

    List<AlertEvent> findByAlertInstance(AlertInstance instance);

    boolean existsByAlertHash(AlertState state);
}
