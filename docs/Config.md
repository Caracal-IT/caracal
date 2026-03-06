# Configuration Guide for Caracal

This document describes how configuration works in the caracal application.

## Configuration Loading Strategy

The caracal application uses a layered configuration approach:

1. **Base Configuration (`config.yml`)** - Loaded first as defaults
2. **Environment-Specific Configuration** - Loaded second and overrides base values
3. **Environment Variables** - Can override any YAML configuration at runtime

### Priority Order (Highest to Lowest)
1. Environment variables (`CARACAL_*`)
2. Environment-specific YAML file (e.g., `config.dev.yml`, `config.staging.yml`, `config.prod.yml`)
3. Base YAML file (`config.yml`)
4. Application hardcoded defaults

## Configuration Files

### `config.yml`
The base configuration file with default settings.
- **Use case:** Base defaults for all environments
- **Loaded:** Always loaded first
- **Overrides:** Can be overridden by environment-specific configs
- **Contains:** Application metadata and settings

**Example:**
```yaml
app:
  name: caracal
  version: 0.1.0
  description: "Caracal - A high-performance Go application"
  environment: development
```

### `config.dev.yml`
Development environment configuration.
- **Use case:** Local development with sensible overrides
- **Environment:** Set when `CARACAL_ENV=dev`
- **Overrides:** Base config values for development
- **Contains:** Application config specific to development

**Example:**
```yaml
app:
  name: caracal
  version: 0.1.0
  description: "Caracal - A high-performance Go application"
  environment: development
```

### `config.staging.yml`
Staging/Pre-production environment configuration.
- **Use case:** Pre-production testing environment
- **Environment:** Set when `CARACAL_ENV=staging`
- **Overrides:** Base config values for staging
- **Contains:** Application config specific to staging

**Example:**
```yaml
app:
  name: caracal
  version: 0.1.0
  description: "Caracal - A high-performance Go application"
  environment: staging
```

### `config.prod.yml`
Production environment configuration.
- **Use case:** Production deployment
- **Environment:** Set when `CARACAL_ENV=prod`
- **Overrides:** Base config values for production
- **Contains:** Application config specific to production

**Example:**
```yaml
app:
  name: caracal
  version: 0.1.0
  description: "Caracal - A high-performance Go application"
  environment: production
```

## Application Configuration Section

### `app`
Application metadata and settings.

**Fields:**
- `name` (string): Application name
  - Example: `caracal`
  - Required: Yes
  
- `version` (string): Semantic version (X.Y.Z)
  - Example: `0.1.0`
  - Required: Yes
  
- `description` (string): Application description
  - Example: `Caracal - A high-performance Go application`
  - Required: Yes
  
- `environment` (string): Environment type
  - Allowed values: `development`, `staging`, `production`
  - Default: `development`
  - Required: Yes

**Example:**
```yaml
app:
  name: caracal
  version: 0.1.0
  description: "Caracal - A high-performance Go application"
  environment: development
```

## Usage

### Loading Configuration

#### Using Environment Variable (Recommended)
```bash
export CARACAL_ENV=dev
./caracal
# Loads config.yml first, then config.dev.yml overrides
```

#### Default Behavior
```bash
./caracal
# Uses CARACAL_ENV=development by default
# Loads config.yml first, then config.dev.yml overrides
```

#### Different Environments
```bash
# Development
export CARACAL_ENV=dev
./caracal

# Staging
export CARACAL_ENV=staging
./caracal

# Production
export CARACAL_ENV=prod
./caracal
```

### Environment Variables

Override configuration with environment variables using the format:
```
CARACAL_<SECTION>_<KEY>=value
```

Examples:
```bash
CARACAL_APP_NAME=my-app
CARACAL_APP_VERSION=1.0.0
CARACAL_APP_ENVIRONMENT=production
```

## Configuration in Go Code

The `config/config.go` file provides functions to load and access configuration:

### Load Configuration
```go
package main

import (
	"caracal/config"
)

func main() {
	cfg := config.Load()
	// cfg.App.Name contains the loaded application name
	// cfg.App.Version contains the loaded application version
	// cfg.App.Environment contains the loaded environment
}
```

### Get App Info
```go
appInfo := cfg.GetAppInfo()
// Returns: "caracal v0.1.0"
```

## How Configuration Loading Works

1. **Load Base Config**
   - Reads `config/config.yml`
   - Serves as defaults for all fields

2. **Determine Environment**
   - Reads `CARACAL_ENV` environment variable
   - Defaults to `development` if not set

3. **Load Environment-Specific Config**
   - Reads `config/config.{environment}.yml`
   - Example: if `CARACAL_ENV=staging`, loads `config/config.staging.yml`

4. **Merge Configurations**
   - Non-empty values from environment-specific config override base values
   - Base values are preserved if environment-specific values are empty

5. **Return Final Configuration**
   - Returns merged configuration with all values

## Examples

### Example 1: Development Environment
```bash
export CARACAL_ENV=dev
./caracal
```

**Loading process:**
1. Load `config/config.yml` (base)
   - name: caracal
   - version: 0.1.0
   - environment: development

2. Load `config/config.dev.yml` (overrides)
   - All values same as base

3. **Result:** Development configuration loaded

### Example 2: Production Environment
```bash
export CARACAL_ENV=prod
./caracal
```

**Loading process:**
1. Load `config/config.yml` (base)
   - name: caracal
   - version: 0.1.0
   - environment: development

2. Load `config/config.prod.yml` (overrides)
   - name: caracal
   - version: 0.1.0
   - environment: production

3. **Result:** Production configuration loaded with environment set to production

### Example 3: Override with Environment Variable
```bash
export CARACAL_ENV=dev
export CARACAL_APP_VERSION=2.0.0
./caracal
```

**Loading process:**
1. Load `config/config.yml` (base)
2. Load `config/config.dev.yml` (overrides)
3. **Environment variable overrides both:** `CARACAL_APP_VERSION=2.0.0`
4. **Result:** Version is 2.0.0 from environment variable

## Best Practices

### Security

1. **Never commit sensitive data**
   - Use environment variables for any sensitive values
   - Keep production credentials out of YAML files

2. **Use different configurations per environment**
   - Development: Local settings
   - Staging: Production-like settings for testing
   - Production: Security-hardened settings

3. **Environment Variables in Production**
   - Override all sensitive values via environment variables
   - Never store credentials in config files

### Development

1. **Use `config.dev.yml` for local overrides**
   - Keep base `config.yml` as universal defaults
   - Override only what's different in development

2. **Test all environments locally**
   ```bash
   CARACAL_ENV=dev ./caracal
   CARACAL_ENV=staging ./caracal
   CARACAL_ENV=prod ./caracal
   ```

3. **Keep configurations minimal**
   - Only override values that differ from base
   - Leave other values in base config

### Deployment

1. **Set `CARACAL_ENV` in deployment**
   - Development: `CARACAL_ENV=dev`
   - Staging: `CARACAL_ENV=staging`
   - Production: `CARACAL_ENV=prod`

2. **Use environment variables for secrets**
   - Database passwords
   - API keys
   - JWT secrets

3. **Validate configuration on startup**
   - Check that all required fields are set
   - Log configuration for debugging

## Troubleshooting

### Configuration Not Loading

1. **Check file exists:**
   ```bash
   ls -la config/config.yml
   ls -la config/config.dev.yml
   ```

2. **Verify environment variable:**
   ```bash
   echo $CARACAL_ENV
   ```

3. **Check file permissions:**
   ```bash
   chmod 644 config/*.yml
   ```

4. **Validate YAML syntax:**
   - Use online YAML validator
   - Check for indentation errors

### Wrong Environment Loaded

1. **Verify `CARACAL_ENV` is set:**
   ```bash
   export CARACAL_ENV=prod
   echo $CARACAL_ENV
   ```

2. **Restart application after changing env var:**
   ```bash
   export CARACAL_ENV=staging
   ./caracal
   ```

3. **Check which config file is being used:**
   - Add debug logging to see which files are loaded
   - Print loaded configuration values

### Values Not Overriding

1. **Ensure environment-specific file exists:**
   ```bash
   ls -la config/config.{environment}.yml
   ```

2. **Check YAML indentation:**
   - Must use spaces (not tabs)
   - Same structure as base config

3. **Verify field names match:**
   - Base config: `app.name`
   - Override should be: `app.name`

## Adding New Configuration Fields

When adding new configuration fields to the app section:

1. **Update `config.yml` (base):**
   ```yaml
   app:
     name: caracal
     version: 0.1.0
     description: "..."
     environment: development
     new_field: default_value
   ```

2. **Update `config/config.go` (Go struct):**
   ```go
   type AppConfig struct {
     Name        string `yaml:"name"`
     Version     string `yaml:"version"`
     Description string `yaml:"description"`
     Environment string `yaml:"environment"`
     NewField    string `yaml:"new_field"`
   }
   ```

3. **Update environment-specific configs:**
   ```yaml
   # config.dev.yml, config.staging.yml, config.prod.yml
   app:
     new_field: environment_specific_value
   ```

4. **Update merge logic (if needed):**
   - Add merge handling in `mergeConfigs()` function in `config/config.go`

5. **Update this documentation**

## Configuration File Locations

All configuration files are located in the `config/` directory:
- `config/config.yml` - Base configuration
- `config/config.dev.yml` - Development configuration
- `config/config.staging.yml` - Staging configuration
- `config/config.prod.yml` - Production configuration

---

**Last Updated:** March 6, 2026  
**Version:** 2.0  
**Status:** Production Ready

## Changelog

### Version 2.0 (March 6, 2026)
- Updated for new configuration loading strategy
- Base config loaded first, then overridden by environment-specific configs
- Simplified to only include app configuration section
- Added detailed loading process explanation
- Added multiple examples of configuration loading
- Removed Debug field documentation
- Added Configuration Loading Strategy section

### Version 1.0
- Initial configuration guide
