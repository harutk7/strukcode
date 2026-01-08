#!/usr/bin/env bun
// Test script to check provider database
import { ModelsDev } from "./packages/strukcode/src/provider/models"

console.log("Fetching provider database...")
const database = await ModelsDev.get()

console.log("\n=== Provider IDs in database ===")
const providerIds = Object.keys(database)
console.log(`Total providers: ${providerIds.length}`)
console.log(providerIds.sort().join("\n"))

console.log("\n=== Checking for 'strukcode' provider ===")
if (database["strukcode"]) {
  console.log("✓ strukcode provider EXISTS in database")
  console.log("Details:", JSON.stringify(database["strukcode"], null, 2))
} else {
  console.log("✗ strukcode provider MISSING from database")
  console.log("This will cause the TypeError when custom loader is called")
}

console.log("\n=== Simulating custom loader call ===")
const input = database["strukcode"]
console.log("input =", input)
if (input === undefined) {
  console.log("Attempting to access input.env will crash...")
  try {
    // @ts-ignore
    const env = input.env
    console.log("Unexpectedly succeeded:", env)
  } catch (error) {
    console.log("✓ Caught expected error:", error.message)
  }
}
