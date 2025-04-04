## Example of genai sdk of google

## Env variable
Please create .env following by
```sh
GENAI_API_KEY=<YOUR_API_KEY>
```

## How to run ?
1. export env from .env
```sh
export $(grep -v '^#' .env | xargs)
```
2. go run
```sh
go run .
```
