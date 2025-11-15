class ExtractAlertsUseCase:
    def __init__(self):
        pass

    async def execute(self, payload: dict) -> list[dict]:
        alerts_json = payload.get("alerts", [])

        if not isinstance(alerts_json, list) or not alerts_json:
            return []
        return alerts_json
