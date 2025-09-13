# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

OpenCode is an AI coding agent built for the terminal, designed to provide a powerful alternative to other AI coding assistants with a focus on open source, provider-agnostic design, and terminal-first experience. The project features a client/server architecture that allows for flexible deployment and use cases.

## Architecture

### Monorepo Structure
- **Root**: Main workspace configuration with Bun package manager
- **packages/opencode**: Core TypeScript CLI application and server
- **packages/tui**: Go-based Terminal User Interface (TUI) client
- **packages/web**: Astro-based documentation website
- **packages/sdk**: TypeScript SDK for API clients
- **cloud/**: Cloud infrastructure packages (app, core, function, resource, scripts)
- **infra/**: SST-based infrastructure configuration
- **github/**: GitHub Action implementation

### Key Components

#### Core Application (packages/opencode)
- Entry point: `src/index.ts` - Yargs-based CLI with commands
- CLI Commands: `src/cli/cmd/` - Individual command implementations (run, generate, auth, agent, etc.)
- Server: `src/server/` - API server with project management
- Providers: `src/provider/` - AI model provider integrations
- LSP Integration: `src/lsp/` - Language Server Protocol support
- MCP Support: `src/mcp/` - Model Context Protocol integration

#### Terminal UI (packages/tui)
- Go-based TUI client that communicates with the TypeScript server
- Provides rich terminal interface for interactive AI coding sessions
- Located in Go modules with custom input handling and app state management

#### Infrastructure
- **SST v3**: Infrastructure as Code using SST with Cloudflare Workers
- **Cloud Deployment**: Cloudflare-based with Durable Objects for sync
- **API**: Cloudflare Worker serving the main API
- **Documentation**: Astro-based docs site

## Development Commands

### Primary Development
```bash
# Install dependencies
bun install

# Start development server (main CLI)
bun dev

# Run specific package development
bun run --conditions=development packages/opencode/src/index.ts

# Test opencode CLI directly
cd packages/opencode && bun dev
```

### Code Quality
```bash
# Type checking across all packages
bun typecheck

# Format code (uses Prettier)
bun run script/format.ts

# Generate SDK clients
bun generate
```

### Building & Testing
```bash
# Build TUI (requires Go 1.24.x)
cd packages/tui && go build

# Run stats collection
bun run script/stats.ts
```

## Configuration

### Development Environment
- **Package Manager**: Bun 1.2.21 (specified in packageManager field)
- **TypeScript**: Configured with @tsconfig/bun base config
- **Node**: Requires Node.js 22.x (based on @tsconfig/node22)
- **Go**: Requires Go 1.24.x for TUI development

### Key Configuration Files
- **opencode.json**: Project-specific OpenCode configuration with MCP settings
- **sst.config.ts**: Infrastructure deployment configuration
- **bunfig.toml**: Bun package manager configuration (exact installs)
- **.editorconfig**: Code formatting standards

## Coding Standards (from AGENTS.md)

### Core Principles
- Keep functionality in single functions unless composable/reusable
- Avoid unnecessary destructuring of variables
- Minimize `else` statements - prefer early returns
- Avoid `try`/`catch` blocks where possible
- Avoid `any` type usage
- Prefer `const` over `let`
- Use single-word variable names when clear
- Leverage Bun APIs (e.g., `Bun.file()`) over Node.js equivalents

### Code Style
- **Prettier Configuration**: 120 character line length, no semicolons
- **Import Organization**: ES modules with explicit imports
- **Error Handling**: Use NamedError class for structured error handling

## Provider Integration

The system supports multiple AI providers:
- **Anthropic** (recommended)
- **OpenAI**
- **Google**
- **Amazon Bedrock**
- **Local models**

Provider configuration is handled through the provider system in `packages/opencode/src/provider/`.

## Infrastructure Deployment

### SST Configuration
- **Home Platform**: Cloudflare
- **Stages**: production (retained), development (removable)
- **Protection**: Production stage protected from accidental removal
- **Secrets**: GitHub App integration, Stripe for billing

### Cloud Architecture
- **API Worker**: Main application logic on Cloudflare Workers
- **Durable Objects**: SyncServer for real-time synchronization
- **Bucket**: File storage via Cloudflare R2
- **Domain Setup**: api.${domain} for API, docs.${domain} for documentation

## GitHub Integration

### Actions & Automation
- **opencode Workflow**: Responds to `/oc` and `/opencode` comments in PRs
- **CI Pipeline**: Typecheck, format checking, publishing
- **Release Process**: Automated through GitHub Actions
- **Stats Collection**: Automated download statistics tracking

### Development Workflow
- **Main Branch**: `dev` (not main/master)
- **PR Requirements**: Typecheck must pass
- **Code Formatting**: Automated via GitHub Actions

## Authentication & Security

- **GitHub App Integration**: For repository access and PR interactions
- **OpenAuth**: Authentication system (@openauthjs/openauth)
- **API Keys**: Environment-based configuration
- **Secrets Management**: SST Secret handling for sensitive data

## Contributing Guidelines

### Acceptable Contributions
- Bug fixes
- LLM performance improvements
- New provider support
- Environment-specific fixes
- Missing standard behavior
- Documentation improvements

### Core Feature Policy
- Fundamental features require design process with core team
- No PRs accepted for core architectural changes
- Focus on incremental improvements over major rewrites

## MCP (Model Context Protocol) Support

OpenCode includes MCP server capabilities for extending functionality:
- Configuration via `opencode.json`
- Local and remote MCP server support
- Weather example server included in default config

## Testing & Development Notes

### API Client Generation
Changes to TypeScript API endpoints in `packages/opencode/src/server/server.ts` require SDK regeneration by the OpenCode team using Stainless.

### Package Structure
- **Workspace Catalog**: Centralized dependency management
- **Workspace Packages**: `cloud/*`, `packages/*`, and `packages/sdk/js`
- **Patches**: Custom patches for dependencies (e.g., SolidJS Start)

### Development Tools
- **VSCode Support**: Configured workspace with appropriate TypeScript settings
- **Language Server**: Built-in LSP client/server for enhanced editor integration
- **File Watching**: Integrated file system watching for development workflow

## Performance Considerations

- **Bundle Size**: Optimized for terminal usage
- **Memory Usage**: Designed for efficient resource utilization
- **Startup Time**: Fast CLI startup through optimized dependency loading
- **Network**: Efficient API communication between TUI and server components

## External Dependencies

### Critical Dependencies
- **AI SDK**: Vercel AI SDK for provider abstraction
- **Hono**: Web framework for API server
- **Zod**: Runtime type validation and API schema
- **Tree-sitter**: Code parsing and analysis
- **yargs**: CLI argument parsing
- **SST**: Infrastructure deployment

### Platform-Specific
- **Bun Runtime**: Primary JavaScript runtime for development and execution
- **Go**: For TUI client implementation
- **Cloudflare**: Production deployment platform

## Distribution

### Installation Methods
- **Direct Install**: `curl -fsSL https://opencode.ai/install | bash`
- **Package Managers**: npm, bun, pnpm, yarn, brew, paru (Arch)
- **Custom Install Directory**: Support for XDG Base Directory Specification

### Versioning
- **Current**: 0.7.9 (from packages/opencode/package.json)
- **GitHub Releases**: Automated release process with download tracking
- **npm Package**: Published as `opencode-ai`