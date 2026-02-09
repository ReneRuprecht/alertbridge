import { act, fireEvent, render, screen, waitFor } from '@testing-library/react';
import AlertHistoryView from '../../src/pages/AlertHistoryView';
import { MemoryRouter, Route, Routes } from 'react-router';
import { alertHistoryMock } from '../mocks/AlertHistory';
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
  vi.mock(import("react-router"), async (importOriginal) => {
    const actual = await importOriginal()
    return {
      ...actual,
      useNavigate: () => mockedUsedNavigate
    }
  })
});

describe('AlertHistoryView', () => {
  const alertHistory = alertHistoryMock;

  it('renders loading', async () => {
    vi.useFakeTimers()
    vi.mocked(getHistoryAlert).mockReturnValue(new Promise(() => { }));

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
    ).toBeInTheDocument()

    vi.useRealTimers()
  });

  it('renders error', async () => {
    const errorText = 'Network Error';
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

    await waitFor(() =>
      expect(
        screen.getByText((content) => content.includes(`Fehler: ${errorText}`)),
      ).toBeInTheDocument(),
    );
  });

  it('renders correct instance headline', async () => {
    vi.mocked(getHistoryAlert).mockResolvedValue(alertHistory);
    vi.unmock('../../src/utils/formatter')

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

    await waitFor(() => {
      expect(screen.getByText(alertHistory.instance)).toBeInTheDocument();

      const backButton = screen.getByText('BACK');
      fireEvent.click(backButton);
      expect(mockedUsedNavigate).toHaveBeenCalledWith('/');

    });
  });
});
