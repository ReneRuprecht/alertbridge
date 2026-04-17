export interface Alert {
  fingerprint: string;
  alertName: string;
  instance: string;
  job: string;
  status: "firing" | "resolved";
  startsAt: string;
}

export interface ActiveAlerts {
  alerts: Alert[];
}

export interface AlertHistory {
  instance: string;
  alerts: HistoryAlert[];
}

export interface HistoryAlert {
  fingerprint: string;
  alertName: string;
  status: "firing" | "resolved";
  receivedAt: string;
  startsAt: string;
  job: string;
  description: string;
  severity: string;
}
