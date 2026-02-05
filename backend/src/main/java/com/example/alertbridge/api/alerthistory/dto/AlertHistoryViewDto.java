package com.example.alertbridge.api.alerthistory.dto;

import com.example.alertbridge.api.alerthistory.dto.value.AlertHistoryEventDto;

import java.util.List;

public record AlertHistoryViewDto(List<AlertHistoryEventDto> events) {
}
