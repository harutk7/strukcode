interface ImportMetaEnv {
  readonly VITE_STRUKCODE_SERVER_HOST: string
  readonly VITE_STRUKCODE_SERVER_PORT: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
