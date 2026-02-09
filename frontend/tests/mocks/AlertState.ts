import type { AlertState } from '../../src/types/Alert';

export const alertStateMock: AlertState = {
  fingerprint: 'abc123',
  alertName: 'CPU High',
  environment: 'test-env',
  instance: 'server-1',
  job: 'Job1',
  severity: 'CRITICAL',
  startsAt: '2026-02-08T10:00:00Z',
  status: 'Firing',
};
