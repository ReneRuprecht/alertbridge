export interface Alert {
  fingerprint: string;
  status: string;
  alertName: string;
  severity: string;
  environment: string;
  instance: string;
  job: string;
  startsAt: string;
  lastUpdatedAt: string;
}
