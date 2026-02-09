import { describe, expect, it, vi } from 'vitest';
import type { AlertState } from '../../src/types/Alert';
import { alertStateMock } from '../mocks/AlertState';
import { render, screen, act } from '@testing-library/react';
import AlertTable from '../../src/components/AlertTable';
import { MemoryRouter } from 'react-router';
import { getCurrentAlerts } from '../../src/api/GetCurrentAlerts';

export const alertMock: AlertState = alertStateMock;

vi.mock('../../src/components/AlertTableItem', () => ({
  default: () => <tr />,
}));

vi.mock('../../src/api/GetCurrentAlerts', () => ({
  getCurrentAlerts: vi.fn(),
}));

afterAll(() => {
  vi.useRealTimers();
});

describe('AlertTable', () => {
  it('shows loading state initially', async () => {
    vi.useFakeTimers();
    vi.mocked(getCurrentAlerts).mockReturnValue(new Promise(() => {}));
    render(
      <MemoryRouter>
        <AlertTable />
      </MemoryRouter>,
    );

    await act(async () => {
      vi.advanceTimersByTime(600);
    });
    expect(screen.getByText('Lädt aktuelle Alerts')).toBeInTheDocument();
    vi.useRealTimers();
  });

  it('renders error message on API failure', async () => {
    vi.useFakeTimers();
    vi.mocked(getCurrentAlerts).mockRejectedValue(
      new Error('Failed to fetch alert states'),
    );

    render(
      <MemoryRouter>
        <AlertTable />
      </MemoryRouter>,
    );
    await act(async () => {
      vi.advanceTimersByTime(600);
    });

    expect(
      screen.getByText((content) =>
        content.includes('Fehler: Failed to fetch alert states'),
      ),
    ).toBeInTheDocument();
  });
  it('renders table header', async () => {
    vi.useFakeTimers();
    vi.mocked(getCurrentAlerts).mockResolvedValue([alertMock]);

    render(
      <MemoryRouter>
        <AlertTable />
      </MemoryRouter>,
    );
    await act(async () => {
      vi.advanceTimersByTime(600);
    });

    expect(screen.getByText('Name')).toBeInTheDocument();
    expect(screen.getByText('Instance')).toBeInTheDocument();
    expect(screen.getByText('Job')).toBeInTheDocument();
  });
});
