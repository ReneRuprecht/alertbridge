import { renderHook, waitFor } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import { useAlertHistory } from './useAlertHistory';
import * as api from '../api/historyAlertApi';

vi.mock('../api/historyAlertApi');

describe('useAlertHistory', () => {
  it('should load alert history', async () => {
    vi.mocked(api.getAlertHistory).mockResolvedValue([
      {
        fingerprint: 'fp1',
        alertName: 'InstanceDown',
        instance: 'node-01:9100',
        severity: 'CRITICAL',
        environment: 'prod',
        job: 'node_exporter',
        status: 'FIRING',
        startsAt: '2026-19-07T20:57:11.872Z',
        receivedAt: '2026-19-07T20:58:11.872Z',
      },
    ] as any);

    const { result } = renderHook(() => useAlertHistory('node-01:9100'));

    expect(result.current.loading).toBe(true);

    await waitFor(() => {
      expect(result.current.loading).toBe(false);
    });

    expect(result.current.history).toHaveLength(1);
  });
});
