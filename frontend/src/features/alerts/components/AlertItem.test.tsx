import { render, screen } from '@testing-library/react';
import { describe, it, expect } from 'vitest';
import { AlertItem } from './AlertItem';

describe('AlertItem', () => {
  it('should render alert details', () => {
    render(
      <AlertItem
        alert={{
          fingerprint: 'fp-1',
          status: 'FIRING',
          alertName: 'InstanceDown',
          severity: 'CRITICAL',
          environment: 'prod',
          instance: 'server-01',
          job: 'node_exporter',
          startsAt: '2026-01-01T00:00:00Z',
          lastUpdatedAt: '2026-01-01T00:05:00Z',
        }}
      />,
    );

    expect(screen.getByText('InstanceDown')).toBeInTheDocument();

    expect(screen.getByText('server-01')).toBeInTheDocument();

    expect(screen.getByText('CRITICAL')).toBeInTheDocument();
  });
});
