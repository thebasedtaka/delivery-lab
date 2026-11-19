# Delivery Lab — GitHub Actions learning project

A small Go HTTP API for Micah's progressively harder delivery assignments.
The application is supplied; implementing GitHub Actions is your assignment.
There is deliberately no completed workflow, Dockerfile, or Terraform yet.

## Start here

1. Install a supported Go release compatible with `go.mod` (minimum 1.24).
2. Extract this directory and open a terminal **inside delivery-lab**, where `go.mod` lives.
3. Verify your tools and the baseline:

```sh
go version
go test -v ./...
go build ./...
go run ./cmd/api
```

The last command keeps running. Use another terminal to open:

- http://127.0.0.1:8080/health → `{"status":"ok"}`
- http://127.0.0.1:8080/version → `{"version":"dev"}`
- http://127.0.0.1:8080/greet?name=Micah → `{"message":"Hello, Micah!"}`

Stop the server with Ctrl+C. The tests do not need a running server.
Optional: set the `ADDR` environment variable to change the listening address.
The loopback default is for local practice; container networking comes later.

## Create your repository

Create an empty GitHub repository called `delivery-lab`, then run:

```sh
git init -b main
git add .
git commit -m "Add Go learning project"
git remote add origin https://github.com/YOUR-USERNAME/delivery-lab.git
git push -u origin main
git switch -c level-01-ci
```

Replace YOUR-USERNAME first. Use your normal GitHub authentication; never
place tokens in source files. Push the new branch after adding your workflow,
then open a pull request into main. A public learning repository is convenient
if you are comfortable sharing this sample code; a private repository also works,
subject to your account's Actions allowance and feature availability.

## Project map

- `cmd/api/main.go`: starts the HTTP server.
- `internal/api/handler.go`: routes and the greeting function.
- `internal/api/handler_test.go`: unit and HTTP handler tests.
- `docs/LEVEL-01.md`: your active ticket and exact acceptance criteria.
- `docs/ROADMAP.md`: future levels, unlocked after review.
- `docs/PROGRESS.md`: progress record and submission template.
- `AGENTS.md`: coaching rules for future sessions.

No third-party packages are used, so there is no `go.sum`. Do not invent one.
If your Go setup action attempts dependency caching, disable it for this first
level; caching becomes a separate exercise once dependencies are introduced.

## Reference material

Use documentation freely. Try writing the workflow before copying an example.

- https://docs.github.com/en/actions/tutorials/build-and-test-code/go
- https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax
- https://go.dev/doc/install

Begin with **docs/LEVEL-01.md**. The project manager reviews evidence; a green
badge alone is not enough to establish that failures are detected correctly.
