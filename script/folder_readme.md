# script - Build & Utility Scripts

## 📦 Overview
This folder contains **build scripts, automation tools, and utility scripts** for the OpenCode project. These TypeScript scripts handle everything from code formatting to release management, statistics collection, and development workflow automation.

## 🎯 Purpose
- **Build Automation**: Compile and package the project
- **Code Quality**: Formatting and linting scripts
- **Release Management**: Version bumping and publishing
- **Statistics**: Download and usage tracking
- **Development Tools**: Helper scripts for developers

## 📁 Script Files

### Core Scripts:

#### `format.ts` - Code Formatting
Ensures consistent code style across the project.
- Runs Prettier on all TypeScript/JavaScript files
- Formats JSON, YAML, and Markdown files
- Checks for formatting issues in CI
- Auto-fixes formatting problems

#### `stats.ts` - Statistics Collection
Tracks project metrics and usage.
- Collects GitHub release download statistics
- NPM package download counts
- Updates STATS.md with latest numbers
- Generates usage reports

#### `hooks` - Git Hooks Setup
Configures Git hooks for the repository.
- Pre-commit formatting checks
- Commit message validation
- Pre-push test execution
- Branch protection rules

#### `build.ts` - Build Orchestration
Manages the build process for all packages.
- Compiles TypeScript code
- Bundles for different targets
- Generates distributable packages
- Creates release artifacts

#### `release.ts` - Release Automation
Handles version management and publishing.
- Bumps version numbers
- Generates changelogs
- Creates Git tags
- Publishes to NPM
- Creates GitHub releases

## 🚀 Usage Examples

### Running Scripts:
```bash
# Format all code
bun run script/format.ts

# Check formatting (CI mode)
bun run script/format.ts --check

# Collect statistics
bun run script/stats.ts

# Build all packages
bun run script/build.ts

# Create a release
bun run script/release.ts --version minor
```

### Script Options:

#### Format Script:
- `--check` - Check formatting without fixing
- `--staged` - Only format staged files
- `--since` - Format files changed since commit

#### Stats Script:
- `--output` - Output file path
- `--format` - Output format (json, markdown)
- `--period` - Time period for stats

#### Release Script:
- `--version` - Version type (major, minor, patch)
- `--dry-run` - Preview without publishing
- `--skip-tests` - Skip test execution
- `--npm-tag` - NPM distribution tag

## 🔧 Technologies Used
- **TypeScript** - Type-safe scripting
- **Bun** - Script runtime and tooling
- **Prettier** - Code formatting
- **Octokit** - GitHub API interaction
- **Semver** - Version management

## 📝 Script Configuration

### Environment Variables:
```env
# GitHub Access
GITHUB_TOKEN=ghp_xxxxx

# NPM Publishing
NPM_TOKEN=npm_xxxxx

# Release Settings
RELEASE_BRANCH=main
SKIP_CHANGELOG=false
```

### Configuration Files:
- `.prettierrc` - Formatting rules
- `.gitignore` - Files to exclude
- `release.config.js` - Release settings

## 🔄 Automation Workflows

### Pre-commit:
1. Format staged files
2. Run type checking
3. Execute unit tests
4. Validate commit message

### Pre-release:
1. Run all tests
2. Build all packages
3. Generate changelog
4. Update version numbers
5. Create Git tag
6. Publish packages

### Post-release:
1. Create GitHub release
2. Update documentation
3. Notify Discord/Slack
4. Collect initial stats

## 💡 Development Scripts

### Utility Functions:
```typescript
// Common utilities used across scripts
async function runCommand(cmd: string)
async function updateFile(path: string, updater: Function)
async function collectStats()
async function publishPackage(name: string)
```

### Helper Scripts:
- **`clean.ts`** - Remove build artifacts
- **`link.ts`** - Link local packages
- **`doctor.ts`** - Check dev environment
- **`migrate.ts`** - Run migrations

## 🎯 Best Practices

### Script Writing:
- Use TypeScript for type safety
- Handle errors gracefully
- Provide clear console output
- Support dry-run mode
- Add progress indicators

### Performance:
- Parallelize where possible
- Cache expensive operations
- Use streaming for large files
- Minimize external API calls

## 🔍 Debugging Scripts

### Debug Mode:
```bash
# Enable verbose logging
DEBUG=* bun run script/build.ts

# Inspect script execution
bun --inspect script/release.ts
```

### Common Issues:
- **Permission Errors**: Check file permissions
- **API Rate Limits**: Use authentication tokens
- **Version Conflicts**: Clear node_modules
- **Build Failures**: Check TypeScript errors

## 📊 Script Metrics

### Performance Benchmarks:
- Format check: ~2 seconds
- Full build: ~30 seconds
- Test suite: ~45 seconds
- Release process: ~5 minutes

### Usage Statistics:
- Most used: `format.ts` (daily)
- CI critical: `build.ts`, `format.ts --check`
- Release cycle: `release.ts` (weekly)

## 🔒 Security Considerations
- Never commit tokens or secrets
- Validate external inputs
- Use minimal permissions
- Audit dependencies regularly
- Sign releases with GPG

## 🚦 CI Integration

These scripts are used in GitHub Actions:
```yaml
- name: Format Check
  run: bun run script/format.ts --check

- name: Build
  run: bun run script/build.ts

- name: Collect Stats
  run: bun run script/stats.ts
```