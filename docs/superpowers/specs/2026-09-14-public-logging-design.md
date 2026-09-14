# Public logging boundary

Approved approach: preserve Debug/Info/Warn/Error/Fatal formatting calls and
standard logging output, without database persistence or generated DB schemas.
Fatal remains a logging level, not process termination, matching existing behavior.

The existing executable only calls Info for build metadata. No in-repository
caller initializes database logging. Remove LogSetup, LogEntries, background
database/CSV retention and automatic directory deletion. This intentionally
breaks compatibility for external callers of the removed persistence API.

Implement a small standard-library logger wrapper with level, target, caller
and formatted message. Report caller basename rather than build-machine paths.
Do not add buffering, goroutines, DB adapters or automatic secret-redaction claims.
Callers remain responsible for not logging credentials.

Remove the 61 generated model source files after checking all references. Remove
GORM dependencies only once no other source uses them. Preserve legacy binaries
for separate review; they are not rebuilt or certified by this source change.

Validate all levels, formatting, caller information and concurrent calls with
Go tests and race detection. Add a regression guard against the old schema and
persistence dependencies. Build all packages; do not start the service or access
operating databases. Existing Python configuration checks must still pass.

No history rewrite, dev deletion, operational deployment or token changes.
Source cleanup does not remove old commits, PR refs, binaries or caches.
