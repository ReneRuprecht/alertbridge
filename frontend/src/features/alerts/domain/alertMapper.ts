import type { Alert } from "../types/alert";
import type { CurrentAlertItemResponse } from "../types/currentAlertsResponse";

export class AlertMapper {
  static toAlert(response: CurrentAlertItemResponse): Alert {
    return {
      fingerprint: response.fingerprint,
      status: response.status,
      alertName: response.alert_name,
      severity: response.severity,
      environment: response.environment,
      instance: response.instance,
      job: response.job,
      startsAt: response.starts_at,
      lastUpdatedAt: response.last_updated_at,
    };
  }
}
