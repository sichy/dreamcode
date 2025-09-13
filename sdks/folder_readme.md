# sdks - SDK Implementations

## 📦 Overview
This folder contains **SDK implementations for various platforms and editors** to integrate with OpenCode. These SDKs allow developers to use OpenCode from their favorite development environments, including VS Code, JetBrains IDEs, Vim, and more.

## 🎯 Purpose
- **IDE Integration**: Bring OpenCode to popular editors
- **Platform SDKs**: Native integrations for different platforms
- **Language Bindings**: Client libraries for various programming languages
- **Tool Integration**: Connect OpenCode with development tools
- **Extension Development**: Framework for building OpenCode extensions

## 📁 Directory Structure

### `/vscode` - Visual Studio Code Extension
Full-featured VS Code extension for OpenCode.

#### Features:
- **Command Palette Integration**: Access OpenCode commands
- **Inline Code Generation**: Generate code in editor
- **Chat Panel**: Interactive AI chat sidebar
- **Code Actions**: Quick fixes and refactoring
- **Status Bar**: Connection status and model selection
- **Keybindings**: Customizable shortcuts

#### Key Files:
- `extension.ts` - Main extension entry point
- `package.json` - Extension manifest
- `webview/` - Chat UI implementation

### `/jetbrains` - JetBrains IDEs Plugin
Plugin for IntelliJ IDEA, WebStorm, PyCharm, etc.

#### Features:
- **Tool Window**: Dedicated OpenCode panel
- **Code Completion**: AI-powered suggestions
- **Inspections**: AI-based code analysis
- **Intentions**: Quick AI actions
- **Project Integration**: Full project context

### `/vim` - Vim/Neovim Plugin
Vim plugin for terminal-based development.

#### Features:
- **Commands**: `:OpenCode` command set
- **Mappings**: Customizable key mappings
- **Floating Windows**: Response display
- **Async Operations**: Non-blocking AI calls
- **Integration**: Works with existing plugins

### `/emacs` - Emacs Package
Emacs Lisp package for OpenCode.

#### Features:
- **Major Mode**: `opencode-mode`
- **Minor Mode**: Global integration
- **Hydra Menus**: Quick action menus
- **Org Mode**: Integration with org-mode
- **REPL**: Interactive AI REPL

### `/sublime` - Sublime Text Package
Package for Sublime Text editor.

#### Features:
- **Command Palette**: OpenCode commands
- **Panel**: Output and chat panel
- **Build System**: AI-powered builds
- **Snippets**: AI snippet generation

## 🚀 Common SDK Features

### Core Functionality:
1. **Connection Management**
   - Connect to local/remote OpenCode server
   - Auto-discovery of running instances
   - Connection status indicators

2. **Chat Interface**
   - Interactive conversation with AI
   - Code block rendering
   - Markdown support
   - File references

3. **Code Generation**
   - Inline code completion
   - Function generation
   - Test generation
   - Documentation generation

4. **Code Analysis**
   - Error explanation
   - Code review
   - Performance suggestions
   - Security analysis

5. **File Operations**
   - Apply AI-suggested changes
   - Diff preview
   - Multi-file edits
   - Safe rollback

## 🔧 SDK Development

### Common Architecture:
```typescript
// SDK Base Class
class OpenCodeSDK {
  private client: OpencodeClient
  private connection: Connection
  
  async connect(options: ConnectionOptions) {
    this.client = new OpencodeClient(options)
    await this.client.connect()
  }
  
  async sendMessage(message: string) {
    return this.client.session.message({ content: message })
  }
  
  async applyChanges(changes: FileChange[]) {
    // Apply changes to files
  }
}
```

### Extension Points:
- **Commands**: Register custom commands
- **Providers**: Code completion, hover, etc.
- **Views**: Custom UI panels
- **Decorations**: Code highlighting
- **Languages**: Language-specific features

## 📝 Installation

### VS Code:
```bash
# Install from marketplace
code --install-extension opencode-ai

# Or build locally
cd sdks/vscode
npm install
npm run build
code --install-extension opencode-*.vsix
```

### JetBrains:
```bash
# Install from marketplace
# Settings -> Plugins -> Search "OpenCode"

# Or build locally
cd sdks/jetbrains
./gradlew buildPlugin
```

### Vim/Neovim:
```vim
" Using vim-plug
Plug 'opencode-ai/vim-opencode'

" Using packer.nvim
use 'opencode-ai/vim-opencode'
```

### Emacs:
```elisp
;; Using straight.el
(straight-use-package
 '(opencode :type git :host github :repo "opencode-ai/emacs-opencode"))

;; Using use-package
(use-package opencode
  :ensure t
  :config
  (opencode-mode))
```

## 🔌 Configuration

### Common Settings:
```json
{
  "opencode.serverUrl": "http://localhost:4096",
  "opencode.apiKey": "oc_xxx",
  "opencode.model": "anthropic/claude-3-5-sonnet",
  "opencode.autoConnect": true,
  "opencode.showInlineCompletions": true
}
```

### Keybindings:
```json
{
  "key": "ctrl+shift+o",
  "command": "opencode.chat",
  "key": "ctrl+shift+g",
  "command": "opencode.generate",
  "key": "ctrl+shift+r",
  "command": "opencode.review"
}
```

## 💡 SDK Guidelines

### Best Practices:
- **Non-blocking**: Never block the editor UI
- **Cancelable**: Allow users to cancel operations
- **Configurable**: Provide settings for customization
- **Accessible**: Follow accessibility guidelines
- **Performant**: Minimize resource usage

### User Experience:
- Clear status indicators
- Progress feedback
- Error messages with solutions
- Undo/redo support
- Keyboard navigation

### Integration:
- Respect editor conventions
- Work with existing plugins
- Support editor themes
- Handle multiple projects
- Preserve user settings

## 🚦 Publishing

### VS Code Marketplace:
```bash
vsce package
vsce publish
```

### JetBrains Marketplace:
```bash
./gradlew publishPlugin
```

### Package Registries:
- **NPM**: JavaScript/TypeScript SDKs
- **PyPI**: Python SDK
- **RubyGems**: Ruby SDK
- **NuGet**: .NET SDK
- **Cargo**: Rust SDK

## 🔍 Testing SDKs

### Unit Tests:
```bash
# VS Code
npm test

# JetBrains
./gradlew test

# Vim
vader test/*.vader
```

### Integration Tests:
- Test with real OpenCode server
- Verify all commands work
- Check error handling
- Performance testing

## 📊 SDK Metrics

### Usage Tracking:
- Installation count
- Active users
- Feature usage
- Error rates
- Performance metrics

### Feedback:
- User reviews
- GitHub issues
- Feature requests
- Bug reports

## 🔒 Security Considerations
- Secure API key storage
- Encrypted communication
- Input validation
- Safe file operations
- Permission checks

## 🌟 Future SDKs
- **Zed** - Modern collaborative editor
- **Helix** - Terminal-based modal editor
- **Nova** - Mac-native code editor
- **Lapce** - Lightning-fast editor
- **Xi** - Modern text editor