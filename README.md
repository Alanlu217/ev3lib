# ev3lib

## Building Instructions
### For EV3
To build for ev3, set the environment flags GOOS to linux, GOARCH to arm and GOARM to 5.

e.g.

```bash
GOOS=linux GOARCH=arm GOARM=5 go build -o main .
```

### For Testing

To build as a test, run with the tag ev3test

e.g.
```bash
go run -tags ev3test .
```

