# Level 02 — Quality gates and dependent jobs

**Status:** PASS — reviewed November 20, 2026  
**Estimated effort:** 60–120 minutes  
**Cloud spend:** None

## Client brief

Delivery Lab's Go API now has automated tests and a build. However, reviewers
still have to spot formatting mistakes and suspicious code manually. We need
quality checks that fail appropriately and prevent unnecessary downstream work.

You own the CI design. Do not change the API's intended behaviour.

## Your assignment

Extend the existing `.github/workflows/go.yml` workflow instead of creating a
duplicate pipeline. Keep the existing triggers for pushes to `main` and pull
requests targeting `main`.

Introduce **three separate jobs**:

1. **Quality gate:** Verify Go formatting without modifying files and run
   static analysis with `go vet ./...`.
2. **Test:** Run `go test -v ./...` independently of quality.
3. **Build:** Run `go build ./...` only after both quality and test succeed.

Name jobs and steps clearly. Explicitly install/select a Go version compatible
with `go.mod` wherever you run Go commands. Remember that GitHub-hosted jobs
use independent runners.

Keep this exercise focused on jobs, dependencies, and quality checks. No Docker,
AWS, Terraform, external dependencies, secrets, or deployment are required.

## Acceptance criteria

- [ ] Workflow still triggers on pushes to `main` and PRs targeting `main`.
- [ ] Formatting verification catches misformatted Go source with a nonzero exit
      status **without rewriting source files in CI**.
- [ ] The quality job runs `go vet ./...` and fails if vet reports an issue.
- [ ] Tests and build execute in a different job from the quality checks.
- [ ] An explicit job dependency prevents the build job from starting until
      both quality and test succeed.
- [ ] If quality fails, the dependent job is visibly skipped.
- [ ] Tests and build still fail when their commands fail.
- [ ] No job requires write access or repository secrets.
- [ ] A repaired commit makes all three jobs pass, with the Go API unchanged.
- [ ] Investigate branch protection or rulesets. If required status checks are
      available for this repository, configure them for merges to `main`.
      If unavailable, document the limitation and explain the difference
      between a green CI check and enforced merge protection.

## Required failure drills

### Drill A — Formatting

1. Capture a successful workflow run before the drill.
2. On a feature branch, deliberately misformat one tracked Go file without
   changing its behaviour.
3. Push the commit and capture the formatting failure.
4. Confirm the dependent build job was skipped.
5. Restore formatting, push a new commit, and capture the successful run.

### Drill B — Static analysis

1. Introduce a temporary Go code issue that `go vet` reports. Preserve the
   existing test assertions.
2. Commit and push. Capture the failed vet step and skipped dependent job.
3. Fix the issue without disabling vet; push and capture a successful run.

Do not hide failure exit codes, weaken tests, add `continue-on-error`, or use
`|| true` to make the checks look green.

## Submission for review

Send the PR URL, final workflow YAML, and the following evidence:

- Initial green run URL and commit SHA.
- Formatting-failure run URL and commit SHA, followed by repair run and SHA.
- Vet-failure run URL and commit SHA, followed by repair run and SHA.
- Evidence the dependent job was skipped for each quality failure.
- Branch-protection/ruleset outcome: configured checks or an explanation of
  why this is not supported for the current repository.

Use the submission format in [PROGRESS.md](PROGRESS.md).

## Review questions — answer in your own words

1. What is the difference between steps executing sequentially and two jobs
   connected by an explicit dependency?
   the latter will only run if the dependent job has ran successfully
2. Why does each job need its own checkout and Go setup?
   the jobs run on different virtual machines so each of the machines need a copy of their own code
3. Why should CI **check** formatting rather than automatically fix it?
   if its automatically fixed the client side code never fixes the issues and can cause drift
4. What kind of problem does `go vet` catch that ordinary unit tests might not?
   go vet catches mistakes such as variables misused that a compiler would still mark as non problematic
5. What happen to dependent jobs when a prerequisite fails?
   the jobs do not execute
6. Why doesn't a passing CI run necessarily stop someone from merging a failing
   PR? Which repository setting changes that?
   branch protection prevents manual merges, without that it is up to the discression of the user

## Review gate

I will return **PASS**, **REVISE**, or **BLOCKED**, based on the YAML, actual
workflow-run evidence, failure drills, and merge-check investigation. A green
badge alone is not sufficient.

I will not write the workflow for you unless you request the solution.
Ask for **Hint 1** (concept), **Hint 2** (documentation/keywords), or **Hint 3**
(partial scaffold).

## Official references

- [Workflow syntax](https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax)
- [Using jobs](https://docs.github.com/en/actions/how-tos/write-workflows/choose-what-workflows-do/use-jobs)
- [Protected branches](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/about-protected-branches)
- [go vet](https://pkg.go.dev/cmd/vet)
- [gofmt](https://pkg.go.dev/cmd/gofmt)

## Final review — PASS

All implementation, drill, protection configuration, and concept review requirements accepted; see PROGRESS.md for evidence. Original learner answers above are preserved. Reviewer clarifications: prerequisite jobs must succeed before the dependent build runs; Go setup selects the toolchain on each runner; branch protection blocks noncompliant merges rather than all manual merges. No additional Level 02 drills required.
