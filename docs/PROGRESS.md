# Progress and review log

Learner: Micah Philpotts
Project: delivery-lab
Current level: 02
Status: REVISE — Level 02 in progress
Repository URL: https://github.com/thebasedtaka/delivery-lab
Approved levels: 01

## Submission template

- Level / ticket:
- Repository / PR:
- Workflow YAML (path + content or accessible link):
- Initial green run URL + commit SHA:
- Intentional failure run URL + commit SHA:
- Repaired green run URL + commit SHA:
- Relevant logs (if links cannot be accessed):
- What failed and why:
- How I fixed it:
- Answers to the level's review questions:
- Any constraints or requested hint:

## Review history

### Level 01 — PASS

Decision made in chat on November 19, 2026; recorded here on November 20, 2026.

Reviewed the workflow configuration, run steps, and failure/restoration code changes:
- [Run #7](https://github.com/thebasedtaka/delivery-lab/actions/runs/44649346760), commit `3cc1531d6fa25b17b0b358399e3f4894c153ff52`: intentional greeting regression caused tests to fail and the build to be skipped.
- [Run #8](https://github.com/thebasedtaka/delivery-lab/actions/runs/44649434824), commit `db87d3cc7a8f800a9ba4b35457d396b8515dbd44`: greeting restored; tests and build passed.
- [Run #9](https://github.com/thebasedtaka/delivery-lab/actions/runs/44651482516): pull-request trigger verified; tests and build passed.
- [PR #1](https://github.com/thebasedtaka/delivery-lab/pull/1): merged into main.
- [Run #10](https://github.com/thebasedtaka/delivery-lab/actions/runs/44651592944): post-merge main run passed.

Verified explicit Go setup, tests before build, and read-only contents permissions.
The reviewer accepted the combined failure/recovery evidence on main and successful PR execution as sufficient for Level 01. Repeating the same failure drill on a PR is not required. The observed PR run was successful; this record does not claim a failed PR run.

Concept feedback covered parallel jobs versus sequential steps, checkout versus toolchain setup, nonzero exit codes, and skipped downstream steps.

## How to resume

Bring this file, the current ticket, and your current workflow into a new session.
Say: "Act as my Delivery Lab project manager. Review my current level before
unlocking the next one." Update this record after each review so progress is
explicit rather than dependent on chat memory.

## Next assignment

Level 02 is unlocked at [LEVEL-02.md](LEVEL-02.md). It introduces
formatting and vet quality gates, dependent jobs, and required-check research.
Level 02 has been reviewed and needs revision; it has not passed.

## Level 02 review — November 20, 2026 — REVISE

Reviewed workflow at `d803b3f874c3270b96a143249e7b9bf4c909e51e` and runs [#20](https://github.com/thebasedtaka/delivery-lab/actions/runs/44848484365) and [#21](https://github.com/thebasedtaka/delivery-lab/actions/runs/44848656747).

- Three jobs are correctly separated: quality and test run independently; build needs both. Each checks out code, sets Go 1.24, and uses contents: read.
- Run #20: formatting step failed, test passed, build skipped. Workflow used `go fmt -l ./...`, an invalid invocation; this is not evidence of rejecting misformatted source.
- Run #21: all three jobs passed. Formatting step now uses `go fmt ./...`, which rewrites files rather than rejecting formatting differences. Vet remains present and passed.
- Next: use a non-mutating formatting check and explicitly fail when files need formatting; demonstrate a formatting-only failure and repair.
- The three-job layout clarified in chat is authoritative over the older two-job layout in LEVEL-02.md. Do not require restructuring back to two jobs.

### Level 02 follow-up — REVISE

[Run #24](https://github.com/thebasedtaka/delivery-lab/actions/runs/44850007097), commit `9b2d87afcd22c1c68e22123b11a12835d3c0fc28`: quality, test, and build passed. Command corrected to `gofmt -l .`; it does not rewrite files. The step still lacks a shell condition to exit nonzero when the output lists files. Next: add that condition, then demonstrate formatting-only failure and repair. Three-job structure and vet remain correct.
