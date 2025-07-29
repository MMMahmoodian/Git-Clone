# Step 5: Commit Objects (`git commit`)

## Scope:
- Generate commit objects linking file snapshots.
- Include commit metadata (author, message, timestamp).

## Goals:
- Build a functioning commit history.

## Implementation Steps:
- Create commit objects referencing a snapshot (tree object) and parent commits.
- Store commits in `.git/objects`.

## Expected Result:
- Commits created and retrievable through commit history.
