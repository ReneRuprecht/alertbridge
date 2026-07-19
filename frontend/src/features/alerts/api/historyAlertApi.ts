import { AlertHistoryMapper } from "../domain/alertHistoryMapper";
import type { AlertHistory } from "../types/alertHistory";
import type { AlertHistoryResponse } from "../types/historyAlertsResponse";

export async function getAlertHistory(instance: string): Promise<AlertHistory[]> {
  const response = await fetch(`/api/v1/alerts/history?instance=${instance}`);

  if (!response.ok) {
    throw new Error('Failed to load alert history');
  }

  const body: AlertHistoryResponse = await response.json();

  return body.alerts.map(AlertHistoryMapper.toAlertHistory);
}
