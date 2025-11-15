from psycopg_pool import AsyncConnectionPool

_pg_pool: AsyncConnectionPool | None = None


async def get_postgres_db_pool() -> AsyncConnectionPool:
    global _pg_pool

    if not _pg_pool:
        _pg_pool = AsyncConnectionPool(
            conninfo="postgresql://root:root@localhost:5432/alerts_db",
            min_size=1,
            max_size=10,
            open=False,
        )

    return _pg_pool
