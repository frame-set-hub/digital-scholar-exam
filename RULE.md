# Documentation sync rules

When you **change code** (or project layout) and that change **affects** what Markdown docs describe — **update the relevant `.md` files in the same spirit** so docs do not describe old behavior.

## What counts as “affects”

- Changing flows, API, state, or file/folder names that docs reference
- Moving/deleting/merging lines so line ranges in [docs/code_analyze.md](./docs/code_analyze.md) no longer match the code
- Changing behavior described in [README.md](./README.md), [execute.md](./execute.md), [docs/planning.md](./docs/planning.md), [docs/architech.md](./docs/architech.md), [docs/api.md](./docs/api.md)

## Practice

1. Identify what the change affects (e.g. no frontend mock fallback anymore → update every mention of MOCK / fallback)
2. Update **every related `.md` file** in the same merge as the code change (or immediately after)
3. For the **code map** in `docs/code_analyze.md`: verify line ranges against real files in the repo (open files or use a line counter), then fix tables to match

## Primary documentation

| Role | File |
|--------|------|
| Repo entry | [README.md](./README.md) |
| In-depth index | [docs/README.md](./docs/README.md) |
| Code walkthrough (lines) | [docs/code_analyze.md](./docs/code_analyze.md) |
| Architecture / flows | [docs/architech.md](./docs/architech.md) |
| API | [docs/api.md](./docs/api.md) |
| Plans / roadmap | [docs/planning.md](./docs/planning.md) |
| Progress | [execute.md](./execute.md) |

This file is a project rule — it does not replace each doc’s own explanation, but requires keeping docs in sync when changes affect them.
