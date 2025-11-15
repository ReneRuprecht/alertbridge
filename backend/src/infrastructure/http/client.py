import aiohttp

_http_client: aiohttp.ClientSession | None = None


async def get_http_client() -> aiohttp.ClientSession:
    global _http_client
    if _http_client is None or _http_client.closed:
        _http_client = aiohttp.ClientSession()
    return _http_client


async def shutdown_http_client():
    global _http_client
    if _http_client and not _http_client.closed:
        await _http_client.close()
    _http_client = None
