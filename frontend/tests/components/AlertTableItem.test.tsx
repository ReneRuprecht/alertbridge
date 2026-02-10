import { render, screen, fireEvent } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';
import AlertTableItem from '../../src/components/AlertTableItem';
import type { AlertState } from '../../src/types/Alert';
import { alertStateMock } from '../mocks/AlertState';

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

describe('AlertTableItem', () => {
  const alertMock: AlertState = alertStateMock;
  it('renders all cells and calls navigate on row click', () => {
    render(
      <table>
        <tbody>
          <AlertTableItem alert={alertMock} />
        </tbody>
      </table>,
    );

    expect(screen.getByText('CPU High')).toBeInTheDocument();
    expect(screen.getByText('server-1')).toBeInTheDocument();
    expect(screen.getByText('Job1')).toBeInTheDocument();
    expect(screen.getByText('CRITICAL')).toBeInTheDocument();
    expect(screen.getByText('Firing')).toBeInTheDocument();

    const row = screen.getByText('CPU High').closest('tr')!;
    fireEvent.click(row);
    expect(mockedUsedNavigate).toHaveBeenCalledWith('/history/server-1');
  });
});
