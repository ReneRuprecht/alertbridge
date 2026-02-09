import AlertHistoryTable from '../../src/components/AlertHistoryTable';

import { render, screen } from '@testing-library/react';
import { it, expect, vi } from 'vitest';
import { alertHistoryMock } from '../mocks/AlertHistory';
import { describe } from 'node:test';


vi.mock('../../src/components/AlertHistoryTableItem', () => ({
  default: () => <tr />,
}));

vi.mock('../../src/utils/formatter', () => ({
  formatAlertTime: vi.fn(),
}));

describe('AlertHistoryTable', () => {

  it('renders table header', () => {
    const alertHistory = alertHistoryMock;
    render(<AlertHistoryTable alertHistory={alertHistory} />);

    expect(screen.getByText('Alertname')).toBeInTheDocument();
    expect(screen.getByText('Job')).toBeInTheDocument();
    expect(screen.getByText('Severity')).toBeInTheDocument();
    expect(screen.getByText('Angefangen')).toBeInTheDocument();
    expect(screen.getByText('Erhalten')).toBeInTheDocument();
    expect(screen.getByText('Status')).toBeInTheDocument();

  });

});
