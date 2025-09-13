# packages/sdk - TypeScript SDK

## 📦 Overview
This package contains the **TypeScript SDK for OpenCode**, providing a type-safe client library for interacting with the OpenCode server API. The SDK is auto-generated using Stainless and provides full TypeScript support with comprehensive type definitions.

## 🎯 Purpose
- **API Client Library**: Type-safe client for OpenCode server
- **Auto-generated Code**: Generated from OpenAPI specifications
- **TypeScript Support**: Full type definitions and IntelliSense
- **Multiple Environments**: Support for different server endpoints
- **Stream Support**: Handle Server-Sent Events and streaming responses

## 📁 Directory Structure

### `/js` - JavaScript/TypeScript SDK
The main SDK implementation.

#### Key Directories:
- **`/src`** - Source code
  - `/gen` - Auto-generated code from Stainless
    - `sdk.gen.ts` - Main SDK client class
    - `types.gen.ts` - TypeScript type definitions
    - HTTP client wrappers
  - `/client` - Custom client implementations
  - `/utils` - Utility functions

- **`/script`** - Build and generation scripts
  - `generate.ts` - SDK generation script

### `/stainless` - Stainless Configuration
Configuration for SDK generation:
- API specification
- Generation settings
- Custom templates

## 🚀 Main Features

1. **Type-Safe API Client**
   - Full TypeScript types for all endpoints
   - Request/response validation
   - Automatic serialization/deserialization

2. **Comprehensive API Coverage**
   - Session management
   - File operations
   - Provider configuration
   - Tool execution
   - TUI control

3. **Streaming Support**
   - Server-Sent Events handling
   - Real-time message streaming
   - Progress updates

4. **Error Handling**
   - Typed error responses
   - Retry logic
   - Timeout configuration

## 🔧 SDK Classes & Methods

### Main Client: `OpencodeClient`
```typescript
const client = new OpencodeClient({
  baseUrl: 'http://localhost:4096'
})
```

### Key API Groups:
- **`client.session`** - Session operations
  - `.create()` - Create new session
  - `.get()` - Get session details
  - `.message()` - Send messages
  - `.abort()` - Cancel operations

- **`client.file`** - File operations
  - `.list()` - List files
  - `.read()` - Read file content
  - `.status()` - Get file status

- **`client.config`** - Configuration
  - `.get()` - Get current config
  - `.providers()` - List providers

- **`client.tui`** - TUI control
  - `.appendPrompt()` - Add to prompt
  - `.submitPrompt()` - Submit input
  - `.showToast()` - Show notifications

- **`client.auth`** - Authentication
  - `.set()` - Set credentials

## 📝 Type Definitions

Key types exported:
- `Session` - Session information
- `Message` - Message structure
- `Provider` - AI provider config
- `Model` - Model information
- `Config` - Configuration schema
- `Agent` - Agent definition
- `Tool` - Tool registration

## 🔌 Usage Example

```typescript
import { OpencodeClient } from '@opencode-ai/sdk'

const client = new OpencodeClient({
  baseUrl: 'http://localhost:4096'
})

// Create a session
const session = await client.session.create({
  title: 'New Session'
})

// Send a message
const response = await client.session.message({
  sessionId: session.id,
  content: 'Hello, OpenCode!'
})

// Stream responses
for await (const event of response) {
  console.log(event)
}
```

## 🔧 Technologies Used
- **TypeScript** - Primary language
- **Stainless** - SDK generation
- **Zod** - Schema validation
- **Hey API** - HTTP client base

## 📝 Key Files
- **`package.json`** - Package configuration
- **`tsconfig.json`** - TypeScript settings
- **`README.md`** - Usage documentation
- **`.npmignore`** - NPM publish configuration

## 🚦 Generation Process

```bash
# Generate SDK from OpenAPI spec
bun run generate

# This runs:
# 1. Fetches OpenAPI spec from server
# 2. Runs Stainless generation
# 3. Outputs to src/gen/
```

## 💡 Development Notes
- SDK is auto-generated - don't edit gen/ files directly
- Custom functionality goes in non-generated directories
- Stainless configuration controls generation behavior
- Version synced with main OpenCode package