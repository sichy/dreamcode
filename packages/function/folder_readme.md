# packages/function - Serverless Functions

## 📦 Overview
This package contains **serverless function implementations** for OpenCode's cloud infrastructure. These functions run on Cloudflare Workers and handle specific background tasks, webhooks, and asynchronous operations that don't belong in the main API server.

## 🎯 Purpose
- **Background Processing**: Async tasks that don't block the main API
- **Webhook Handlers**: Process GitHub, Stripe, and other webhooks
- **Scheduled Jobs**: Cron-based maintenance and cleanup tasks
- **Event Processing**: Handle events from queues and streams
- **Data Processing**: Batch operations and transformations

## 📁 Directory Structure

### `/src` - Function Source Code
Individual function implementations.

#### Key Functions:
- **`sync.ts`** - Session synchronization across regions
- **`webhook.ts`** - Webhook event processing
- **`cleanup.ts`** - Scheduled cleanup tasks
- **`analytics.ts`** - Usage data aggregation
- **`notification.ts`** - Send notifications and alerts

## 🚀 Function Types

### 1. **HTTP Functions**
Handle incoming HTTP requests:
```typescript
export async function handler(request: Request) {
  // Process webhook
  const body = await request.json()
  // Return response
  return new Response('OK')
}
```

### 2. **Scheduled Functions**
Run on a schedule via Cron triggers:
```typescript
export async function scheduled(event: ScheduledEvent) {
  // Cleanup old sessions
  await cleanupExpiredSessions()
}
```

### 3. **Queue Functions**
Process messages from queues:
```typescript
export async function queue(batch: MessageBatch) {
  for (const message of batch.messages) {
    await processMessage(message)
  }
}
```

### 4. **Durable Object Functions**
Stateful serverless functions:
```typescript
export class SyncServer {
  constructor(state: DurableObjectState) {
    this.state = state
  }
  
  async fetch(request: Request) {
    // Handle sync operations
  }
}
```

## 🔧 Technologies Used
- **TypeScript** - Type-safe function development
- **Cloudflare Workers** - Serverless runtime
- **Wrangler** - Development and deployment
- **Miniflare** - Local testing

## 📝 Function Configuration

### Environment Variables:
```typescript
interface Env {
  DATABASE: D1Database
  BUCKET: R2Bucket
  KV: KVNamespace
  QUEUE: Queue
  SECRET_KEY: string
}
```

### Bindings:
- **D1** - SQL database access
- **R2** - Object storage
- **KV** - Key-value store
- **Queues** - Message queuing
- **Durable Objects** - Stateful compute

## 🎯 Common Use Cases

### Session Synchronization:
```typescript
// Sync sessions across regions
export async function syncSessions(env: Env) {
  const sessions = await env.DATABASE.prepare(
    "SELECT * FROM sessions WHERE updated > ?"
  ).bind(lastSync).all()
  
  await broadcastToRegions(sessions)
}
```

### Webhook Processing:
```typescript
// Handle GitHub webhooks
export async function githubWebhook(request: Request, env: Env) {
  const signature = request.headers.get('X-Hub-Signature')
  if (!verifySignature(signature, env.GITHUB_SECRET)) {
    return new Response('Unauthorized', { status: 401 })
  }
  
  const event = await request.json()
  await processGitHubEvent(event)
  
  return new Response('OK')
}
```

### Cleanup Jobs:
```typescript
// Clean expired data
export async function cleanup(env: Env) {
  // Delete old sessions
  await env.DATABASE.prepare(
    "DELETE FROM sessions WHERE expires < ?"
  ).bind(Date.now()).run()
  
  // Remove orphaned files
  await cleanupOrphanedFiles(env.BUCKET)
}
```

## 🚦 Deployment

### Deploy Functions:
```bash
# Deploy to production
wrangler deploy --env production

# Deploy specific function
wrangler deploy src/webhook.ts

# View logs
wrangler tail
```

### Local Development:
```bash
# Run locally
wrangler dev src/sync.ts

# Test with local data
miniflare src/webhook.ts
```

## 💡 Best Practices

### Performance:
- Keep functions lightweight
- Minimize cold start time
- Use efficient data structures
- Cache frequently accessed data

### Error Handling:
- Implement retry logic
- Use dead letter queues
- Log errors comprehensively
- Return appropriate status codes

### Security:
- Validate all inputs
- Verify webhook signatures
- Use environment secrets
- Implement rate limiting

## 🔍 Monitoring

### Metrics to Track:
- Function execution time
- Error rates
- Queue depth
- Memory usage
- CPU time

### Logging:
```typescript
console.log('INFO:', 'Processing started')
console.error('ERROR:', error.message)
console.warn('WARN:', 'Rate limit approaching')
```

## 📊 Performance Limits

### Cloudflare Workers:
- **CPU Time**: 10ms (free), 30s (paid)
- **Memory**: 128MB
- **Request Size**: 100MB
- **Subrequests**: 50 (free), 1000 (paid)

## 🔒 Security Considerations
- Validate webhook signatures
- Sanitize user inputs
- Use least privilege access
- Encrypt sensitive data
- Audit function logs

## 📈 Scaling Patterns

### Auto-scaling:
- Workers scale automatically
- No manual intervention needed
- Global distribution by default

### Rate Limiting:
```typescript
const limiter = new RateLimiter({
  requests: 100,
  window: 60 // seconds
})

if (!limiter.allow(clientId)) {
  return new Response('Too Many Requests', { status: 429 })
}
```

## 🎯 Testing

### Unit Tests:
```typescript
describe('Webhook Handler', () => {
  it('validates signatures', async () => {
    const response = await handler(mockRequest)
    expect(response.status).toBe(401)
  })
})
```

### Integration Tests:
- Test with Miniflare
- Mock external services
- Verify queue processing
- Check database operations