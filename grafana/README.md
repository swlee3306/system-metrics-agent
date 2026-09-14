# Grafana configuration boundary

Reviewed on 2026-09-14. The existing `datasource.yaml` contains Prometheus,
InfluxDB and TimescaleDB definitions, including password literals. Their validity
and any external provisioning/mount usage are unknown. No Grafana service was
contacted and no datasource was provisioned during this review.

Do not apply that file as a ready-to-use public example. It is retained unchanged
because an external deployment may consume this path. The absence of an in-repo
reference does not prove that the file is unused.

## Preparing a separate example

1. Select only the datasource actually used by your test environment. This project
   does not claim verified integration with all three listed backends.
2. Create a separately named, untracked local configuration. Use synthetic names
   and documentation-only endpoints when sharing the structure publicly; omit all
   real passwords, tokens and internal host details.
3. Supply credentials using the supported secret mechanism for the installed
   Grafana and datasource versions. Do not assume a legacy password field has the
   correct semantics for your version.
4. Verify network reachability, authentication, TLS and a read-only sample query in
   a disposable environment before connecting it to an operating deployment.
5. Switch an existing mount/provisioning path only after its owner and consumers
   are identified. Keep rollback material private if it contains secrets.

This document does not provide a replacement runtime configuration or claim a
working Grafana dashboard. Existing password literals and prior Git history are
not remediated by adding this guide. No passwords were rotated or tested.
