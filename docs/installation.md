# Quick Start - Install crypto-scripts

> [!TIP]
> crypto-scripts is installable via [instl.sh](https://instl.sh).\
> You just have to run the following command and you're ready to go!

<!-- tabs:start -->

#### ** Windows **

### Windows Command

```powershell
iwr -useb instl.sh/floaust/crypto-scripts/windows | iex
```

#### ** Linux **

### Linux Command

```bash
curl -fsSL instl.sh/floaust/crypto-scripts/linux | bash
```

#### ** macOS **

### macOS Command

```bash
curl -fsSL instl.sh/floaust/crypto-scripts/macos | bash
```

#### ** Compile from source **

### Compile from source with Golang

?> **NOTICE**
To compile crypto-scripts from source, you have to have [Go](https://golang.org/) installed.

Compiling crypto-scripts from source has the benefit that the build command is the same on every platform.\
It is not recommended to install Go only for the installation of crypto-scripts.

```command
go install github.com/floaust/crypto-scripts@latest
```

<!-- tabs:end -->
