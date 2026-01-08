# Sessions

List sessions and inspect them. The wrapper exposes a convenience method while the generated API remains available under `strukcode_ai.api.default`.

```python
from strukcode_ai import StrukCodeClient
from strukcode_ai.api.default import session_list as generated

client = StrukCodeClient()

# Wrapper
sessions = client.list_sessions() or []

# Generated function
sessions2 = generated.sync(client=client.client)

print(len(sessions), len(sessions2))
```
