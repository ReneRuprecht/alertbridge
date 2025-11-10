import json
import logging
import os
from typing import List, Optional, Tuple

import psycopg2
from src.models.alert import Alert


class PSQLClient:
    def __init__(
        self,
        dbname: str = os.getenv("POSTGRES_DB", "postgres"),
        user: str = os.getenv("POSTGRES_USER", "postgres"),
        password: str = os.getenv("POSTGRES_PASSWORD", "postgres"),
        host: str = os.getenv("POSTGRES_HOST", "localhost"),
        port: int = int(os.getenv("POSTGRES_PORT", 5432)),
        logger: Optional[logging.Logger] = None,
    ):
        self.logger = logger or logging.getLogger("PSQLClient")
        self.dbname = dbname
        self.user = user
        self.password = password
        self.host = host
        self.port = port
        self.conn = self._connect()
        """
        Initialize a PSQLClient instance.
        """

    def _connect(self):
        """
        Establish a new connection to the Postgres database.
        """
        try:
            conn = psycopg2.connect(
                dbname=self.dbname,
                user=self.user,
                password=self.password,
                host=self.host,
                port=self.port,
            )
            conn.autocommit = True
            self.logger.info("Connected to %s:%d/%s", self.host, self.port, self.dbname)
            return conn
        except Exception as e:
            self.logger.error("Connection failed: %s", e)
            raise

    def _ensure_connection(self):
        """
        Check if the current database connection is alive and reconnect if needed.
        """
        try:
            with self.conn.cursor() as cur:
                cur.execute("SELECT 1;")
        except Exception:
            self.logger.error("Connection lost, attempting reconnect...")
            self.conn = self._connect()

    def save_alert(self, alert: Alert):
        """
        Save an Alert into the database
        """

        ends_at_value = alert.ends_at if alert.status == "resolved" else None

        query = """
        INSERT INTO alerts_history (
            alertname, status, labels, fingerprint,
            starts_at, ends_at
        )
        VALUES (%s, %s, %s, %s, %s, %s)
        ON CONFLICT (fingerprint, status, labels, starts_at) DO NOTHING;
        """

        try:
            self._ensure_connection()
            with self.conn.cursor() as cur:
                cur.execute(
                query,
                    (
                        alert.alertname,
                        alert.status,
                        json.dumps(alert.labels),
                        alert.fingerprint,
                        alert.starts_at,
                        ends_at_value,
                    ),
                )
        except Exception as e:
            self.logger.error("Error while saving alert: %s", e)
            self.conn.rollback()

    def read_alerts(self) -> List[Tuple]:
        """
        Read alerts from database
        """
        self._ensure_connection()
        try:
            with self.conn.cursor() as cur:
                cur.execute("SELECT * FROM alerts_history;")
                rows = cur.fetchall()
            return rows
        except Exception as e:
            self.logger.error("Error while reading alerts: %s", e)
            return []
