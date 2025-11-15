from datetime import datetime
from typing import Dict, Optional

from pydantic import BaseModel


class Alert(BaseModel):
    alertname: str
    status: str
    labels: Dict[str, str]
    fingerprint: str

    starts_at: Optional[datetime] = None
    ended_at: Optional[datetime] = None
    updated_at: Optional[datetime] = None
