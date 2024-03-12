package types

// PolcyCommandMaxExpirationDelta configures the maximum lifetime of a PolicyCmd signed payload.
// Since SourceHub is expected to operate roughly with 1-2 seconds block time,
// the paylaod would live for roughly 12-24h
const DefaultPolicyCommandMaxExpirationDelta = 60 * 60 * 12
