package com.example.alertbridge.alerts.domain.ports;

import com.example.alertbridge.alerts.domain.model.CurrentAlert;

import java.util.List;

public interface AlertCurrentStateReaderPort {

    List<CurrentAlert> findCurrentAlerts();
}
