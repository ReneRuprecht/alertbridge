import type { AlertHistory } from '../../src/types/Alert';

export const alertHistoryMock: AlertHistory = {
  fingerprint: 'abc123',
  instance: 'server-1',
  events: [
    {
      alertName: 'CPU High',
      job: 'Server-1',
      severity: 'CRITICAL',
      status: 'Firing',
      startsAt: '2026-02-08T10:00:00Z',
      receivedAt: '2026-02-08T10:05:00Z',
    },
    {
      alertName: 'Memory Low',
      job: 'Server-2',
      severity: 'WARNING',
      status: 'Resolved',
      startsAt: '2026-02-07T08:30:00Z',
      receivedAt: '2026-02-07T08:35:00Z',
    },
  ],
};
