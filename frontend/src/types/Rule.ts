export interface RulesResponse {
  rules: Rule[];
}
export interface Rule {
  id: string;
  name: string;
  description: string;
  priority: number;
  enabled: boolean;
}

export interface CreateRuleRequest {
  name: string;
  description: string;
  priority: number;
  enabled: boolean;
}
