# crypto-scripts

## Usage
> Outputs ggt, sam and ss with a nice output path

crypto-scripts [global options] command [options] [arguments...]

## Description

```
A small tool to output

- ggt (greatest common divisor)
- sam (square-and-multiply)
- ss (shift cipher for english and german texts)

in a nice way with the computing path.

Nice output designed with [PTerm](https://github.com/pterm/pterm).
```

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`crypto-scripts help`|Help about any command|
|`crypto-scripts sam`|square-and-multiply|
# ... help
`crypto-scripts help`

## Usage
> Help about any command

crypto-scripts [global options] command [options] [arguments...]

## Description

```
Help provides help for any command in the application.
Simply type crypto-scripts help [path to command] for full details.
```
# ... sam
`crypto-scripts sam`

## Usage
> square-and-multiply

crypto-scripts [global options] command [options] [arguments...]

## Description

```
Calculates the square-and-multiply algorithm.
```

## Flags
|Flag|Usage|
|----|-----|
|`-b, --base int`|base|
|`-e, --exp int`|exponent|
|`-m, --mod int`|modulo|


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 23 May 2021**
