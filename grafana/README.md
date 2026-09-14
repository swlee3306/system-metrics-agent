# Grafana configuration boundary

## Latest verification — supersedes the historical review below

On 2026-09-14 the public example was tested in a disposable Grafana 13.2.1
container with `--network none`, no published ports or host-directory mounts,
and a synthetic administrator password. No operating service was contacted.
Legacy top-level `password` properties did not register secure password fields.
They now use `secureJsonData`, with a regression test. The values remain explicit
non-working placeholders. No real datasource query or authentication is claimed.

After reprovisioning, `/api/health` reported database `ok`; all three datasources
were registered, and individual InfluxDB/TimescaleDB API responses showed
`secureJsonFields.password: true` without returning the placeholder values.
Image digest: `sha256:f772d434e8fab0049deb2b1b30abd43342bcfca1537614aa8d36080232cf4283`.
The network-blocked plugin-pattern update warning is expected in this test.

Inside a container, localhost means that container, not the host or another
datasource container. Use a separate private configuration for your actual
addresses, authentication and TLS. Never commit operational credentials.

Known fixed passwords and generated DB models were removed from published
main/dev histories on 2026-09-14. Read-only GitHub PR refs, caches and external
copies remain a separate scope. No operating password was rotated or tested.

Official format: https://grafana.com/docs/grafana/latest/administration/provisioning/

## Historical notes — describe the state before the changes above

## Public configuration update — 2026-09-14

Both password fields in `datasource.yaml` now contain explicit `REPLACE_WITH_...`
placeholders. They are not working credentials and are not automatically expanded.
Use a private configuration and an appropriate secret mechanism before provisioning.
The ARM64 activation/deactivation scripts no longer contain a default server IP or
personal SSH key path: arguments 2 and 4 must be supplied. Do not run the destructive
deactivation script as a smoke test. Verification only exercised its missing-argument
failure path, before any SSH command.

No operating Grafana or remote service was changed. Historical literals may still
exist in Git history. The notes below describe the pre-cleanup configuration.

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
