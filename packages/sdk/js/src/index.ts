export * from "./client.js"
export * from "./server.js"

import { createStrukcodeClient } from "./client.js"
import { createStrukcodeServer } from "./server.js"
import type { ServerOptions } from "./server.js"

export async function createStrukcode(options?: ServerOptions) {
  const server = await createStrukcodeServer({
    ...options,
  })

  const client = createStrukcodeClient({
    baseUrl: server.url,
  })

  return {
    client,
    server,
  }
}
