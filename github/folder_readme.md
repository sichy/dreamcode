# github - GitHub Integration & Actions

## 📦 Overview
This folder contains **GitHub-specific integrations for OpenCode**, including GitHub Actions workflows, the GitHub App implementation, and automation tools that allow OpenCode to work directly within GitHub repositories through comments and pull requests.

## 🎯 Purpose
- **GitHub Actions**: Automated workflows for CI/CD
- **GitHub App**: Bot that responds to PR comments
- **Issue/PR Automation**: AI-powered code reviews and assistance
- **Release Management**: Automated releases and changelogs
- **Integration Tools**: GitHub API interactions

## 📁 Directory Structure

### `/action` - GitHub Action Implementation
The main GitHub Action that users can add to their repositories.

#### Components:
- **`action.yml`** - Action definition and inputs
- **`index.js`** - Action entry point
- **`runner.ts`** - Core logic for running OpenCode in Actions

### `/app` - GitHub App
The OpenCode GitHub App that responds to repository events.

#### Features:
- **Comment Handler** - Responds to `/oc` and `/opencode` commands
- **PR Review** - Automated code review assistance
- **Issue Triage** - Help with issue management
- **Commit Suggestions** - AI-powered commit messages

### `/workflows` - CI/CD Workflows
GitHub Actions workflows for the OpenCode project itself.

#### Workflows:
- **`ci.yml`** - Continuous integration (tests, linting)
- **`release.yml`** - Automated releases
- **`publish.yml`** - NPM and binary publishing
- **`stats.yml`** - Download statistics tracking

## 🚀 Main Features

### 1. **Comment Commands**
Users can invoke OpenCode in PRs/issues:
```
/opencode review this PR
/oc fix the failing tests
/opencode explain this change
```

### 2. **Automated Code Review**
- Analyzes PR changes
- Suggests improvements
- Identifies potential bugs
- Checks for best practices

### 3. **Issue Management**
- Auto-labeling issues
- Suggesting solutions
- Creating fix PRs
- Triaging bug reports

### 4. **CI/CD Integration**
- Run OpenCode in GitHub Actions
- Automated testing with AI
- Code quality checks
- Documentation generation

## 🔧 GitHub App Setup

### Installation:
1. Install the OpenCode GitHub App
2. Grant repository permissions
3. Configure in `.github/opencode.yml`

### Permissions Required:
- **Read**: Code, issues, pull requests
- **Write**: Comments, pull requests
- **Admin**: Webhooks (for app only)

### Configuration:
```yaml
# .github/opencode.yml
version: 1
settings:
  auto_review: true
  commands_enabled: true
  model: anthropic/claude-3-5-sonnet
  
rules:
  - pattern: "*.test.js"
    action: "generate tests"
  - pattern: "*.md"
    action: "check documentation"
```

## 📝 GitHub Action Usage

### Basic Workflow:
```yaml
name: OpenCode Review
on:
  pull_request:
    types: [opened, synchronize]

jobs:
  review:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: sst/opencode@v1
        with:
          command: "review"
          api_key: ${{ secrets.OPENCODE_API_KEY }}
```

### Available Commands:
- `review` - Review PR changes
- `test` - Generate/run tests
- `document` - Update documentation
- `fix` - Auto-fix issues
- `explain` - Explain code changes

## 🔌 API Integration

### Webhook Events:
- `issue_comment.created`
- `pull_request.opened`
- `pull_request.synchronize`
- `issues.opened`
- `push` (on main branch)

### GitHub API Usage:
- GraphQL for efficient queries
- REST API for mutations
- Octokit SDK integration
- Rate limit handling

## 🔧 Technologies Used
- **Node.js** - Action runtime
- **TypeScript** - Type-safe development
- **Octokit** - GitHub API client
- **Probot** - GitHub App framework
- **Actions Toolkit** - GitHub Actions utilities

## 💡 Development Notes

### Testing GitHub Actions:
```bash
# Local testing with act
act -j review

# Test with specific event
act pull_request -e event.json
```

### Debugging:
- Enable debug logging: `ACTIONS_STEP_DEBUG=true`
- Check webhook deliveries in GitHub settings
- Use ngrok for local webhook testing

### Best Practices:
- Cache dependencies in workflows
- Use Action secrets for sensitive data
- Implement rate limit handling
- Provide clear error messages

## 🚦 Publishing

### GitHub Action:
```bash
# Tag new version
git tag v1.0.0
git push origin v1.0.0

# Update major version tag
git tag -f v1
git push -f origin v1
```

### GitHub App:
- Deployed automatically via SST
- Webhook URL points to Cloudflare Worker
- Private key stored in secrets

## 🔒 Security Considerations
- Never log sensitive information
- Validate webhook signatures
- Use minimal required permissions
- Implement rate limiting
- Sanitize user inputs

## 📊 Analytics & Monitoring
- GitHub Insights for app usage
- Action run statistics
- Error tracking with Sentry
- Performance metrics logging

## 🌟 Example Use Cases

### Automated PR Review:
```
/opencode review focusing on security
```

### Fix Failing Tests:
```
/oc fix the test failures in api.test.js
```

### Generate Documentation:
```
/opencode document the new API endpoints
```

### Explain Complex Changes:
```
/opencode explain what this refactor does
```