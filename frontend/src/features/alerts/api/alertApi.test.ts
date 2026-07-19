import { describe, expect, it, vi, beforeEach } from 'vitest';
import { getCurrentAlerts } from './alertApi';

describe('alert api', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
  });

  it('should load current alerts', async () => {
    const response = {
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
      vi.fn(() =>
        Promise.resolve({
          ok: true,
          json: () => Promise.resolve(response),
        })
      )
    );

    const result = await getCurrentAlerts();

    expect(fetch).toHaveBeenCalledWith('/api/v1/alerts/current');

    expect(result).toHaveLength(1);
    expect(result[0].alertName).toBe('CPUHigh');
  });

  it('should throw error when request fails', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn(() =>
        Promise.resolve({
          ok: false,
        })
      )
    );

    await expect(getCurrentAlerts()).rejects.toThrow(
      'Failed to load current alerts'
    );
  });
});
