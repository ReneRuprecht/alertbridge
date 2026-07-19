import { AlertMapper } from '../domain/alertMapper';
import type { Alert } from '../types/alert';
import type { CurrentAlertsResponse } from '../types/currentAlertsResponse';

export async function getCurrentAlerts(): Promise<Alert[]> {
  const response = await fetch('/api/v1/alerts/current');

  if (!response.ok) {
    throw new Error('Failed to load current alerts');
  }

  const body: CurrentAlertsResponse = await response.json();

  return body.alerts.map(AlertMapper.toAlert);
}
