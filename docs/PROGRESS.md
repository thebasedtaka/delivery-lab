# Progress and review log

Learner: Micah Philpotts
Project: delivery-lab
Current level: 03
Status: READY — Level 01 passed; Level 02 unlocked
Repository URL: https://github.com/thebasedtaka/delivery-lab
Approved levels: 01, 02

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
Level 02 has not yet been reviewed or passed.

### Level 02 — PASS (November 20, 2026)

Formatting drill accepted in runs #29/#30 above. Vet drill verified: [#35](https://github.com/thebasedtaka/delivery-lab/actions/runs/44854527668) passed formatting, failed vet on a printf verb/type mismatch, failed test, and skipped build; [#36](https://github.com/thebasedtaka/delivery-lab/actions/runs/44854569570) repaired the verb and passed all jobs. [#37](https://github.com/thebasedtaka/delivery-lab/actions/runs/44854647885) also passed.

[Basic Rule](https://github.com/thebasedtaka/delivery-lab/rules/27777735) verified active, targeting only refs/heads/main, requiring PRs and quality/test/build checks, with zero required approvals and no bypass actors. Configuration evidence accepted; no additional failure drill is required.

Learner's six review answers in LEVEL-02.md accepted. Clarifications: a dependent job waits for its prerequisites to succeed; each runner needs both source checkout and the selected Go toolchain. Protection blocks merges that violate its requirements, not all manual merges. CI formatting checks ensure the committed source itself satisfies formatting requirements.

Level 02 complete. Level 03 unlocked, assignment pending. No repeat Level 02 drills required.
