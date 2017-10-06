# SIGFUZZ
SIGFUZZ is a signal fuzzer that is capable of delivering signals to multiple
processes at regular intervals.

```sh
$ sigfuzz -h
Usage:                                                                                                                                                                                                                                        
  sigfuzz [flags]

Flags:
  -c, --config string        config file (default is $HOME/.sigfuzz.yaml)
  -x, --exit                 exit on signal failure
  -i, --interval duration    time duration between signals (default 1s)
  -n, --number int           number of times to signal (default 1)
  -p, --pid stringSlice      pids to fuzz
  -s, --signal stringSlice   signal to fuzz
```
