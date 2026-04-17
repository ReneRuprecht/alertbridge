export interface Alert {
  fingerprint: string;
  alertName: string;
  instance: string;
  status: "firing" | "resolved";
  startsAt: string;
}

export interface ActiveAlerts {
  alerts: Alert[];
}
