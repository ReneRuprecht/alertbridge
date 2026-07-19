import { render, screen } from '@testing-library/react';
import { describe, it, expect } from 'vitest';
import { AlertHistoryItem } from './AlertHistoryItem';

describe('AlertHistoryItem', () => {
  it('renders alert history details', () => {
    const alert = {
      fingerprint: 'fp1',
      status: 'resolved',
      alertName: 'InstanceDown',
      severity: 'critical',
      environment: 'prod',
      instance: 'node-01:9100',
      job: 'node_exporter',
      startsAt: new Date('2026-07-19T20:57:11Z'),
      receivedAt: new Date('2026-07-19T20:58:11Z'),
    };

    render(<AlertHistoryItem alert={alert} />);

    expect(screen.getByText('InstanceDown')).toBeInTheDocument();

    expect(screen.getByText('node-01:9100')).toBeInTheDocument();

    expect(screen.getByText('resolved')).toBeInTheDocument();

    expect(screen.getByText(alert.startsAt.toLocaleString())).toBeInTheDocument();

    expect(screen.getByText(alert.receivedAt.toLocaleString())).toBeInTheDocument();
  });
});
