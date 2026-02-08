package com.example.alertbridge.infrastructure.postgres;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;

@Repository
public interface JpaAlertHistoryRepository extends JpaRepository<AlertHistoryEntity, UUID> {

    List<AlertHistoryEntity> findByInstanceContaining(String instance);

    boolean existsByAlertHash(
            String alertHash
    );
}
