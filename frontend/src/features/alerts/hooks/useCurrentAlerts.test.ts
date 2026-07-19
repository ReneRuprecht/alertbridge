import { renderHook, waitFor } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import { useCurrentAlerts } from './useCurrentAlerts';
import * as alertApi from '../api/alertApi';

vi.mock('../api/alertApi');

describe('useCurrentAlerts', () => {
  it('should load current alerts', async () => {
    vi.mocked(alertApi.getCurrentAlerts).mockResolvedValue([
      {
        fingerprint: 'fp1',
        alertName: 'InstanceDown',
        instance: 'instance-01',
        severity: 'CRITICAL',
        environment: 'prod',
        job: 'node_exporter',
        status: 'FIRING',
        startsAt: '2026-19-07T20:57:11.872Z',
        lastUpdatedAt: '2026-19-07T20:58:11.872Z',
      },
    ]);

    const { result } = renderHook(() => useCurrentAlerts());

    expect(result.current.loading).toBe(true);

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    expect(result.current.alerts).toHaveLength(1);
    expect(result.current.alerts[0].fingerprint).toBe('fp1');
    expect(result.current.alerts[0].alertName).toBe('InstanceDown');
  });
});
