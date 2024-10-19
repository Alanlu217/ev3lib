# ev3lib

## Installation

```bash
go get github.com/Alanlu217/ev3lib/ev3lib
go get github.com/Alanlu217/ev3lib/ev3lib/ev3
go get github.com/Alanlu217/ev3lib/ev3lib/testUtils # If doing simulations

```

## Docs
 - [ev3lib](https://pkg.go.dev/github.com/Alanlu217/ev3lib/ev3lib)
 - [ev3](https://pkg.go.dev/github.com/Alanlu217/ev3lib/ev3lib/ev3)
 - [testUtils](https://pkg.go.dev/github.com/Alanlu217/ev3lib/ev3lib/testUtils)

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
