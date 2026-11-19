# Delivery Lab progression

Only Level 01 is active. Later levels are a backlog, not simultaneous homework.
Each level gets its own client brief, acceptance criteria, incident drill, and
review when unlocked. We increase one major source of complexity at a time.

GitHub Actions is extensible, so no fixed course can exhaust every possible
integration or use case. Our target is broad practical coverage of its major
capabilities and the judgement needed to build new workflows independently.
Recheck current documentation and account/plan support when each level starts.

| Level | Client request / capability |
| --- | --- |
| 01 | Run Go tests and build on PRs and main; prove failure detection. |
| 02 | Add formatting and vet checks; job dependencies; enforce required checks where available. |
| 03 | Learn contexts, expressions, variables, env scope, step/job outputs and conditions. |
| 04 | Cancel superseded PR work; timeouts, concurrency groups and cancellation behaviour. |
| 05 | Test Go versions and operating systems; matrix include/exclude and fail-fast. |
| 06 | Retain test reports and coverage after failures; artifacts, retention and job summaries. |
| 07 | Introduce a real dependency; cache keys, restore strategies and cache invalidation. |
| 08 | Add PostgreSQL integration tests; service containers, health checks and job containers. |
| 09 | Build and smoke-test Docker images; publish to GHCR with scoped permissions. |
| 10 | Release tags, cross-platform binaries, checksums, release notes and GitHub Releases. |
| 11 | Manual dispatch with typed inputs; scheduled maintenance and repository dispatch. |
| 12 | Reusable workflows and composite actions; contracts, outputs and secret boundaries. |
| 13 | Monorepo change detection; dynamic matrices, shared-code changes and selective execution. |
| 14 | GitHub environments, approvals where supported, configuration and deployment protection. |
| 15 | AWS OIDC; least privilege, scoped trust and negative authorization tests. |
| 16 | Terraform fmt/validate/plan/apply; state locking, reviewable plans and safe concurrency. |
| 17 | Build once; promote the same image digest through staging and production. |
| 18 | Smoke failures, rollback, reruns, idempotency and interrupted-deployment recovery. |
| 19 | Fork PRs, Dependabot, token permissions, script injection, action pinning and supply-chain checks. |
| 20 | SBOMs, provenance/attestations where supported, verification and vulnerability policy. |
| 21 | Custom JavaScript and Docker actions; action metadata, inputs, outputs and testing. |
| 22 | GitHub API/CLI automation; issues, labels, PR events and permission boundaries. |
| 23 | Cross-workflow and cross-repository orchestration; workflow_run trust and artifact integrity. |
| 24 | Self-hosted/ephemeral runners; labels, groups, isolation and cleanup; larger runners/ARC if appropriate. |
| 25 | Operational controls: quotas, performance, billing, retention, auditability and organization standards. |
| 26 | Capstone: operate a multi-service delivery platform under injected incidents. |

Some features depend on repository visibility, account plan, or organization
settings. Use an equivalent simulated gate where necessary and document the
difference. Paid/cloud steps require a budget decision at that level. No cloud
resources or paid runners are provisioned as part of this starter project.

The capstone ends with a coverage review and follow-up tickets for gaps,
including new GitHub capabilities introduced while you learn.
