import type { RulesResponse } from "../types/Rule";

export async function getRules(): Promise<RulesResponse> {
  const base_api_url = import.meta.env.VITE_API_BASE_URL;
  const response = await fetch(base_api_url + "/api/v1/rules");
  if (!response.ok) throw new Error("Failed to fetch rules");
  return response.json();
}
