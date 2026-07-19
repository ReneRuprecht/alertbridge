import { render, screen } from '@testing-library/react';
import { describe, it, expect } from 'vitest';
import { AlertList } from './AlertList';

describe('AlertList', () => {
  it('should render all alerts', () => {
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
});
