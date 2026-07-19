export interface AlertHistoryItemResponse {
  fingerprint: string;
  status: string;
  alert_name: string;
  severity: string;
  environment: string;
  instance: string;
  job: string;
  starts_at: string;
  received_at: string;
}

export interface AlertHistoryResponse {
  alerts: AlertHistoryItemResponse[];
}
