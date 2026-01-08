# Streaming (SSE)

Subscribe to the event stream. The wrapper provides both sync and async interfaces.

```python
from strukcode_ai import StrukCodeClient

client = StrukCodeClient()

# Sync streaming
for event in client.subscribe_events():
    print(event)
    break
```

Async variant:

```python
import asyncio
from strukcode_ai import StrukCodeClient

async def main():
    client = StrukCodeClient()
    async for event in client.subscribe_events_async():
        print(event)
        break

asyncio.run(main())
```
