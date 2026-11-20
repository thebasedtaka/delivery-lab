# Level 01 — Make every change test itself

Status: READY · Suggested effort: 45–90 minutes · Cloud spend: none required

## Client brief

We have a working Go API, but nobody checks changes consistently. Before the
team adds features, we need GitHub to test and build every proposed change.
You own the automation. Keep the existing application behaviour intact.

## Scope

Create `.github/workflows/ci.yml` on your `level-01-ci` branch. Implement one
job on a GitHub-hosted Linux runner, with clearly named steps. No deployment,
Docker, AWS, secrets, matrix, caching, or reusable workflows are required yet.

## Acceptance criteria

- [ ] The workflow runs on pushes to `main`.
- [ ] It runs when a pull request targeting `main` is opened or updated.
- [ ] It checks out the repository and explicitly installs/selects Go using a setup action.
- [ ] The selected Go version is compatible with `go.mod` and not left to runner defaults.
- [ ] It executes the real test suite with `go test -v ./...`.
- [ ] It executes `go build ./...` after tests pass.
- [ ] A failing test causes the job/run to fail; build is skipped after that failure.
- [ ] A subsequent commit fixing the defect makes the PR run pass again.
- [ ] The workflow requests only read access to repository contents; no credentials are needed.
- [ ] The final submitted application and tests retain the starter behaviour.

## Required failure drill

1. First demonstrate a successful run.
2. In `Greeting`, temporarily change the empty-name fallback from `friend` to `stranger`.
3. Commit and push that change on the PR branch. Do not change the tests.
4. Confirm that GitHub reports failure and identify the failing test cases.
5. Restore `friend`, commit, push, and confirm that the new run succeeds.

Do not suppress errors, delete assertions, use `continue-on-error`, or add
`|| true` to make the check green. The defect must be caught, then fixed.

## Submit for review

Send your workflow YAML, the PR link, and evidence for the initial green run,
failed run, and repaired green run. Paste relevant logs if the repository is
private or I cannot access the run links. Include the commit SHA for each run.
Use the template in PROGRESS.md.

Answer in your own words:
1. What is the difference between a workflow, a job, and a step?
    a workflow is triggered by an event 
    a job is work triggered in sequence in a workflow
    a step is an action inside the job
2. Why does a fresh runner need checkout and Go setup?
    github runs the tests in a remote server and the server needs to download from your code base before any work is done
3. What told GitHub the test step failed?
    an error will trigger github
4. Why was the build step skipped on the failing run?
    because if one step fails the steps after will fail
## Review gate

I will return PASS, REVISE, or BLOCKED, with concrete evidence and next actions.
PASS requires all acceptance criteria and the failure drill. A workflow that
runs on PRs does not itself prevent merging; enforced merge checks come later.
After approval, merge the PR, verify the main-branch run, and start Level 02.

## Hints on request

Ask for Hint 1 (concept), Hint 2 (documentation/keywords), or Hint 3 (partial
scaffold). A complete workflow solution is available if explicitly requested.
