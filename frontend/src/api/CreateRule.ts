import type { CreateRuleRequest } from "../types/Rule";

export async function createRule(
  createRuleRequest: CreateRuleRequest,
): Promise<Response> {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(createRuleRequest),
  };

  const base_api_url = import.meta.env.VITE_API_BASE_URL;
  const response = await fetch(base_api_url + "/api/v1/rules", requestOptions);
  if (!response.ok) throw new Error("Failed to post rule");
  return response;
}
