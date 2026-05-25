# yarr — Tester-Env Deployment Notes

## Baseline
- Upstream: https://github.com/nkanaev/yarr
- Baseline commit: `31274d17a542173490e0359ebd86b6a2c394486a`
- Tag: v2.6 (pre-release, master HEAD ~2025-01)

## Build
```bash
docker build -t tester-env-yarr -f etc/dockerfile .
```
Builds Go binary with `CGO_ENABLED=1` and SQLite tags (`sqlite_foreign_keys sqlite_json`).
Embedded web UI served from `src/assets/` (Go `embed.FS`).

## Run
```bash
docker run -d --name tester-env-yarr -p 7070:7070 -v yarr-data:/data tester-env-yarr -addr 0.0.0.0:7070 -db /data/yarr.db
```
- Port: 7070
- SQLite database: `/data/yarr.db` (volume `yarr-data`)
- Auth optional (open access without `-auth` flag)

## Reset
```bash
docker stop tester-env-yarr && docker rm tester-env-yarr && docker volume rm yarr-data
```
Completely destroys container + database volume. Fresh start on next deploy.

## Seed (Deterministic)
Uses `./tester-env` CLI (this repo):
```bash
./tester-env build
./tester-env deploy
./tester-env seed
./tester-env verify
```
Creates 3 folders, 3 static RSS feeds, 15 articles (12 unread, 1 starred, 2 read):
- **Product Updates** / Aurora SaaS Changelog (5 items)
- **Engineering Radar** / Platform Engineering Radar (5 items)
- **Customer Signals** / Customer Voice Digest (5 items)

Static RSS fixtures live under `src/assets/tester-env/rss/` and are embedded into the Go binary via `src/assets/assetsfs.go`.

## Verify
```bash
./tester-env verify
```
Asserts:
- 3 folders, 3 feeds, 15 items, 12 unread, 1 starred
- Search "quartz lighthouse" returns exactly 1 result
- Newest item: "Enterprise onboarding notes request clearer SSO copy"
- Ordering: newest-first by published date

## Browser Smoke
- URL: http://localhost:7070/
- Visible: 3 folders sidebar, article list with read/unread/starred indicators, search bar, settings cog.
- Starred item: "Mobile receipt preview wins praise from beta cohort"

## Environment Variables
| Variable | Default | Description |
|---|---|---|
| YARR_IMAGE | tester-env-yarr | Docker image tag |
| YARR_PORT | 7070 | Host port |
| YARR_CONTAINER | tester-env-yarr-<run-id> | Container name |
| YARR_VOLUME | yarr-data-<run-id> | Docker volume name |
| YARR_URL | http://localhost:<port> | Base URL |
