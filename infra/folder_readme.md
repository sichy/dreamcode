# infra - Infrastructure Configuration

## 📦 Overview  
This folder contains the **SST (Serverless Stack Toolkit) v3 infrastructure configuration** for OpenCode. It defines all cloud resources, deployment settings, and infrastructure as code for the production environment on Cloudflare.

## 🎯 Purpose
- **Infrastructure as Code**: Define all cloud resources in TypeScript
- **Multi-stage Deployments**: Manage dev, staging, and production environments
- **Resource Management**: Provision databases, storage, and compute
- **Secret Management**: Handle API keys and sensitive configuration
- **CI/CD Integration**: Automated deployments via GitHub Actions

## 📁 Directory Structure

### Root Configuration Files:
- **`sst.config.ts`** - Main SST configuration file
- **`sst-env.d.ts`** - TypeScript definitions for SST resources
- **`.sst/`** - SST build artifacts (gitignored)

### Resource Definitions:
Infrastructure resources are defined in the main config and cloud packages.

## 🚀 SST Configuration

### Main Config Structure:
```typescript
export default $config({
  app(input) {
    return {
      name: "opencode",
      home: "cloudflare",
      removal: input?.stage === "production" ? "retain" : "remove",
    }
  },
  async run() {
    // Resource definitions
  }
})
```

## 🏗️ Infrastructure Components

### 1. **Cloudflare Workers**
- API server deployment
- Edge computing
- WebSocket support
- Global distribution

### 2. **Storage Resources**
```typescript
// R2 Bucket for file storage
const bucket = new sst.cloudflare.Bucket("Bucket")

// D1 Database for structured data
const database = new sst.cloudflare.D1("Database")

// KV Namespace for caching
const kv = new sst.cloudflare.Kv("Cache")
```

### 3. **Durable Objects**
```typescript
// Real-time sync server
const sync = new sst.cloudflare.DurableObject("SyncServer", {
  handler: "cloud/app/sync.ts",
  class: "SyncServer",
})
```

### 4. **Secrets Management**
```typescript
// Stripe integration
const stripe = new sst.Secret("StripeSecretKey")

// GitHub App credentials
const githubAppId = new sst.Secret("GitHubAppId")
const githubPrivateKey = new sst.Secret("GitHubAppPrivateKey")

// Authentication
const openAuthSecret = new sst.Secret("OpenAuthSecret")
```

### 5. **Domain & DNS**
```typescript
// Custom domain setup
const domain = "opencode.ai"

// API subdomain
new sst.cloudflare.Worker("Api", {
  url: `api.${domain}`,
  // ...
})

// Docs subdomain
new sst.aws.StaticSite("Docs", {
  url: `docs.${domain}`,
  // ...
})
```

## 🌍 Environment Stages

### Development Stage:
- **Purpose**: Testing and development
- **URL**: `dev.api.opencode.ai`
- **Removal**: Resources deleted on `sst remove`
- **Debugging**: Enhanced logging enabled

### Production Stage:
- **Purpose**: Live production environment
- **URL**: `api.opencode.ai`
- **Removal**: Resources retained (protected)
- **Monitoring**: Full analytics and alerts

## 🔧 Technologies Used
- **SST v3** - Infrastructure framework
- **Cloudflare** - Cloud platform
- **TypeScript** - Type-safe IaC
- **Pulumi** - Under the hood engine

## 📝 Deployment Commands

### Basic Operations:
```bash
# Deploy to development
sst deploy --stage dev

# Deploy to production
sst deploy --stage production

# View deployment status
sst status --stage production

# Stream logs
sst logs --stage production

# Remove development stage
sst remove --stage dev
```

### Secret Management:
```bash
# Set a secret
sst secret set StripeSecretKey sk_live_xxx --stage production

# List secrets
sst secret list --stage production

# Remove a secret
sst secret remove GitHubToken --stage dev
```

## 🔌 Resource Linking

### Linking Resources to Code:
```typescript
// In application code
import { Resource } from "sst"

// Access bucket
const bucket = Resource.Bucket

// Access database
const db = Resource.Database

// Access secrets
const stripeKey = Resource.StripeSecretKey
```

## 💡 Configuration Best Practices

### Resource Naming:
- Use PascalCase for resource names
- Prefix with environment when needed
- Keep names descriptive but concise

### Stage Management:
- Never deploy directly to production
- Test in dev stage first
- Use feature branches for experiments
- Implement gradual rollouts

### Security:
- Store all secrets in SST Secrets
- Never commit sensitive data
- Use least privilege IAM roles
- Enable audit logging

## 🚦 CI/CD Integration

### GitHub Actions Workflow:
```yaml
- name: Deploy to Production
  run: |
    npx sst deploy --stage production
  env:
    CLOUDFLARE_API_TOKEN: ${{ secrets.CLOUDFLARE_API_TOKEN }}
```

## 📊 Monitoring & Observability

### Cloudflare Analytics:
- Request metrics
- Error rates
- Performance data
- Geographic distribution

### Custom Metrics:
- Session counts
- API usage by endpoint
- Model usage statistics
- Cost tracking

## 🔍 Debugging

### Local Development:
```bash
# Run locally with remote resources
sst dev --stage dev

# Console into resources
sst console --stage dev
```

### Production Debugging:
```bash
# Tail production logs
sst logs --stage production --tail

# Filter logs
sst logs --stage production --filter "ERROR"
```

## 🎯 Cost Optimization

### Cloudflare Pricing:
- **Workers**: 100,000 requests/day free
- **R2**: $0.015/GB storage
- **D1**: 5GB free
- **Durable Objects**: $0.15/million requests

### Optimization Tips:
- Use caching aggressively
- Implement request coalescing
- Optimize asset sizes
- Monitor usage patterns

## 🔒 Security Features
- Automatic SSL/TLS
- DDoS protection
- Web Application Firewall
- Rate limiting
- IP whitelisting options

## 📈 Scaling Considerations
- Auto-scaling built into Cloudflare
- Global edge deployment
- No cold starts with Workers
- Unlimited concurrent executions
- Geographic load balancing