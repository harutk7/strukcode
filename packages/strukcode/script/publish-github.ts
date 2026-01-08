#!/usr/bin/env bun
import { $ } from "bun"
import pkg from "../package.json"
import { Script } from "@strukcode-ai/script"
import { fileURLToPath } from "url"

const dir = fileURLToPath(new URL("..", import.meta.url))
process.chdir(dir)

// Check if binaries exist, if not build them
const { binaries } = await import("./build.ts")

console.log("Creating archives for GitHub release...")

for (const key of Object.keys(binaries)) {
  if (key.includes("linux")) {
    await $`cd dist/${key}/bin && tar -czf ../../${key}.tar.gz *`
  } else {
    await $`cd dist/${key}/bin && zip -r ../../${key}.zip *`
  }
  console.log(`✓ Created archive for ${key}`)
}

console.log("\n✓ All archives created successfully!")
console.log("\nArchives location: packages/strukcode/dist/")
console.log("\nTo create a GitHub release, run:")
console.log(`  gh release create v${Script.version} --title "v${Script.version}" --notes "Release notes here" ./packages/strukcode/dist/*.zip ./packages/strukcode/dist/*.tar.gz`)
