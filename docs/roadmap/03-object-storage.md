# Step 3: Object Storage (`git hash-object`, `git cat-file`)

## Scope:
- Implement SHA-1 object hashing.
- Compress and store file contents.

## Goals:
- Store and retrieve file contents in `.git/objects`.

## Implementation Steps:
- Hash file contents using SHA-1 with appropriate Git object format.
- Compress objects using zlib compression.
- Retrieve stored objects based on hash.

## Expected Result:
- Objects stored and retrieved accurately from `.git/objects`.
