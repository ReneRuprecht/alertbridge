import { render, screen, waitFor } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';
import AlertTableView from '../../src/pages/AlertTableView';
import { getCurrentAlerts } from '../../src/api/GetCurrentAlerts';
import { alertStateMock } from '../mocks/AlertState';
import { MemoryRouter } from 'react-router';

vi.mock('../../src/components/AlertTable', () => ({
  default: () => <div />,
}));

describe('AlertTableView', () => {
  const headline = 'Aktuelle Alerts'

  it('renders headline', async () => {
    render(
      <AlertTableView />
    );

    expect(screen.getByText(headline)).toBeInTheDocument()
  })
})


vi.mock('../../src/api/GetCurrentAlerts', () => ({
  getCurrentAlerts: vi.fn(),
}));

describe('AlertTableView Integration', () => {
  vi.unmock('../../src/components/AlertTable')
  vi.mocked(getCurrentAlerts).mockResolvedValue([alertStateMock])

  it('renders current alerts', async () => {
    render(
      <MemoryRouter>      <AlertTableView />
      </MemoryRouter>

    );


    await waitFor(() => {
      expect(screen.getByText('CPU High')).toBeInTheDocument();
      expect(screen.getByText('server-1')).toBeInTheDocument();
      expect(screen.getByText('Job1')).toBeInTheDocument();
      expect(screen.getByText('CRITICAL')).toBeInTheDocument();
      expect(screen.getByText('Firing')).toBeInTheDocument();
    })




  })

})