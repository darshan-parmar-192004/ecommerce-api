# Backup and Restore Runbook

## Overview

This runbook covers database backup and restore procedures for the Ecommerce API PostgreSQL database.

## Backup Strategy

### Types of Backups

| Type | Frequency | Purpose |
|------|-----------|---------|
| Full dump (pg_dump) | Daily | Complete data recovery |
| WAL archiving | Continuous | Point-in-time recovery |
| Config backup | On change | System configuration |

### Backup Schedule

| Backup Type | Schedule | Retention |
|-------------|----------|-----------|
| Daily full | 02:00 UTC | 7 days |
| Weekly full | Sunday 02:00 UTC | 4 weeks |
| Monthly full | 1st of month | 12 months |
| WAL archives | Continuous | 7 days |

## Manual Backup

### Full Database Backup (pg_dump)

```bash
# Create backup directory
mkdir -p /backups/ecommerce

# Run pg_dump (replace with your values)
pg_dump -h localhost -U ecommerce_user -d ecommerce \
  -F c -b -v -f /backups/ecommerce/backup_$(date +%Y%m%d_%H%M%S).dump

# List backups
ls -la /backups/ecommerce/
```

### Backup Options Explained

| Flag | Description |
|------|-------------|
| `-F c` | Custom format (compressed, parallel) |
| `-b` | Include large objects (blobs) |
| `-v` | Verbose output |
| `-f` | Output file |

### Backup to S3 (Optional)

```bash
# Install AWS CLI first
# Then upload backup to S3
aws s3 cp /backups/ecommerce/backup_20240115_020000.dump s3://your-bucket/backups/
```

### Compressed SQL Backup (Alternative)

```bash
# Plain SQL format (readable, larger)
pg_dump -h localhost -U ecommerce_user -d ecommerce \
  -F p | gzip > /backups/ecommerce/backup_$(date +%Y%m%d).sql.gz
```

## Automated Backup Script

```bash
#!/bin/bash
# /opt/ecommerce-api/scripts/backup.sh

set -e

# Configuration
BACKUP_DIR="/backups/ecommerce"
DB_NAME="ecommerce"
DB_USER="ecommerce_user"
DB_HOST="localhost"
RETENTION_DAYS=7

# Create backup directory if not exists
mkdir -p "$BACKUP_DIR"

# Generate backup filename
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="$BACKUP_DIR/ecommerce_${TIMESTAMP}.dump"

# Run backup
echo "Starting backup at $(date)"
pg_dump -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" \
  -F c -b -v -f "$BACKUP_FILE"

# Verify backup
if [ -f "$BACKUP_FILE" ]; then
    SIZE=$(du -h "$BACKUP_FILE" | cut -f1)
    echo "Backup completed: $BACKUP_FILE ($SIZE)"
else
    echo "ERROR: Backup failed"
    exit 1
fi

# Cleanup old backups
find "$BACKUP_DIR" -name "ecommerce_*.dump" -mtime +$RETENTION_DAYS -delete
echo "Old backups cleaned up"

echo "Backup process completed at $(date)"
```

### Make Script Executable

```bash
chmod +x /opt/ecommerce-api/scripts/backup.sh

# Add to crontab for daily backup
crontab -e
# Add: 0 2 * * * /opt/ecommerce-api/scripts/backup.sh
```

## Restore Procedures

### Full Database Restore

```bash
# Stop the application
systemctl stop ecommerce-api
# Or: podman-compose stop app

# Drop existing database (optional - creates fresh start)
# psql -h localhost -U ecommerce_user -d postgres -c "DROP DATABASE ecommerce;"
# psql -h localhost -U ecommerce_user -d postgres -c "CREATE DATABASE ecommerce;"

# Restore from backup
pg_restore -h localhost -U ecommerce_user -d ecommerce \
  -v /backups/ecommerce/backup_20240115_020000.dump

# Start the application
systemctl start ecommerce-api
# Or: podman-compose up -d
```

### Restore Options

| Flag | Description |
|------|-------------|
| `-v` | Verbose output |
| `-j 4` | Parallel jobs (faster restore) |
| `--data-only` | Restore data only, not schema |
| `--schema-only` | Restore schema only, not data |

### Point-in-Time Recovery (PITR)

For PITR, you need:
1. Base backup
2. WAL archives

```bash
# Stop PostgreSQL
systemctl stop postgresql

# Configure recovery.conf (PostgreSQL 12+ uses postgresql.conf)
# Edit /etc/postgresql/16/main/postgresql.conf:
# restore_command = 'cp /path/to/wal/%f %p'

# Create recovery signal file
touch /var/lib/postgresql/16/main/recovery.signal

# Start PostgreSQL
systemctl start postgresql
```

## Verification

### Verify Backup Integrity

```bash
# List contents of backup
pg_restore --list /backups/ecommerce/backup_20240115_020000.dump

# Test restore to temporary database
createdb -h localhost -U ecommerce_user ecommerce_test
pg_restore -h localhost -U ecommerce_user -d ecommerce_test \
  /backups/ecommerce/backup_20240115_020000.dump

# Verify tables exist
psql -h localhost -U ecommerce_user -d ecommerce_test -c "\dt"

# Clean up test database
dropdb -h localhost -U ecommerce_user ecommerce_test
```

### Verify Database Functionality

```bash
# After restore, verify:
curl http://localhost:8080/health
# Should return: {"status": "ok"}
```

## Disaster Recovery

### Complete Data Loss

```bash
# 1. Stop all services
podman-compose down

# 2. Recreate PostgreSQL volume (if using containers)
podman volume rm ecommerce-api_psql_volume_bp
podman volume create ecommerce-api_psql_volume_bp

# 3. Restore from latest backup
podman-compose up -d psql_bp
# Wait for PostgreSQL to be ready
pg_restore -h localhost -U postgres -d ecommerce \
  /backups/ecommerce/backup_latest.dump

# 4. Start all services
podman-compose up -d

# 5. Verify
curl http://localhost:8080/health
```

### Off-Site Backup (DR)

```bash
# Copy backups to off-site location
rsync -avz --delete /backups/ecommerce/ user@backup-server:/path/to/backups/

# Or use cloud storage
aws s3 sync /backups/ecommerce/ s3://your-bucket/ecommerce-backups/
```

## Monitoring

### Check Backup Success

```bash
# Add to monitoring script
LAST_BACKUP=$(ls -lt /backups/ecommerce/*.dump | head -1 | awk '{print $NF}')
AGE_HOURS=$(($(date +%s) - $(stat -c %Y "$LAST_BACKUP")) / 3600)

if [ $AGE_HOURS -gt 25 ]; then
    echo "WARNING: Last backup is older than 25 hours"
    # Send alert
fi
```

## Backup Retention Policy

| Backup Type | Keep For |
|-------------|----------|
| Daily backups | 7 days |
| Weekly backups | 4 weeks |
| Monthly backups | 12 months |

## Testing Restore Procedure

**IMPORTANT**: Test your restore procedure regularly!

```bash
# Quarterly: Test restore to a test environment
# 1. Create test database
createdb -h localhost -U ecommerce_user ecommerce_restore_test

# 2. Restore backup
pg_restore -h localhost -U ecommerce_user -d ecommerce_restore_test \
  /backups/ecommerce/backup_latest.dump

# 3. Verify data integrity
psql -h localhost -U ecommerce_user -d ecommerce_restore_test -c "SELECT COUNT(*) FROM customers;"

# 4. Clean up
dropdb -h localhost -U ecommerce_user ecommerce_restore_test
```