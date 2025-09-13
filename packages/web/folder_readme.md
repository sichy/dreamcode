# packages/web - Documentation Website

## 📦 Overview
This package contains the **official OpenCode documentation website** built with Astro and Starlight. It provides comprehensive documentation, guides, and API references for OpenCode users and developers.

## 🎯 Purpose
- **User Documentation**: Guides for getting started and using OpenCode
- **API Reference**: Detailed API documentation for developers
- **Feature Showcase**: Examples and use cases
- **Changelog & Updates**: Version history and release notes
- **Community Resources**: Links to GitHub, Discord, and other resources

## 📁 Directory Structure

### `/src`
Main source directory for the website.

#### Key Directories:
- **`/content`** - Documentation content in MDX format
  - `/docs` - Main documentation pages
  - `/config.ts` - Starlight configuration
  
- **`/components`** - Custom Astro/React components
  - Reusable UI components
  - Interactive examples
  - Code blocks with syntax highlighting
  
- **`/layouts`** - Page layouts and templates
  - Documentation layout
  - Landing page layout
  - Blog/changelog layout
  
- **`/styles`** - CSS and styling
  - Global styles
  - Component-specific styles
  - Theme customization

### `/public`
Static assets served directly:
- Images and icons
- Fonts
- Downloads
- Robots.txt and sitemap

## 📄 Documentation Structure

### Core Documentation:
1. **Getting Started**
   - Installation guide
   - Quick start tutorial
   - Basic usage

2. **CLI Reference**
   - Command documentation
   - Options and flags
   - Examples

3. **Configuration**
   - `opencode.json` reference
   - Provider setup
   - Agent configuration

4. **Features**
   - AI providers
   - Tool system
   - Session management
   - MCP integration

5. **API Documentation**
   - Server endpoints
   - SDK usage
   - WebSocket/SSE events

6. **Guides**
   - Best practices
   - Advanced usage
   - Troubleshooting

## 🚀 Main Features

1. **Starlight Documentation**
   - Automatic navigation generation
   - Search functionality
   - Dark/light theme toggle
   - Mobile responsive design

2. **MDX Support**
   - Markdown with JSX components
   - Interactive code examples
   - Embedded demos

3. **SEO Optimized**
   - Meta tags
   - Open Graph support
   - Sitemap generation
   - Schema.org markup

4. **Fast Performance**
   - Static site generation
   - Optimized images
   - Minimal JavaScript
   - Edge deployment ready

## 🔧 Technologies Used
- **Astro** - Static site generator
- **Starlight** - Documentation theme
- **MDX** - Markdown with components
- **TypeScript** - Type-safe development
- **Tailwind CSS** - Utility-first styling

## 📝 Key Files
- **`astro.config.mjs`** - Astro configuration
- **`package.json`** - Dependencies and scripts
- **`tsconfig.json`** - TypeScript configuration
- **`tailwind.config.mjs`** - Tailwind CSS configuration

## 🚦 Build & Development

```bash
# Install dependencies
bun install

# Start development server
bun run dev

# Build for production
bun run build

# Preview production build
bun run preview
```

## 📚 Content Guidelines

### Writing Documentation:
- Use clear, concise language
- Include code examples
- Add screenshots where helpful
- Link to related topics
- Keep content up-to-date

### MDX Components Available:
- `<Tabs>` - Tab containers
- `<Code>` - Syntax-highlighted code
- `<Callout>` - Important notes
- `<Card>` - Feature cards
- `<Video>` - Embedded videos

## 🌐 Deployment
The website is deployed to:
- **Production**: https://opencode.ai
- **Docs Subdomain**: https://docs.opencode.ai
- Hosted on Cloudflare Pages
- Automatic deployment on push to main

## 💡 Development Notes
- Content is in MDX format for rich interactivity
- Starlight provides automatic features like search and navigation
- Optimized for documentation sites
- Supports versioning for different releases
- Integrates with OpenCode's main site