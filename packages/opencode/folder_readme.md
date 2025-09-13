# packages/opencode - Core CLI and Server

## 📦 Overview
This is the **heart of OpenCode** - the main TypeScript package containing the CLI application, HTTP server, and all core functionality. This package handles AI interactions, file operations, session management, and provides the API that other components connect to.

## 🎯 Purpose
- **CLI Interface**: Command-line interface for running OpenCode commands
- **HTTP Server**: REST API server that TUI and other clients connect to
- **AI Integration**: Manages connections to various AI providers (Anthropic, OpenAI, etc.)
- **Session Management**: Handles conversation sessions, history, and state
- **Tool System**: Implements tools that AI can use (file editing, command execution, etc.)

## 📁 Key Directories

### `/src` - Source Code Root
The main source directory containing all TypeScript code.

#### Core Modules:
- **`/agent`** - AI agent configurations and management
  - Different modes (build, plan, general)
  - Agent prompt templates and behavior definitions
  
- **`/auth`** - Authentication and credential management
  - API key storage and retrieval
  - Provider authentication handling
  
- **`/bus`** - Event bus system for internal communication
  - Pub/sub pattern for component communication
  - Real-time event streaming
  
- **`/cli`** - Command-line interface implementation
  - `/cmd` - Individual CLI commands (run, serve, auth, etc.)
  - Uses Yargs for command parsing
  
- **`/command`** - Custom command system
  - User-defined command templates
  - Command execution and parameter substitution
  
- **`/config`** - Configuration management
  - Reads and validates `opencode.json`
  - User preferences and settings
  
- **`/file`** - File system operations
  - File reading/writing with diff tracking
  - Ripgrep integration for fast searching
  - Git operations for version control
  
- **`/format`** - Code formatting integrations
  - Prettier, Black, Ruff support
  - Language-specific formatters
  
- **`/lsp`** - Language Server Protocol client
  - Code intelligence features
  - Symbol search and navigation
  
- **`/mcp`** - Model Context Protocol implementation
  - External tool integration via MCP servers
  - Local and remote MCP server support
  
- **`/permission`** - Permission system
  - User consent for file edits and commands
  - Security controls for AI actions
  
- **`/plugin`** - Plugin system
  - Dynamic plugin loading
  - Plugin API and lifecycle
  
- **`/provider`** - AI provider integrations
  - Anthropic, OpenAI, Google, Amazon Bedrock, etc.
  - Provider abstraction layer
  - Model management and selection
  
- **`/server`** - HTTP API server
  - RESTful endpoints
  - OpenAPI documentation
  - WebSocket/SSE for real-time updates
  
- **`/session`** - Session management
  - Conversation history
  - Message handling
  - Session persistence and recovery
  
- **`/storage`** - Data persistence layer
  - File-based storage system
  - Session and configuration storage
  
- **`/tool`** - AI tool implementations
  - File editing tools
  - Command execution tools
  - Web fetching tools
  - Custom tool registry
  
- **`/util`** - Utility functions
  - Logging system
  - Error handling
  - Common helpers

### Key Files:
- **`index.ts`** - Main entry point for the CLI
- **`package.json`** - Package configuration and dependencies
- **`tsconfig.json`** - TypeScript configuration
- **`AGENTS.md`** - Development guidelines for the codebase

## 🚀 Main Features

1. **Multi-Provider AI Support**
   - Seamless switching between AI providers
   - Unified interface for different models
   - Streaming response support

2. **Intelligent File Operations**
   - Smart diff application
   - Syntax-aware editing
   - Automatic formatting

3. **Session Management**
   - Persistent conversation history
   - Session branching and merging
   - Export/import capabilities

4. **Extensible Tool System**
   - Built-in tools for common operations
   - Custom tool registration
   - Safe execution with permissions

5. **Real-time Communication**
   - Server-Sent Events for updates
   - WebSocket support for interactive features
   - Event bus for component communication

## 🔧 Technologies Used
- **TypeScript** - Primary language
- **Bun** - JavaScript runtime and package manager
- **Hono** - Lightweight web framework
- **Zod** - Schema validation
- **Yargs** - CLI argument parsing
- **AI SDK** - Vercel's AI SDK for provider abstraction

## 📝 Configuration
The package uses several configuration files:
- `opencode.json` - User configuration
- `.env` - Environment variables for API keys
- `auth.json` - Stored credentials (in ~/.local/share/opencode/)

## 🔌 API Endpoints
Key server endpoints include:
- `/session` - Session management
- `/provider` - AI provider operations  
- `/file` - File operations
- `/tool` - Tool execution
- `/config` - Configuration management

## 💡 Development Notes
- Uses namespace pattern for code organization
- Implements Result pattern for error handling
- Extensive use of Zod for runtime validation
- Event-driven architecture for loose coupling