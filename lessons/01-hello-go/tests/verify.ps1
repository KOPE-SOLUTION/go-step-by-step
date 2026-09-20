# Run from PowerShell after installing Go. No system settings are changed.
$ErrorActionPreference = 'Stop'
$lessonRoot = Split-Path -Parent $PSScriptRoot

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    throw 'Go is not available. Follow notes/setup-windows.md, reopen PowerShell, then retry. No Go checks were run.'
}
if (-not (Get-Command gofmt -ErrorAction SilentlyContinue)) {
    throw 'gofmt is not available. Check the Go installation. No Go checks were run.'
}

function Assert-ProgramOutput {
    param([string]$ProgramPath, [string[]]$ExpectedLines)
    $actualLines = @(& go run $ProgramPath)
    if ($LASTEXITCODE -ne 0) { throw "go run failed: $ProgramPath" }
    if (($actualLines.Count -ne $ExpectedLines.Count) -or
        (($actualLines -join "`n") -cne ($ExpectedLines -join "`n"))) {
        throw "Output mismatch for $ProgramPath. Expected [$($ExpectedLines -join ' | ')], got [$($actualLines -join ' | ')]."
    }
    Write-Output "PASS output: $ProgramPath"
}

Push-Location -LiteralPath $lessonRoot
try {
    & go version
    if ($LASTEXITCODE -ne 0) { throw 'go version failed.' }

    $formatIssues = @(& gofmt -l .)
    if ($LASTEXITCODE -ne 0) { throw 'gofmt could not parse the source.' }
    if ($formatIssues.Count -gt 0) {
        throw "Formatting differs: $($formatIssues -join ', '). Review before formatting your file."
    }
    Write-Output 'PASS formatting (no files modified)'

    & go test ./...
    if ($LASTEXITCODE -ne 0) { throw 'Package compilation failed. Restore the intentional exercise error before retrying.' }
    Write-Output 'PASS package compilation; this lesson has no unit test functions.'

    & go vet ./...
    if ($LASTEXITCODE -ne 0) { throw 'go vet reported a problem.' }
    Write-Output 'PASS go vet'

    Assert-ProgramOutput './examples/hello' @('Hello, Go!')
    Assert-ProgramOutput './solutions/01-greeting' @('Hello, learner!')
    Assert-ProgramOutput './solutions/02-two-lines' @('Hello, learner!', 'Gateway simulator is starting.')

    # A fresh directory on each run keeps any previous executable intact.
    # Use Windows TEMP: this machine blocks program-created files under Documents.
    $runDirectory = Join-Path ([IO.Path]::GetTempPath()) ('go-course-verify-' + [guid]::NewGuid().ToString('N'))
    New-Item -ItemType Directory -Path $runDirectory | Out-Null
    $executablePath = Join-Path $runDirectory 'hello.exe'
    & go build -o $executablePath ./examples/hello
    if ($LASTEXITCODE -ne 0) { throw 'go build failed.' }
    $builtOutput = @(& $executablePath)
    if (($LASTEXITCODE -ne 0) -or ($builtOutput.Count -ne 1) -or ($builtOutput[0] -cne 'Hello, Go!')) {
        throw 'The built executable did not produce the expected output.'
    }
    Write-Output "PASS built executable: $executablePath"
    Write-Output 'All checks passed. Record the actual version and results in tests/RESULTS.md.'
}
finally {
    Pop-Location
}
