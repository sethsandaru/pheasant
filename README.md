# Pheasant / Pheasant-API

<p align="center">
  <img src="./docs/imgs/phesant-logo.png" title="Pheasant API SethPhat SethSandaru" alt="Pheasant API SethPhat SethSandaru">
</p>

Pheasant is a dynamic CRUD Management which helps end-users to create the CRUD listing/form in seconds then use it with ease and fun.

Pheasant comes with 2 parts:
- API (this repo)
- Frontend (on paper/head)
  - (WIP) still thinking in the middle of Svelte and Vue 🤔 Will make a decision after finished v0.0.1.

Pheasant is an open-source project and planned to deploy on CLOUD with 100% FREE usage. 
But if you don't want PheasantCloud keeping your data, feel free to clone and deploy for your own usage privately.

## Why?
- From high-level/end-users PoV: I want to manage data with ease with Good UI, I'm bored with Excel and stuff.
- From low-level/code-wise: CRUD mostly copy-paste, let's do less of that and Pheasant will help you.

## Dependencies
- Go 1.16+
- PostgreSQL
- Redis (for cache and pub/sub queue)

## Development

Install the dependencies:

```bash
go get
```

Create the `.env` file based on the `.env.example` and add your configuration values there.

Starting the project:
```bash
go run main.go 
# OR
make start-app
```

Run the queue worker:
```bash
go run queue-worker.go
# OR
make start-worker
```

## Deployment

### Build

```bash
go build main.go # BUILD APP
go build queue-worker # Build Worker
# OR
make build-app
make build-worker
```

## Milestones

### v0.0.1
- Basic functionalities
  - Authentication (Login / Register / Forgot Password)
  - Manage Entity
  - Manage Entity's Data

### v0.0.2
- More filtering options

### v0.0.3
- Relationship(s) between entity.

## License

MIT License

## Contributor(s)
- [Seth Phat](https://github.com/sethsandaru)
