from dataclasses import dataclass
from typing import Any, Dict


@dataclass
class FrontendAlert:
    alertname: str
    status: str
    fingerprint: str
    instance: str

    @staticmethod
    def from_dict(frontend_alert: Dict[str, Any]):
        return FrontendAlert(
            fingerprint=frontend_alert.get("fingerprint", ""),
            status=frontend_alert.get("status", ""),
            alertname=frontend_alert.get("alertname", ""),
            instance=frontend_alert.get("instance", ""),
        )
