## go cli for quick cognito user pool management

- list users with permissions and modify permissions implemented
- requires `.env` file in root dir for context to authorize requests
  - keys inside env: AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, USRPOOL, CLNTID
- to run: `go build && go install && ldc list` 

## todo:

- DONE add user to group in pool
- DONE listusers in group (there's a function in the sdk)
- DONE list groups in pool