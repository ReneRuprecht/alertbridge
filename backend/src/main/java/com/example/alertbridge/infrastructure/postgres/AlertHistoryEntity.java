package com.example.alertbridge.infrastructure.postgres;

import com.example.alertbridge.domain.value.AlertStatus;
import jakarta.persistence.*;

import java.time.Instant;
import java.util.UUID;

@Entity
@Table(name = "alert_history", uniqueConstraints = @UniqueConstraint(columnNames = "alert_hash"))
public class AlertHistoryEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private UUID id;

    @Column(name = "alert_hash")
    private String alertHash;

    private String fingerprint;

    @Enumerated(EnumType.STRING)
    private AlertStatus status;

    @Column(name = "starts_at")
    private Instant startsAt;

    @Column(name = "alert_name")
    private String alertName;
    private String instance;
    private String severity;
    private String job;
    private String environment;

    public String getFingerprint() {
        return fingerprint;
    }

    public void setFingerprint(String fingerprint) {
        this.fingerprint = fingerprint;
    }

    public AlertStatus getStatus() {
        return status;
    }

    public void setStatus(AlertStatus status) {
        this.status = status;
    }

    public Instant getStartsAt() {
        return startsAt;
    }

    public void setStartsAt(Instant startsAt) {
        this.startsAt = startsAt;
    }

    public String getAlertName() {
        return alertName;
    }

    public void setAlertName(String alertName) {
        this.alertName = alertName;
    }

    public String getInstance() {
        return instance;
    }

    public void setInstance(String instance) {
        this.instance = instance;
    }

    public String getSeverity() {
        return severity;
    }

    public void setSeverity(String severity) {
        this.severity = severity;
    }

    public String getJob() {
        return job;
    }

    public void setJob(String job) {
        this.job = job;
    }

    public String getEnvironment() {
        return environment;
    }

    public void setEnvironment(String environment) {
        this.environment = environment;
    }

    public String getAlertHash() {
        return alertHash;
    }

    public void setAlertHash(String alertHash) {
        this.alertHash = alertHash;
    }

}
