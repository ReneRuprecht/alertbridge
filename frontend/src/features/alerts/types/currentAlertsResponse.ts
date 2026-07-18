export interface CurrentAlertItemResponse {
  fingerprint: string;
  status: string;
  alert_name: string;
  severity: string;
  environment: string;
  instance: string;
  job: string;
  starts_at: string;
  last_updated_at: string;
}

export interface CurrentAlertsResponse {
    alerts: CurrentAlertItemResponse[]
}
