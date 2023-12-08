## go cli for quick cognito user pool management

- only list users implemented
- requires `.env` file in root dir for context to authorize requests
  - keys inside env: AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, USRPOOL, CLNTID
- to run: `go build && go install && ldc list` 