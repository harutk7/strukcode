#!/usr/bin/env bun

import { Script } from "@strukcode-ai/script"
import { $ } from "bun"

if (!Script.preview) {
  await $`gh release edit v${Script.version} --draft=false`
}

await $`bun install`

await $`gh release download --pattern "strukcode-linux-*64.tar.gz" --pattern "strukcode-darwin-*64.zip" -D dist`

await import(`../packages/strukcode/script/publish-registries.ts`)
