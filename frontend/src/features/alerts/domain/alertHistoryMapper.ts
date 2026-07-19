import type { AlertHistory } from '../types/alertHistory';
import type { AlertHistoryItemResponse } from '../types/historyAlertsResponse';

export class AlertHistoryMapper {
  static toAlertHistory(response: AlertHistoryItemResponse): AlertHistory {
    return {
      fingerprint: response.fingerprint,
      status: response.status,
      alertName: response.alert_name,
      severity: response.severity,
      environment: response.environment,
      instance: response.instance,
      job: response.job,
      startsAt: new Date(response.starts_at),
      receivedAt: new Date(response.received_at),
    };
  }
}
