# ev3lib

## Building Instructions

To build for ev3, set the environment flags GOOS to linux, GOARCH to arm and GOARM to 5.

e.g.

```bash
GOOS=linux GOARCH=arm GOARM=5 go build -o main .
```