# httpepe.back

Go backend part of httpepe project.

## Build & Run

Using `Taskfile.yml` <https://taskfile.dev/>

> Build

```sh
❯ task build
task: [build] go build -o bin/httpepe main.go
```

> Run dev

```sh
❯ task run-dev
task: [run-dev] go run main.go
```

> Run with binary

```sh
❯ task run
task: [build] go build -o bin/httpepe main.go
task: [run] bin/httpepe
```

## API Calls

> Query

```sh
❯ curl http://127.0.0.1:3000/api/codes | jq
```

> Answer

```json
{
  "data": [
    {
      "id": 100,
      "message": "Continue"
    },
    // truncated
    {
      "id": 599,
      "message": "Network Connect Timeout Error"
    }
  ]
}
```
