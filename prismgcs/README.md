# Reproduce Failures with Prism

This code doesn't work with the prism runner and a GCS file.

To reproduce

1. Copy hackedlogs.json to a GCS bucket
2. Run it

```bash
go run . --inputs=gs://app-logs-dev-sailplane/hackedlogs.json
```

Runner appears to keep reading the same file over and over again.

```bash
2024/02/19 15:17:19 INFO Reading from gs://<BUCKET>/hackedlogs.json source=/Users/jlewi/go/pkg/mod/github.com/apache/beam/sdks/v2@v2.54.0/go/pkg/beam/io/textio/textio.go:226 time=2024-02-19T23:17:19.971Z worker.ID=job-001[go-job-1-1708384631797458000]_go worker.endpoint=localhost:65326
```

Using local input works just fine.