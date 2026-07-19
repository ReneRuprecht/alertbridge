import { render, screen } from '@testing-library/react';
import { describe, it, expect } from 'vitest';
import { AlertHistoryList } from './AlertHistoryList';

describe('AlertHistoryList', () => {
  it('renders all alert history items', () => {
    const alerts = [
      {
        fingerprint: 'fp1',
        status: 'resolved',
        alertName: 'InstanceDown',
        severity: 'critical',
        environment: 'prod',
        instance: 'node-01:9100',
        job: 'node_exporter',
        startsAt: new Date('2026-07-19T10:00:00Z'),
        receivedAt: new Date('2026-07-19T10:05:00Z'),
      },
      {
        fingerprint: 'fp2',
        status: 'resolved',
        alertName: 'CPUHigh',
        severity: 'warning',
        environment: 'prod',
        instance: 'node-01:9100',
        job: 'node_exporter',
        startsAt: new Date('2026-07-18T10:00:00Z'),
        receivedAt: new Date('2026-07-18T10:05:00Z'),
      },
    ];

    render(<AlertHistoryList alerts={alerts} />);

    expect(screen.getByText('InstanceDown')).toBeInTheDocument();

    expect(screen.getByText('CPUHigh')).toBeInTheDocument();
  });
});
