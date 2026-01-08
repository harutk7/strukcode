from strukcode_ai import StrukCodeClient

client = StrukCodeClient()
files = client.file_status() or []
for f in files:
    print(f.path, f.type)
