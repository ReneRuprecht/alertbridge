import type { AlertEvent } from '../../src/types/Alert';

export const alertEventMock: AlertEvent = {
  alertName: 'CPU High',
  job: 'Server-1',
  severity: 'Critical',
  startsAt: '2026-02-08T10:00:00Z',
  receivedAt: '2026-02-08T10:05:00Z',
  status: 'Active',
};
