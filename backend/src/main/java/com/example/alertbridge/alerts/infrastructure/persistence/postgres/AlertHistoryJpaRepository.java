package com.example.alertbridge.alerts.infrastructure.persistence.postgres;

import com.example.alertbridge.alerts.infrastructure.persistence.postgres.entity.AlertHistoryEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;

import java.time.Instant;
import java.util.List;
import java.util.UUID;

public interface AlertHistoryJpaRepository extends JpaRepository<AlertHistoryEntity, UUID> {

    @Modifying
    @Query(value = """
            INSERT INTO alert_history (
                id,
                fingerprint,
                status,
                alert_name,
                severity,
                environment,
                instance,
                job,
                starts_at,
                received_at,
                event_key
            )
            VALUES (
                :id,
                :fingerprint,
                :status,
                :alertName,
                :severity,
                :environment,
                :instance,
                :job,
                :startsAt,
                :receivedAt,
                :eventKey
            )
            ON CONFLICT (event_key) DO NOTHING
            """, nativeQuery = true)
    void saveWithoutDuplicateEventKey(UUID id,
                                      String fingerprint,
                                      String status,
                                      String alertName,
                                      String severity,
                                      String environment,
                                      String instance,
                                      String job,
                                      Instant startsAt,
                                      Instant receivedAt,
                                      String eventKey);

    List<AlertHistoryEntity> findByInstance(String instance);
}
