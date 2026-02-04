package com.example.alertbridge.infrastructure.redis;

import com.example.alertbridge.domain.model.AlertState;
import org.springframework.data.annotation.Id;
import org.springframework.data.redis.core.RedisHash;

import java.io.Serializable;

@RedisHash(value = "alert_state", timeToLive = 172800)
public class AlertStateRedis implements Serializable {

    @Id
    private String fingerprint;
    private AlertState state;

    public AlertStateRedis() {

    }

    public AlertStateRedis(String fingerprint, AlertState state) {
        this.fingerprint = fingerprint;
        this.state = state;
    }

    public String getFingerprint() {
        return this.fingerprint;
    }

    public AlertState getState() {
        return this.state;
    }

}
