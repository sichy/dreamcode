# packages/identity - Authentication & Identity

## 📦 Overview
This package handles **authentication and identity management** for OpenCode's cloud services. It provides user authentication, API key management, OAuth integration, and access control for the OpenCode platform using OpenAuth.

## 🎯 Purpose
- **User Authentication**: Sign up, login, and session management
- **OAuth Integration**: Social login providers (GitHub, Google, etc.)
- **API Key Management**: Generate and validate API keys
- **Access Control**: Role-based permissions and authorization
- **Identity Federation**: Single sign-on across OpenCode services

## 📁 Directory Structure

### `/src` - Identity Source Code
Core authentication implementation.

#### Key Components:
- **`auth.ts`** - Authentication logic and flows
- **`provider.ts`** - OAuth provider integrations
- **`session.ts`** - Session management
- **`token.ts`** - JWT token generation/validation
- **`user.ts`** - User management and profiles
- **`apikey.ts`** - API key generation and validation

## 🚀 Authentication Features

### 1. **User Authentication**
```typescript
// Sign up new user
export async function signUp(email: string, password: string) {
  const hashedPassword = await hashPassword(password)
  const user = await createUser({ email, password: hashedPassword })
  const token = generateToken(user)
  return { user, token }
}

// Login existing user
export async function login(email: string, password: string) {
  const user = await findUserByEmail(email)
  if (!await verifyPassword(password, user.password)) {
    throw new Error('Invalid credentials')
  }
  return generateToken(user)
}
```

### 2. **OAuth Providers**
Supported providers:
- **GitHub** - Developer authentication
- **Google** - General users
- **Microsoft** - Enterprise users
- **GitLab** - Alternative git platform

### 3. **API Key System**
```typescript
// Generate API key
export async function generateApiKey(userId: string) {
  const key = `oc_${generateRandomString(32)}`
  await saveApiKey(userId, key)
  return key
}

// Validate API key
export async function validateApiKey(key: string) {
  const apiKey = await findApiKey(key)
  if (!apiKey || apiKey.expired) {
    throw new Error('Invalid API key')
  }
  return apiKey.userId
}
```

### 4. **Session Management**
```typescript
// Create session
export async function createSession(userId: string) {
  const sessionId = generateSessionId()
  const expiresAt = Date.now() + SESSION_DURATION
  
  await saveSession({
    id: sessionId,
    userId,
    expiresAt
  })
  
  return sessionId
}
```

## 🔧 Technologies Used
- **OpenAuth** - Authentication framework
- **JWT** - JSON Web Tokens
- **bcrypt** - Password hashing
- **OAuth 2.0** - Social login standard
- **Cloudflare D1** - User database

## 📝 Configuration

### Environment Variables:
```env
# JWT Configuration
JWT_SECRET=your-secret-key
JWT_EXPIRY=7d

# OAuth Providers
GITHUB_CLIENT_ID=xxx
GITHUB_CLIENT_SECRET=xxx
GOOGLE_CLIENT_ID=xxx
GOOGLE_CLIENT_SECRET=xxx

# Database
DATABASE_URL=xxx

# OpenAuth
OPENAUTH_SECRET=xxx
```

### OpenAuth Setup:
```typescript
import { OpenAuth } from '@openauthjs/openauth'

export const auth = new OpenAuth({
  secret: process.env.OPENAUTH_SECRET,
  providers: [
    GitHubProvider({
      clientId: process.env.GITHUB_CLIENT_ID,
      clientSecret: process.env.GITHUB_CLIENT_SECRET,
    }),
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET,
    })
  ]
})
```

## 🔒 Security Features

### Password Security:
- **Hashing**: bcrypt with salt rounds
- **Complexity**: Minimum requirements enforced
- **Reset**: Secure token-based reset flow
- **History**: Prevent password reuse

### Token Security:
- **JWT Signing**: RS256 algorithm
- **Expiration**: Configurable TTL
- **Refresh Tokens**: Secure rotation
- **Revocation**: Blacklist support

### Session Security:
- **CSRF Protection**: Token validation
- **Session Fixation**: Prevention measures
- **Secure Cookies**: HttpOnly, Secure, SameSite
- **Idle Timeout**: Automatic expiration

## 🔌 API Endpoints

### Authentication:
- `POST /auth/signup` - Create account
- `POST /auth/login` - Login user
- `POST /auth/logout` - Logout user
- `POST /auth/refresh` - Refresh token
- `POST /auth/forgot` - Password reset

### OAuth:
- `GET /auth/github` - GitHub OAuth
- `GET /auth/google` - Google OAuth
- `GET /auth/callback` - OAuth callback

### API Keys:
- `POST /api/keys` - Generate key
- `GET /api/keys` - List keys
- `DELETE /api/keys/:id` - Revoke key

## 💡 Usage Examples

### Frontend Integration:
```typescript
// Login flow
const response = await fetch('/auth/login', {
  method: 'POST',
  body: JSON.stringify({ email, password })
})

const { token } = await response.json()
localStorage.setItem('token', token)
```

### API Authentication:
```typescript
// Authenticate API request
const response = await fetch('/api/session', {
  headers: {
    'Authorization': `Bearer ${token}`,
    'X-API-Key': apiKey
  }
})
```

## 🎯 Access Control

### Role-Based Access:
```typescript
enum Role {
  USER = 'user',
  ADMIN = 'admin',
  DEVELOPER = 'developer'
}

// Check permissions
export function hasPermission(user: User, resource: string) {
  return permissions[user.role].includes(resource)
}
```

### Rate Limiting:
```typescript
// API key rate limits
const limits = {
  free: { requests: 1000, window: 3600 },
  pro: { requests: 10000, window: 3600 },
  enterprise: { requests: -1, window: 0 } // Unlimited
}
```

## 📊 User Management

### User Profile:
```typescript
interface User {
  id: string
  email: string
  username?: string
  avatar?: string
  role: Role
  createdAt: Date
  lastLogin: Date
  emailVerified: boolean
}
```

### Account Operations:
- Profile updates
- Email verification
- Account deletion
- Data export (GDPR)

## 🔍 Monitoring & Analytics

### Track Metrics:
- Login attempts
- Failed authentications
- OAuth provider usage
- API key usage
- Session duration

### Security Events:
- Suspicious login patterns
- Brute force attempts
- Account takeover attempts
- API key abuse

## 🚦 Development

### Testing:
```bash
# Run tests
bun test

# Test OAuth flow
bun test:oauth

# Test API keys
bun test:apikeys
```

### Local Development:
```bash
# Start auth server
bun dev

# Use test OAuth app
OAUTH_REDIRECT_URL=http://localhost:3000/callback
```

## 💡 Best Practices
- Never store plaintext passwords
- Use secure random generators
- Implement proper session management
- Rate limit authentication endpoints
- Log security events
- Regular security audits