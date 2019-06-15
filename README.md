# Raptime
`Raptime` is a command to add the input time to the prefix of standard input. When standard input is done for this command, it is usually passed via a (anonymous) pipe.

## Installation

```
$ go get -u github.com/cia-rana/raptime
```

## Usage

```
$ <preceding-command> | raptime
```

If you want to add time to the standard error output as well, please execute as follows.

```
$ <preceding-command> 2>&1 | raptime
```
