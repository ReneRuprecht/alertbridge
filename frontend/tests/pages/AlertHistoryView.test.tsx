import { act, fireEvent, render, screen } from '@testing-library/react';
import AlertHistoryView from '../../src/pages/AlertHistoryView';
import { MemoryRouter, Route, Routes } from 'react-router';
import {
  alertHistoryMock,
  alertHistoryWithEmptyEventsMock,
} from '../mocks/AlertHistory';
import { getHistoryAlert } from '../../src/api/GetHistoryAlert';

vi.mock('../../src/components/AlertHistoryTable', () => ({
  default: () => <div />,
}));

vi.mock('../../src/api/GetHistoryAlert', async () => ({
  getHistoryAlert: vi.fn(),
}));

vi.mock('../../src/utils/formatter', () => ({
  formatAlertInstance: vi.fn(),
}));

const mockedUsedNavigate = vi.fn();

beforeEach(() => {
  vi.mock(import('react-router'), async (importOriginal) => {
    const actual = await importOriginal();
    return {
      ...actual,
      useNavigate: () => mockedUsedNavigate,
    };
  });
});

afterEach(() => {
  vi.useRealTimers();
});

describe('AlertHistoryView', () => {
  const alertHistory = alertHistoryMock;

  it('renders loading', async () => {
    vi.useFakeTimers();
    vi.mocked(getHistoryAlert).mockReturnValue(new Promise(() => {}));

    render(
      <MemoryRouter initialEntries={[`/history/${alertHistory.instance}`]}>
        <Routes>
          <Route
            path='/history/:alertInstance'
            element={<AlertHistoryView />}
          />
        </Routes>
      </MemoryRouter>,
    );

    await act(async () => {
      vi.advanceTimersByTime(600);
    });

    expect(
      screen.getByText(`Lädt History von ${alertHistory.instance}`),
    ).toBeInTheDocument();

    vi.useRealTimers();
  });

  it('renders error from api', async () => {
    const errorText = 'Network Error';
    vi.useFakeTimers();
    vi.mocked(getHistoryAlert).mockRejectedValue(new Error(errorText));

    render(
      <MemoryRouter initialEntries={[`/history/${alertHistory.instance}`]}>
        <Routes>
          <Route
            path='/history/:alertInstance'
            element={<AlertHistoryView />}
          />
        </Routes>
      </MemoryRouter>,
    );

    await act(async () => {
      vi.advanceTimersByTime(600);
    });

    expect(
      screen.getByText((content) => content.includes(`Fehler: ${errorText}`)),
    ).toBeInTheDocument();

    vi.useRealTimers();
  });

  it('renders missing alertInstance error if param is missing', async () => {
    const errorText = 'Missing alertInstance';
    vi.useFakeTimers();

    render(
      <MemoryRouter initialEntries={['/history/']}>
        <Routes>
          <Route
            path='/history/:alertInstance?'
            element={<AlertHistoryView />}
          />
        </Routes>
      </MemoryRouter>,
    );

    await act(async () => {
      vi.advanceTimersByTime(600);
    });

    expect(
      screen.getByText((content) => content.includes(`Fehler: ${errorText}`)),
    ).toBeInTheDocument();
  });

  it('renders missing alertInstance error if param if undefined', async () => {
    const errorText = 'Missing alertInstance';
    vi.useFakeTimers();

    render(
      <MemoryRouter initialEntries={['/history/undefined']}>
        <Routes>
          <Route
            path='/history/:alertInstance?'
            element={<AlertHistoryView />}
          />
        </Routes>
      </MemoryRouter>,
    );

    await act(async () => {
      vi.advanceTimersByTime(600);
    });

    expect(
      screen.getByText((content) => content.includes(`Fehler: ${errorText}`)),
    ).toBeInTheDocument();
  });

  it('renders correct instance headline', async () => {
    vi.mocked(getHistoryAlert).mockResolvedValue(alertHistory);
    vi.unmock('../../src/utils/formatter');
    vi.useFakeTimers();

    render(
      <MemoryRouter initialEntries={[`/history/${alertHistory.instance}`]}>
        <Routes>
          <Route
            path='/history/:alertInstance'
            element={<AlertHistoryView />}
          />
        </Routes>
      </MemoryRouter>,
    );
    await act(async () => {
      vi.advanceTimersByTime(600);
    });

    expect(screen.getByText(alertHistory.instance)).toBeInTheDocument();

    const backButton = screen.getByText('BACK');
    fireEvent.click(backButton);
    expect(mockedUsedNavigate).toHaveBeenCalledWith('/');
  });

  it('renders missing history error', async () => {
    vi.mocked(getHistoryAlert).mockResolvedValue(undefined as any);
    vi.unmock('../../src/utils/formatter');
    vi.useFakeTimers();

    render(
      <MemoryRouter initialEntries={[`/history/${alertHistory.instance}`]}>
        <Routes>
          <Route
            path='/history/:alertInstance'
            element={<AlertHistoryView />}
          />
        </Routes>
      </MemoryRouter>,
    );

    await act(async () => {
      vi.advanceTimersByTime(600);
    });
    expect(screen.getByText('Keine History gefunden')).toBeInTheDocument();
  });

  it('renders missing event history error', async () => {
    vi.mocked(getHistoryAlert).mockResolvedValue(
      alertHistoryWithEmptyEventsMock,
    );
    vi.unmock('../../src/utils/formatter');
    vi.useFakeTimers();

    render(
      <MemoryRouter initialEntries={[`/history/${alertHistory.instance}`]}>
        <Routes>
          <Route
            path='/history/:alertInstance'
            element={<AlertHistoryView />}
          />
        </Routes>
      </MemoryRouter>,
    );

    await act(async () => {
      vi.advanceTimersByTime(1000);
    });
    expect(
      screen.getByText((content) =>
        content.includes(`Keine Event History für ${alertHistory.instance}`),
      ),
    ).toBeInTheDocument();
  });
});
