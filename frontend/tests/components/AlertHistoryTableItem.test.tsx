import AlertHistoryTableItem from './../../src/components/AlertHistoryTableItem';
import { render, screen } from '@testing-library/react';
import { it, expect, describe } from 'vitest';
import { alertEventMock } from '../mocks/AlertEvent';

describe('AlertHistoryTableItem', () => {
  it('renders table rows', () => {
    const event = alertEventMock;

    render(
      <table>
        <tbody>
          <AlertHistoryTableItem event={event} />
        </tbody>
      </table>,
    );

    expect(screen.getByText(event.alertName)).toBeInTheDocument();
    expect(screen.getByText(event.job)).toBeInTheDocument();
    expect(screen.getByText(event.severity)).toBeInTheDocument();
    expect(screen.getByText(event.status)).toBeInTheDocument();
    expect(screen.getAllByRole('cell')[3].textContent).not.toBe('');
    expect(screen.getAllByRole('cell')[4].textContent).not.toBe('');
  });
});
