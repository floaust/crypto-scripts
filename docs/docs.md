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
|`crypto-scripts gcd`|greatest common divisor|
|`crypto-scripts help`|Help about any command|
|`crypto-scripts sam`|square-and-multiply|
# ... gcd
`crypto-scripts gcd`

## Usage
> greatest common divisor

crypto-scripts [global options] command [options] [arguments...]

## Description

```
The greatest common divisor (GCD) of two nonzero integers n and a is the greatest positive integer d such that d is a divisor of both n and a; that is, there are integers e and f such that n = de and a = df, and d is the largest such integer. The GCD of n and a is generally denoted gcd(n, a).[9]

This definition also applies when one of n and a is zero. In this case, the GCD is the absolute value of the non zero integer: gcd(n, 0) = gcd(0, n) = |n|. This case is important as the terminating step of the Euclidean algorithm.

The above definition cannot be used for defining gcd(0, 0), since 0 × n = 0, and zero thus has no greatest divisor. However, zero is its own greatest divisor if greatest is understood in the context of the divisibility relation, so gcd(0, 0) is commonly defined as 0. This preserves the usual identities for GCD, and in particular Bézout's identity, namely that gcd(n, a) generates the same ideal as {n, a}.[10][11][12] This convention is followed by many computer algebra systems.[13] Nonetheless, some authors leave gcd(0, 0) undefined.[14]

The GCD of n and b is their greatest positive common divisor in the preorder relation of divisibility. This means that the common divisors of n and a are exactly the divisors of their GCD. This is commonly proved by using either Euclid's lemma, the fundamental theorem of arithmetic, or the Euclidean algorithm. This is the meaning of "greatest" that is used for the generalizations of the concept of GCD.
```

## Flags
|Flag|Usage|
|----|-----|
|`-a, --a int`|a|
|`-n, --n int`|n|
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
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 24 May 2021**
