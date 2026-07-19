export interface AlertHistory {
  fingerprint: string;
  status: string;
  alertName: string;
  severity: string;
  environment: string;
  instance: string;
  job: string;
  startsAt: Date;
  receivedAt: Date;
}
