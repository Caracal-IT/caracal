# Configuration Guide for Caracal

This directory contains all YAML configuration files for the caracal application.

## Configuration Files

### `config.yml`
The base configuration file with all available settings and their defaults.
- **Use case:** Reference and template for all configurations
- **Included:** All possible configuration options with descriptions

### `config.dev.yml`
Development environment configuration.
- **Debug Mode:** Enabled (`debug: true`)
- **Logging:** Debug level, console output, file logging enabled
- **Database:** SQLite (local development database)
- **Cache:** In-memory cache for quick development
- **Features:** All alpha and experimental features enabled
- **Profiling:** All profiling enabled (CPU, memory, goroutine)

### `config.staging.yml`
Staging/Pre-production environment configuration.
- **Debug Mode:** Disabled (`debug: false`)
- **Logging:** Info level, JSON format, file logging enabled
- **Database:** PostgreSQL with SSL preference
- **Cache:** Redis backend for performance testing
- **Features:** Alpha features enabled, experimental disabled
- **Monitoring:** Metrics and tracing with 10% sample rate

### `config.prod.yml`
Production environment configuration.
- **Debug Mode:** Disabled (`debug: false`)
- **Logging:** Warning level only, JSON format, no console output
- **Database:** PostgreSQL with SSL required
- **Cache:** Redis with optimized pool size
- **Features:** Only stable features enabled
- **Monitoring:** Tracing with 1% sample rate (low overhead)

## Configuration Sections

### `app`
Application metadata and settings.
- `name`: Application name
- `version`: Semantic version (X.Y.Z)
- `description`: Application description
- `environment`: Environment type (development, staging, production)
- `debug`: Enable debug mode

### `server`
HTTP server configuration.
- `host`: Server bind address
- `port`: Server port
- `read_timeout`: Request read timeout (seconds)
- `write_timeout`: Response write timeout (seconds)
- `idle_timeout`: Connection idle timeout (seconds)
- `max_connections`: Maximum concurrent connections
- `graceful_shutdown`: Enable graceful shutdown
- `shutdown_timeout`: Graceful shutdown timeout (seconds)

### `logging`
Logging configuration.
- `level`: Log level (debug, info, warn, error)
- `format`: Log format (json, text)
- `console`: Enable console output
- `file`:
  - `enabled`: Enable file logging
  - `path`: Log file path
  - `max_size`: Max file size (MB)
  - `max_backups`: Number of backup files to keep
  - `max_age`: Max days to retain logs
  - `compress`: Compress rotated files

### `database`
Database connection configuration.
- `type`: Database type (postgres, mysql, sqlite)
- `host`: Database server host
- `port`: Database server port
- `name`: Database name
- `user`: Database user
- `password`: Database password (use env vars in production)
- `pool`: Connection pool settings
  - `min_connections`: Minimum pool size
  - `max_connections`: Maximum pool size
  - `idle_timeout`: Idle timeout (seconds)
  - `max_lifetime`: Max connection lifetime (seconds)
- `ssl`: SSL configuration
  - `enabled`: Enable SSL
  - `mode`: SSL mode (disable, allow, prefer, require, verify-ca, verify-full)

### `cache`
Cache configuration.
- `backend`: Cache backend (memory, redis)
- `ttl`: Cache TTL (seconds)
- `max_size`: Max cache size (MB, memory backend only)
- `redis`: Redis configuration
  - `host`: Redis host
  - `port`: Redis port
  - `db`: Redis database number
  - `password`: Redis password
  - `pool_size`: Connection pool size

### `security`
Security settings.
- `cors`: CORS configuration
  - `enabled`: Enable CORS
  - `allowed_origins`: List of allowed origins
  - `allowed_methods`: List of allowed HTTP methods
  - `allowed_headers`: List of allowed headers
  - `allow_credentials`: Allow credentials
- `jwt`: JWT configuration
  - `enabled`: Enable JWT
  - `secret_key`: JWT secret key (use env vars)
  - `expiration`: Token expiration (seconds)
  - `algorithm`: JWT algorithm (HS256, HS512, RS256)

### `api`
API configuration.
- `base_path`: API base path
- `version`: API version
- `rate_limit`: Rate limiting
  - `enabled`: Enable rate limiting
  - `requests_per_second`: RPS limit
  - `burst`: Burst size

### `features`
Feature flags.
- `alpha_features`: Enable alpha features
- `experimental_endpoints`: Enable experimental endpoints
- `metrics`: Enable metrics collection
- `health_checks`: Enable health checks

### `monitoring`
Observability configuration.
- `metrics`: Metrics configuration
  - `enabled`: Enable metrics
  - `port`: Metrics port
  - `path`: Metrics path
- `tracing`: Tracing configuration
  - `enabled`: Enable tracing
  - `sample_rate`: Sampling rate (0.0-1.0)
  - `exporter`: Exporter type (jaeger, datadog, stdout)

### `dev`
Development tools configuration.
- `hot_reload`: Enable hot reload
- `watch_dirs`: Directories to watch for changes
- `profiling`: Profiling options
  - `cpu`: CPU profiling
  - `memory`: Memory profiling
  - `goroutine`: Goroutine profiling

## Usage

### Loading Configuration

Using environment variable (recommended):
```bash
export CARACAL_ENV=dev
# Application will load config/config.dev.yml
```

Using command-line flag:
```bash
./caracal --config=config/config.dev.yml
```

### Environment Variables

Override configuration with environment variables using the format:
```
CARACAL_<SECTION>_<KEY>=value
```

Examples:
```bash
CARACAL_APP_DEBUG=true
CARACAL_SERVER_PORT=9000
CARACAL_LOGGING_LEVEL=debug
CARACAL_DATABASE_HOST=db.example.com
```

### Configuration Priority

Configuration is loaded in this priority order (highest to lowest):
1. Environment variables (`CARACAL_*`)
2. Environment-specific YAML file (e.g., `config.dev.yml`)
3. Base YAML file (`config.yml`)
4. Application defaults

## Best Practices

### Security

1. **Never commit sensitive data** - Always use environment variables for:
   - Database passwords
   - JWT secrets
   - API keys
   - Credentials

2. **Use different credentials** for each environment:
   - Development: Local credentials
   - Staging: Staging-specific credentials
   - Production: Production-specific credentials

3. **Enable SSL in production:**
   - Always use `ssl.mode: require` in production
   - Use certificate management (Let's Encrypt, etc.)

### Performance

1. **Development:**
   - Use SQLite for simplicity
   - Use in-memory cache
   - Enable all debug features

2. **Staging:**
   - Use same infrastructure as production
   - Test with realistic data volumes
   - Monitor performance metrics

3. **Production:**
   - Use PostgreSQL for reliability
   - Use Redis for distributed caching
   - Optimize pool sizes based on load
   - Use JSON logging for log aggregation

### Monitoring

1. **Enable metrics** in all environments
2. **Adjust sampling rate** based on volume:
   - Development: 100% (1.0)
   - Staging: 10% (0.1)
   - Production: 1% (0.01)

3. **Log levels by environment:**
   - Development: debug
   - Staging: info
   - Production: warn or error

## Environment-Specific Guidelines

### Development (`config.dev.yml`)
- Maximum verbosity and debugging
- Fast feedback loop
- All features enabled for testing
- Local resources (SQLite, memory cache)

### Staging (`config.staging.yml`)
- Production-like infrastructure
- Balanced logging and monitoring
- Feature testing before production
- Performance validation

### Production (`config.prod.yml`)
- Security hardening
- Performance optimization
- Minimal logging overhead
- Only stable features
- High availability settings

## Troubleshooting

### Configuration Not Loading

1. Check file exists: `ls -la config/config.yml`
2. Verify environment variable: `echo $CARACAL_ENV`
3. Check file permissions: `chmod 644 config/*.yml`
4. Validate YAML syntax: Use YAML validator online

### Environment Variables Not Working

1. Verify prefix: Should be `CARACAL_*`
2. Use underscores for nested keys: `CARACAL_DATABASE_HOST`
3. Export variable: `export CARACAL_*=value`
4. Restart application after setting variables

### Performance Issues

1. Check database pool settings
2. Adjust cache TTL
3. Monitor log level (reduce verbosity in production)
4. Check sampling rate (increase for more data)

## Adding New Configuration Options

When adding new configuration options:

1. Add to base `config.yml` with description
2. Add to environment-specific files with appropriate values
3. Update this README with section and options
4. Ensure backward compatibility
5. Test all environments

---

**Last Updated:** March 6, 2026  
**Version:** 1.0  
**Status:** Production Ready

