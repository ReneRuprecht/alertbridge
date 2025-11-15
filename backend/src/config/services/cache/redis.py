import redis.asyncio as redis

_redis_cache_pool: redis.ConnectionPool | None = None
_redis_cache_client: redis.Redis | None = None


async def get_redis_cache_pool() -> redis.ConnectionPool:
    global _redis_cache_pool
    if not _redis_cache_pool:
        _redis_cache_pool = redis.ConnectionPool.from_url(
            url="redis://localhost:6379/0"
        )

    return _redis_cache_pool


async def get_redis_cache_client() -> redis.Redis:
    global _redis_cache_client
    pool = await get_redis_cache_pool()
    if not _redis_cache_client:
        _redis_cache_client = redis.Redis(
            connection_pool=pool, decode_responses=True
        )

    return _redis_cache_client
