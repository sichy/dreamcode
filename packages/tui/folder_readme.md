# packages/tui - Terminal User Interface

## 📦 Overview
This is the **Go-based Terminal User Interface (TUI)** for OpenCode. It provides a rich, interactive terminal experience for chatting with AI, viewing code changes, and managing sessions. The TUI communicates with the TypeScript server via HTTP API calls.

## 🎯 Purpose
- **Interactive Terminal Interface**: Beautiful, responsive terminal UI for OpenCode
- **Client-Server Architecture**: Connects to the OpenCode server via HTTP
- **Rich Text Rendering**: Markdown rendering, syntax highlighting, and formatted output
- **Keyboard-Driven**: Full keyboard navigation and shortcuts
- **Real-time Updates**: Live streaming of AI responses and file changes

## 📁 Directory Structure

### `/cmd/opencode`
- **`main.go`** - Entry point for the TUI application
- Initializes the Bubble Tea app and starts the UI

### `/input`
- Terminal input handling and key mapping
- Mouse support implementation
- Color definitions and theming
- Cross-platform input drivers (Windows, Unix)

### `/internal`
- Core TUI implementation using Bubble Tea framework

#### Key Components:
- **`/app`** - Main application state and update logic
  - Session management
  - Model selection
  - Agent switching
  
- **`/api`** - HTTP client for server communication
  - OpenCode SDK integration
  - API request/response handling
  
- **`/ui`** - UI components and views
  - Chat interface
  - Dialog boxes (help, sessions, models)
  - Input field with multi-line support
  - Status bar and indicators
  
- **`/state`** - Application state management
  - Current session tracking
  - Message history
  - User preferences
  
- **`/keybinds`** - Keyboard shortcut handling
  - Configurable key mappings
  - Leader key combinations
  - Vim-style navigation

## 🚀 Main Features

1. **Chat Interface**
   - Real-time streaming AI responses
   - Markdown rendering with syntax highlighting
   - Code block display with language detection
   - File diff visualization

2. **Session Management**
   - Create, list, and switch sessions
   - Session history navigation
   - Export sessions to editor
   - Share/unshare sessions

3. **Model Selection**
   - Quick model switching (F2)
   - Provider browsing
   - Recent models list
   - Model capabilities display

4. **Agent System**
   - Tab to cycle through agents
   - Agent-specific prompts
   - Mode indicators (build, plan, general)

5. **File Operations**
   - View file changes in real-time
   - Approve/reject file edits
   - Syntax-highlighted code display
   - Diff view for modifications

6. **Keyboard Navigation**
   - Vim-inspired keybindings
   - Leader key system (Ctrl+X by default)
   - Customizable shortcuts
   - Help dialog with key reference

## 🔧 Technologies Used
- **Go** - Primary language
- **Bubble Tea** - Terminal UI framework
- **Lipgloss** - Styling and layout
- **Glamour** - Markdown rendering
- **Bubbles** - Pre-built UI components
- **OpenCode SDK** - Server communication

## 📝 Key Files
- **`go.mod`** - Go module dependencies
- **`go.sum`** - Dependency checksums
- **`.goreleaser.yml`** - Release configuration
- **`.gitignore`** - Git ignore patterns

## 🎨 UI Components

### Main View Areas:
1. **Header** - Session info, model, agent
2. **Messages Area** - Chat history with scrolling
3. **Input Field** - Multi-line text input
4. **Status Bar** - Connection status, hints

### Dialogs:
- **Help** (Ctrl+X, H) - Keybinding reference
- **Sessions** (Ctrl+X, L) - Session list and management
- **Models** (Ctrl+X, M) - Model selection
- **Themes** (Ctrl+X, T) - Theme switcher

## 🔌 Server Communication
The TUI communicates with the OpenCode server via:
- HTTP REST API for commands
- Server-Sent Events for streaming
- Polling for state updates

## 💡 Development Notes
- Built with Bubble Tea's Elm-inspired architecture
- Uses message passing for state updates
- Implements responsive layout system
- Supports both mouse and keyboard input
- Terminal capability detection for colors/features

## 🚦 Build Instructions
```bash
# Build the TUI
go build ./cmd/opencode/

# Run the TUI
./opencode

# Or use the built binary
packages/tui/opencode
```

## 🎯 Design Principles
- **Performance**: Efficient rendering and minimal redraws
- **Responsiveness**: Smooth scrolling and instant feedback
- **Accessibility**: Clear visual hierarchy and keyboard navigation
- **Cross-platform**: Works on macOS, Linux, and Windows terminals