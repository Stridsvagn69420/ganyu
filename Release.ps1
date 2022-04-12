#!/usr/bin/pwsh
$rawArray = $(go tool dist list).Split([Environment]::NewLine)
$releaseText = ""
$PWD = (Get-Location).Path
foreach ($osArch in $rawArray) {
    # Env
    $osArchArr = $osArch.Split("/")
    $env:GOOS = $osArchArr[0]
    $env:GOARCH = $osArchArr[1]
    # Compile
    $filename = "ganyu-$env:GOOS-$env:GOARCH"
    if ($env:GOOS -eq "windows") {
        $filename += ".exe"
    }
    (go build -o ./build/$filename -ldflags="-s -w" ganyu.go config.go distro.go meta.go commands.go help.go) | Out-Null
    # Release
    if (Test-Path ./build/$filename -PathType Leaf) {
        $hash = (Get-FileHash -Algorithm SHA256 -LiteralPath "./build/$filename").Hash.ToLower()
        $releaseText += "* ``$filename``: $hash`n"
    }
}
Clear-Host
Write-Output $releaseText