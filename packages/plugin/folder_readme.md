# packages/plugin - Plugin System

## 📦 Overview
This package defines the **plugin architecture for OpenCode**, allowing third-party developers to extend OpenCode's functionality with custom tools, providers, and integrations. Plugins can add new AI capabilities, tool functions, or modify behavior.

## 🎯 Purpose
- **Plugin API**: Define the interface for OpenCode plugins
- **Plugin Loading**: Dynamic plugin discovery and loading
- **Tool Registration**: Allow plugins to register custom tools
- **Provider Extensions**: Add custom AI provider support
- **Lifecycle Management**: Handle plugin initialization and cleanup

## 📁 Directory Structure

### `/src`
Core plugin system implementation.

#### Key Files:
- **`index.ts`** - Main plugin API exports
- **`types.ts`** - TypeScript type definitions for plugins
- **`loader.ts`** - Plugin loading and discovery logic
- **`registry.ts`** - Plugin registration system

## 🚀 Plugin Capabilities

### What Plugins Can Do:
1. **Register Custom Tools**
   - Add new tool functions for AI to use
   - Define tool schemas and parameters
   - Handle tool execution

2. **Add Provider Support**
   - Integrate new AI providers
   - Custom model configurations
   - Authentication handling

3. **Extend Commands**
   - Add new CLI commands
   - Modify existing command behavior
   - Custom command templates

4. **Hook into Events**
   - Listen to session events
   - File change notifications
   - Message processing hooks

## 🔧 Plugin Interface

```typescript
interface OpencodePlugin {
  name: string
  version: string
  description?: string
  
  // Lifecycle hooks
  onLoad?: () => Promise<void>
  onUnload?: () => Promise<void>
  
  // Tool registration
  tools?: Tool[]
  
  // Provider registration  
  providers?: Provider[]
  
  // Command extensions
  commands?: Command[]
  
  // Event handlers
  handlers?: {
    onMessage?: (message: Message) => void
    onFileChange?: (file: File) => void
    onSessionStart?: (session: Session) => void
  }
}
```

## 📝 Creating a Plugin

### Basic Plugin Structure:
```typescript
export default {
  name: 'my-plugin',
  version: '1.0.0',
  description: 'My custom OpenCode plugin',
  
  async onLoad() {
    console.log('Plugin loaded!')
  },
  
  tools: [{
    id: 'my-tool',
    description: 'My custom tool',
    execute: async (params) => {
      // Tool implementation
      return { result: 'success' }
    }
  }]
}
```

## 🔌 Plugin Loading

### Discovery Methods:
1. **Local Plugins** - From `~/.opencode/plugins/`
2. **NPM Packages** - Packages prefixed with `opencode-plugin-`
3. **Configuration** - Listed in `opencode.json`
4. **Runtime** - Dynamically loaded via API

### Load Order:
1. Core plugins (built-in)
2. User plugins (from config)
3. Local plugins (from directory)
4. Runtime plugins (dynamic)

## 🎨 Plugin Types

### Tool Plugins
Add new capabilities for AI:
- Database queries
- API integrations
- Custom file operations
- External service connections

### Provider Plugins
Support new AI models:
- Local model support
- Custom API endpoints
- Specialized model handlers

### Formatter Plugins
Code formatting tools:
- Language-specific formatters
- Custom linting rules
- Style enforcement

## 🔧 Technologies Used
- **TypeScript** - Type-safe plugin development
- **Zod** - Schema validation for plugin configs
- **Dynamic Import** - Runtime plugin loading

## 📝 Key Files
- **`package.json`** - Package metadata
- **`tsconfig.json`** - TypeScript configuration
- **`README.md`** - Plugin development guide

## 💡 Best Practices

### Plugin Development:
- Keep plugins focused on single functionality
- Handle errors gracefully
- Provide clear documentation
- Use semantic versioning
- Test across OpenCode versions

### Security Considerations:
- Validate all inputs
- Sandbox execution where possible
- Request minimal permissions
- Avoid storing sensitive data
- Use secure communication

## 🚦 Publishing a Plugin

```bash
# Build the plugin
bun run build

# Publish to NPM
npm publish

# Users install with:
bun add opencode-plugin-yourname
```

## 📚 Example Plugins

### Database Plugin:
```typescript
export default {
  name: 'opencode-plugin-database',
  tools: [{
    id: 'query-database',
    execute: async ({ query }) => {
      const result = await db.query(query)
      return { rows: result }
    }
  }]
}
```

### Custom Provider:
```typescript
export default {
  name: 'opencode-plugin-local-llm',
  providers: [{
    id: 'local-llm',
    models: ['llama2', 'mistral'],
    createModel: (modelId) => {
      return new LocalLLMModel(modelId)
    }
  }]
}
```

## 💡 Development Notes
- Plugins run in the same process as OpenCode
- Access to full OpenCode API via imports
- Can modify behavior through event hooks
- Should be backwards compatible
- Respect user privacy and preferences