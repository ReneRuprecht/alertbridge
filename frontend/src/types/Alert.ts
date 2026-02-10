export interface AlertState {
  fingerprint: string;
  alertName: string;
  environment: string;
  instance: string;
  job: string;
  severity: 'CRITICAL' | 'WARNING' | 'INFO' | string;
  status: 'Firing' | 'Resolved' | string;
  startsAt: string;
}

export interface AlertHistory {
  fingerprint: string;
  instance: string;
  events: AlertEvent[];
}

export interface AlertEvent {
  alertName: string;
  job: string;
  severity: 'CRITICAL' | 'WARNING' | 'INFO' | string;
  status: 'Firing' | 'Resolved' | string;
  startsAt: string;
  receivedAt: string;
}
