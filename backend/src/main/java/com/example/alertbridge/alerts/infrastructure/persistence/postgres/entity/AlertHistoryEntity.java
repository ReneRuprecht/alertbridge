package com.example.alertbridge.alerts.infrastructure.persistence.postgres.entity;

import jakarta.persistence.*;

import java.time.Instant;
import java.util.UUID;

@Entity
@Table(name = "alert_history", uniqueConstraints = {@UniqueConstraint(columnNames = "event_key")})
public class AlertHistoryEntity {
    @Id
    private UUID id;

    @Column(nullable = false)
    private String fingerprint;

    @Column(nullable = false)
    private String status;

    @Column(nullable = false)
    private String alertName;

    @Column(nullable = false)
    private String severity;

    @Column(nullable = false)
    private String environment;

    @Column(nullable = false)
    private String instance;

    @Column(nullable = false)
    private String job;

    @Column(nullable = false)
    private Instant startsAt;

    @Column(nullable = false)
    private Instant receivedAt;

    @Column(nullable = false, unique = true)
    private String eventKey;

    protected AlertHistoryEntity() {
    }

    public AlertHistoryEntity(UUID id,
                              String fingerprint,
                              String status,
                              String alertName,
                              String severity,
                              String environment,
                              String instance,
                              String job,
                              Instant startsAt,
                              Instant receivedAt,
                              String eventKey) {
        this.id = id;
        this.fingerprint = fingerprint;
        this.status = status;
        this.alertName = alertName;
        this.severity = severity;
        this.environment = environment;
        this.instance = instance;
        this.job = job;
        this.startsAt = startsAt;
        this.receivedAt = receivedAt;
        this.eventKey = eventKey;
    }

    public UUID getId() {
        return id;
    }

    public String getFingerprint() {
        return fingerprint;
    }

    public String getStatus() {
        return status;
    }

    public String getAlertName() {
        return alertName;
    }

    public String getSeverity() {
        return severity;
    }

    public String getEnvironment() {
        return environment;
    }

    public String getInstance() {
        return instance;
    }

    public String getJob() {
        return job;
    }

    public Instant getStartsAt() {
        return startsAt;
    }

    public Instant getReceivedAt() {
        return receivedAt;
    }

    public String getEventKey() {
        return eventKey;
    }
}
