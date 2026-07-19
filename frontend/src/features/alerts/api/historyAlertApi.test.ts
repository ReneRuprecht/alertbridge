import { describe, it, expect, vi, beforeEach } from 'vitest';
import { getAlertHistory } from './historyAlertApi';

describe('getAlertHistory', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
  });

  it('should load alert history', async () => {
    const responseBody = {
      alerts: [
        {
          fingerprint: 'fp1',
          status: 'firing',
          alert_name: 'CPUHigh',
          severity: 'critical',
          environment: 'prod',
          instance: 'node-01:9100',
          job: 'node_exporter',
          starts_at: '2026-01-01T00:00:00Z',
          last_updated_at: '2026-01-01T00:05:00Z',
        },
      ],
    };

    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue({
        ok: true,
        json: async () => responseBody,
      }),
    );

    const result = await getAlertHistory('node-01:9100');

    expect(fetch).toHaveBeenCalledWith('/api/v1/alerts/history?instance=node-01:9100');

    expect(result).toHaveLength(1);
    expect(result[0].alertName).toBe('CPUHigh');
  });

  it('should throw error when request fails', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue({
        ok: false,
      }),
    );

    await expect(getAlertHistory('node-01:9100')).rejects.toThrow('Failed to load alert history');
  });
});
