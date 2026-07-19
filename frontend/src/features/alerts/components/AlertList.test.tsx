import { fireEvent, render, screen } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import { AlertList } from './AlertList';

describe('AlertList', () => {
  it('should render all alerts', () => {
    const onSelect = vi.fn();
    render(
      <AlertList
        alerts={[
          {
            fingerprint: 'fp-1',
            status: 'FIRING',
            alertName: 'InstanceDown',
            severity: 'CRITICAL',
            environment: 'prod',
            instance: 'server-01',
            job: 'node_exporter',
            startsAt: '2026-01-01T00:00:00Z',
            lastUpdatedAt: '2026-01-01T00:05:00Z',
          },
          {
            fingerprint: 'fp-2',
            status: 'FIRING',
            alertName: 'CPUHigh',
            severity: 'WARNING',
            environment: 'prod',
            instance: 'server-02',
            job: 'job',
            startsAt: '2026-01-01T00:00:00Z',
            lastUpdatedAt: '2026-01-01T00:05:00Z',
          },
        ]}
        onSelect={onSelect}
      />,
    );

    expect(screen.getByText('InstanceDown')).toBeInTheDocument();
    expect(screen.getByText('server-01')).toBeInTheDocument();
    expect(screen.getByText('node_exporter')).toBeInTheDocument();
    expect(screen.getByText('CRITICAL')).toBeInTheDocument();

    expect(screen.getByText('CPUHigh')).toBeInTheDocument();
    expect(screen.getByText('server-02')).toBeInTheDocument();
    expect(screen.getByText('WARNING')).toBeInTheDocument();
  });
  it('calls onSelect with the selected alert', () => {
    const onSelect = vi.fn();

    const alerts = [
      {
        fingerprint: 'fp-1',
        status: 'FIRING',
        alertName: 'InstanceDown',
        severity: 'CRITICAL',
        environment: 'prod',
        instance: 'node-01',
        job: 'node_exporter',
        startsAt: '2026-01-01T00:00:00Z',
        lastUpdatedAt: '2026-01-01T00:05:00Z',
      },
    ];

    render(<AlertList alerts={alerts} onSelect={onSelect} />);

    fireEvent.click(screen.getByText('node-01'));

    expect(onSelect).toHaveBeenCalledWith(alerts[0]);
  });
});
