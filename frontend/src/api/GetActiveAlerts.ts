import type { ActiveAlerts } from "../types/Alert";

export async function getActiveAlerts(): Promise<ActiveAlerts> {
  const base_api_url = import.meta.env.VITE_API_BASE_URL;
  const response = await fetch(base_api_url + "/api/v1/alerts");
  if (!response.ok) throw new Error("Failed to fetch alerts");
  return response.json();
}
