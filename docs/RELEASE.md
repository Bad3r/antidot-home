# Release Process

This document describes the release process for antidot-home.

## Automatic Release

Releases are automatically triggered when a version tag is pushed to the repository. The release workflow will:

1. Generate and update CHANGELOG.md
2. Build binaries for multiple platforms
3. Create a GitHub release with artifacts
4. Update the Homebrew tap (if configured)

### Creating a Release

To create a new release, follow these steps:

```bash
# 1. Make sure you're on the main branch with latest changes
git checkout main
git pull origin main

# 2. Create and push a version tag
git tag -s v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

The release workflow will automatically:
- Generate the changelog from commit history
- Commit the updated CHANGELOG.md to main
- Build release artifacts
- Create the GitHub release

## Manual Changelog Update

You can manually update the changelog without creating a release:

1. Go to Actions → "Update Changelog"
2. Click "Run workflow"
3. Optionally specify a version tag
4. The workflow will create a PR with the updated changelog

## Versioning

We follow [Semantic Versioning](https://semver.org/):
- MAJOR version for incompatible API changes
- MINOR version for backwards-compatible functionality additions
- PATCH version for backwards-compatible bug fixes

## Commit Message Format

We use [Conventional Commits](https://www.conventionalcommits.org/) format:
- `feat:` New features (bumps MINOR version)
- `fix:` Bug fixes (bumps PATCH version)
- `feat!:` or `BREAKING CHANGE:` Breaking changes (bumps MAJOR version)
- `chore:`, `docs:`, `test:`, `refactor:` Don't trigger releases