export interface AlertState {
  fingerprint: string;
  alertName: string;
  environment: string;
  instance: string;
  job: string;
  severity: "CRITICAL" | "WARNING" | "INFO" | string;
  status: "Firing" | "Resolved" | string;
  startsAt: string;
}
