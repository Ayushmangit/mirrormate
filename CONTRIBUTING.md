# MirrorMate — Team Workflow & Commit Guide

This document defines how our 4-person team (2 frontend, 2 backend) works together on
[mirrormate](https://github.com/Ayushmangit/mirrormate) so nobody's commits, branches, or PRs turn into a mess.

Everyone should read this once and follow it every time. Drop this file at the repo root as `CONTRIBUTING.md`
so it shows up automatically when someone opens a PR.

---

## 1. Roles

| Person | Area | Owns |
|---|---|---|
| Dev 1 | Frontend | `frontend/` |
| Dev 2 | Frontend | `frontend/` |
| Dev 3 | Backend | `backend/` |
| Dev 4 | Backend | `backend/` |

Cross-cutting changes (e.g. API contract changes that touch both sides) should be flagged in the related GitHub
Issue so both frontend devs and both backend devs are aware before work starts.

---

## 2. Branching rules

**Never push directly to `main`.** All work happens on a branch, merged via Pull Request.

### Branch naming

```
<type>/<short-description>
```

Types:
- `feature/` — new functionality
- `fix/` — bug fix
- `refactor/` — code change with no behavior change
- `docs/` — documentation only
- `chore/` — tooling, config, dependencies
- `hotfix/` — urgent production fix

Use kebab-case, no spaces, no ticket-number-only names.

**Good:**
```
feature/user-profile-page
feature/redis-rate-limiter
fix/auth-token-expiry
refactor/store-layer
docs/update-readme
chore/upgrade-vite
```

**Bad:**
```
patch1
new-stuff
ayush-branch
test
```

If you're working from an Issue, include the issue number:
```
feature/23-user-profile-page
fix/41-auth-token-expiry
```

---

## 3. Commit message convention

We use **Conventional Commits**. Every commit message follows this format:

```
<type>(<scope>): <short summary>

[optional body]

[optional footer]
```

### Type (required)

| Type | When to use it |
|---|---|
| `feat` | A new feature |
| `fix` | A bug fix |
| `refactor` | Code change that doesn't add a feature or fix a bug |
| `docs` | Documentation only |
| `style` | Formatting, whitespace, no logic change |
| `test` | Adding or fixing tests |
| `chore` | Build process, dependencies, tooling |
| `perf` | Performance improvement |
| `migration` | Database migration (Goose) |

### Scope (recommended)

The part of the codebase you touched. Keep it short and consistent:

- Backend: `auth`, `store`, `db`, `mailer`, `ratelimiter`, `env`, `api`
- Frontend: `ui`, `auth`, `routing`, `api-client`, `profile`, `comments`
- Other: `docker`, `ci`, `readme`, `env`

### Summary (required)

- Lowercase, imperative mood ("add", not "added" or "adds")
- No period at the end
- Under ~60 characters
- Say *what* changed, not *how*

### Examples

**Good:**
```
feat(auth): add JWT refresh token endpoint
fix(store): correct nil pointer on empty user query
refactor(ratelimiter): simplify redis key generation
docs(readme): update local setup instructions
migration(db): add index on users.email
feat(profile): add avatar upload component
fix(api-client): handle 401 by redirecting to login
chore(deps): bump vite to 6.1.0
```

**Bad:**
```
fixed stuff
update
WIP
asdf
final fix for real this time
changes
```

### Body (optional, use for anything non-trivial)

Explain *why*, not just what — the diff already shows what changed.

```
fix(auth): prevent token refresh race condition

Two concurrent refresh requests could both succeed and issue
different tokens, invalidating the first. Added a mutex lock
around the refresh handler.

Closes #37
```

### Multiple small commits vs one commit

- Commit often locally, but before opening a PR, squash trivial "wip", "typo", "fix test" commits into
  logical, meaningful commits. Nobody wants to review 14 commits called `fix`.
- One commit should represent one logical change.

---

## 4. Pull Request rules

1. **One PR = one feature/fix.** Don't bundle unrelated changes.
2. **Title follows the same convention** as commits, e.g. `feat(auth): add password reset flow`.
3. **Description must include:**
   - What changed and why
   - How to test it
   - `Closes #<issue-number>` if it resolves an issue
4. **Every PR needs at least 1 approval** from a teammate before merging — ideally from someone on the *other*
   side of the stack if the change is cross-cutting, otherwise anyone.
5. **Run checks before opening the PR:**

   Backend:
   ```
   cd backend
   gofmt -w .
   go vet ./...
   go test ./...
   go build ./cmd/api
   ```

   Frontend:
   ```
   cd frontend
   npm run lint
   npm run build
   ```
6. **Never commit secrets** — no `.envrc`, API keys, JWT secrets, or DB passwords. Use `.envrc.example` as the
   template.
7. **Database changes:** never edit an already-applied migration. Create a new one, and commit the migration
   together with the code that needs it.
8. **Merge strategy:** use **Squash and merge** into `main` so the history on `main` stays one clean commit per
   feature/fix, matching the PR title. Delete the branch after merging.
9. **Resolve merge conflicts on your own branch** (`git fetch origin && git rebase origin/main`), not in the
   GitHub web UI, once conflicts get non-trivial.

---

## 5. Daily workflow

```
1. git checkout main && git pull
2. git checkout -b feature/23-user-profile-page
3. Do the work, commit in logical chunks with conventional commit messages
4. Run lint/tests/build locally
5. git push -u origin feature/23-user-profile-page
6. Open PR, link the issue, request review
7. Address review comments (fix, don't force-push over an active review unless asked)
8. Get 1 approval → squash and merge → delete branch
```

---

## 6. GitHub Issues & board

- Every non-trivial task starts as a **GitHub Issue**, assigned to one person.
- Use labels: `frontend`, `backend`, `bug`, `feature`, `docs`.
- Use a **Projects** board with columns `Backlog → In Progress → In Review → Done` so all 4 of you can see who's
  doing what without asking in chat.
- If two people might touch the same area, flag it in the issue before starting.

---

## 7. Quick reference cheat sheet

```
Branch:  <type>/<issue#>-<short-description>
Commit:  <type>(<scope>): <summary>
PR:      same format as commit title, 1 approval required, squash & merge
```

Types: `feat fix refactor docs style test chore perf migration`
