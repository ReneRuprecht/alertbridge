package com.example.alertbridge.api.alerthistory.dto;

import com.example.alertbridge.api.alerthistory.dto.value.AlertHistoryInfoDto;

import java.util.List;

public record AlertHistoryViewDto(
        String fingerprint,
        String instance,
        List<AlertHistoryInfoDto> events) {
}
