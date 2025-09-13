# cloud - Cloud Infrastructure Packages

## 📦 Overview
This folder contains **cloud infrastructure code for OpenCode's production deployment**. It includes the Cloudflare Workers application, serverless functions, and infrastructure resources managed by SST (Serverless Stack Toolkit).

## 🎯 Purpose
- **Cloud Deployment**: Production infrastructure on Cloudflare
- **API Worker**: Main application running on Cloudflare Workers
- **Durable Objects**: Real-time synchronization and state management
- **Edge Functions**: Serverless compute at the edge
- **Resource Management**: Database, storage, and other cloud resources

## 📁 Directory Structure

### `/app` - Main Application Worker
The primary Cloudflare Worker serving the OpenCode API.

#### Key Components:
- **API Routes** - REST endpoints served from the edge
- **WebSocket Handler** - Real-time communication
- **SSE Streaming** - Server-sent events for AI responses
- **Authentication** - Edge-based auth handling
- **Rate Limiting** - Request throttling

### `/core` - Core Infrastructure Components
Shared infrastructure code and utilities.

#### Includes:
- **Database Abstractions** - D1 database interfaces
- **Storage Handlers** - R2 bucket operations
- **Cache Management** - Cloudflare cache APIs
- **Security Middleware** - CORS, CSP, auth checks
- **Error Handling** - Centralized error management

### `/function` - Serverless Functions
Individual serverless functions for specific tasks.

#### Functions:
- **Session Sync** - Cross-region session synchronization
- **File Processing** - Background file operations
- **Webhook Handlers** - GitHub and other integrations
- **Scheduled Tasks** - Cron jobs and maintenance

### `/resource` - Infrastructure Resources
Resource definitions for SST deployment.

#### Resources:
- **D1 Database** - SQLite at the edge
- **R2 Buckets** - Object storage for files
- **KV Namespaces** - Key-value storage
- **Durable Objects** - Stateful serverless compute
- **Queues** - Message queuing system
- **Analytics** - Usage tracking

### `/scripts` - Deployment Scripts
Build and deployment automation.

#### Scripts:
- **Deploy** - Production deployment
- **Migrate** - Database migrations
- **Seed** - Data seeding
- **Monitor** - Health checks
- **Rollback** - Version rollback

## 🚀 Key Features

### 1. **Edge Computing**
- Global distribution via Cloudflare network
- Sub-10ms latency worldwide
- Automatic failover and redundancy
- DDoS protection built-in

### 2. **Durable Objects**
- **SyncServer** - Real-time session synchronization
- **RateLimiter** - Distributed rate limiting
- **SessionStore** - Persistent session state
- **WebSocketManager** - Connection management

### 3. **Storage Solutions**
- **R2** - File and asset storage
- **D1** - Relational database
- **KV** - Fast key-value cache
- **Cache API** - Edge caching

### 4. **Security Features**
- Zero-trust security model
- Encrypted data at rest and in transit
- API key management
- IP-based access controls

## 🔧 Technologies Used
- **Cloudflare Workers** - Edge compute platform
- **TypeScript** - Type-safe development
- **SST v3** - Infrastructure as code
- **Wrangler** - Cloudflare CLI tool
- **Miniflare** - Local development

## 📝 Configuration

### Environment Variables:
```env
# Production secrets (stored in SST)
STRIPE_SECRET_KEY
GITHUB_APP_ID
GITHUB_APP_PRIVATE_KEY
OPENAUTH_SECRET

# Cloudflare resources
DATABASE_ID
BUCKET_NAME
KV_NAMESPACE
```

### SST Configuration:
- Defined in `sst.config.ts`
- Stage-based deployments (dev, production)
- Automatic resource provisioning

## 🌐 Deployment Architecture

```
┌─────────────┐     ┌──────────────┐     ┌──────────────┐
│   Client    │────▶│  Cloudflare  │────▶│   Worker     │
│  (TUI/CLI)  │     │   Network    │     │   (API)      │
└─────────────┘     └──────────────┘     └──────────────┘
                            │                     │
                            ▼                     ▼
                    ┌──────────────┐     ┌──────────────┐
                    │   R2 Bucket  │     │  D1 Database │
                    │   (Files)    │     │  (Sessions)  │
                    └──────────────┘     └──────────────┘
```

## 🚦 Deployment Commands

```bash
# Deploy to production
sst deploy --stage production

# Deploy to development
sst deploy --stage dev

# View logs
sst logs --stage production

# Remove deployment
sst remove --stage dev
```

## 💡 Development Notes

### Local Development:
- Use `wrangler dev` for local worker testing
- Miniflare simulates Cloudflare environment
- SST provides local resource emulation

### Best Practices:
- Keep workers lightweight (<1MB)
- Use Durable Objects for stateful logic
- Leverage edge caching aggressively
- Monitor cold starts and performance

### Limitations:
- 10ms CPU time limit per request
- 128MB memory limit
- No filesystem access
- Limited Node.js API compatibility

## 🔍 Monitoring & Analytics
- Cloudflare Analytics dashboard
- Custom metrics via Workers Analytics
- Error tracking with Sentry integration
- Performance monitoring with Web Vitals

## 🔒 Security Considerations
- All traffic through Cloudflare proxy
- Automatic SSL/TLS termination
- Web Application Firewall (WAF)
- Rate limiting at edge
- Bot protection enabled