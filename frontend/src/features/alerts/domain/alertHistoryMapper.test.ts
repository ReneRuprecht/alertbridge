import { describe, it, expect } from 'vitest';
import { AlertHistoryMapper } from './alertHistoryMapper';

describe('AlertHistoryMapper', () => {
  it('should map alert history response to domain model', () => {
    const response = {
      fingerprint: 'fp1',
      status: 'resolved',
      alert_name: 'InstanceDown',
      severity: 'critical',
      environment: 'prod',
      instance: 'node-01:9100',
      job: 'node_exporter',
      starts_at: '2026-07-19T05:00:00Z',
      received_at: '2026-07-19T06:11:00Z',
    };

    const result = AlertHistoryMapper.toAlertHistory(response);

    expect(result).toEqual({
      fingerprint: 'fp1',
      status: 'resolved',
      alertName: 'InstanceDown',
      severity: 'critical',
      environment: 'prod',
      instance: 'node-01:9100',
      job: 'node_exporter',
      startsAt: new Date('2026-07-19T05:00:00Z'),
      receivedAt: new Date('2026-07-19T06:11:00Z'),
    });
  });
});
