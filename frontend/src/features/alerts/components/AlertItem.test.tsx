import { fireEvent, render, screen } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import { AlertItem } from './AlertItem';
import type { Alert } from '../domain/alert';

describe('AlertItem', () => {
  it('should render alert details', () => {
    const onClick = vi.fn();
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
        onClick={onClick}
      />,
    );

    expect(screen.getByText('InstanceDown')).toBeInTheDocument();

    expect(screen.getByText('server-01')).toBeInTheDocument();

    expect(screen.getByText('CRITICAL')).toBeInTheDocument();
  });
  it('calls onClick when clicked', async () => {
    const onClick = vi.fn();
    const mockAlert: Alert = {
      fingerprint: 'fp-1',
      status: 'FIRING',
      alertName: 'InstanceDown',
      severity: 'CRITICAL',
      environment: 'prod',
      instance: 'node-01',
      job: 'node_exporter',
      startsAt: '2026-01-01T00:00:00Z',
      lastUpdatedAt: '2026-01-01T00:05:00Z',
    };

    render(<AlertItem alert={mockAlert} onClick={onClick} />);

    fireEvent.click(screen.getByText(mockAlert.instance));

    expect(onClick).toHaveBeenCalled();
  });
});
