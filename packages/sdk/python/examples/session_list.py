from strukcode_ai import StrukCodeClient

client = StrukCodeClient()
print([s.id for s in client.list_sessions() or []])
