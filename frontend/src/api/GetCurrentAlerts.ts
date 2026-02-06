import type { AlertState } from "../types/Alert";

export async function getCurrentAlerts(): Promise<AlertState[]> {
  const base_api_url = import.meta.env.VITE_API_BASE_URL;
  const response = await fetch(base_api_url + "/api/v1/alert-states");
  if (!response.ok) throw new Error("Failed to fetch alert states");
  return response.json();
}
