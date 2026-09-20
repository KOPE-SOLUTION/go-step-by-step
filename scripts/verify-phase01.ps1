#requires -Version 7.0
[CmdletBinding()]
param(
    [ValidateRange(0, 28)][int]$Episode = 0,
    [string]$GoCommand = 'go'
)
$ErrorActionPreference = 'Stop'
$courseRoot = Split-Path -Parent $PSScriptRoot
$manifest = @(Get-Content -LiteralPath (Join-Path $PSScriptRoot 'phase01-checks.json') -Raw | ConvertFrom-Json)
if ($Episode -ne 0) { $manifest = @($manifest | Where-Object episode -EQ $Episode) }
$goPath = (Get-Command $GoCommand -ErrorAction Stop).Source
$gofmtPath = Join-Path (Split-Path -Parent $goPath) 'gofmt.exe'
if (-not (Test-Path -LiteralPath $gofmtPath)) { $gofmtPath = (Get-Command gofmt -ErrorAction Stop).Source }
function Invoke-Tool([string]$File, [string[]]$ToolArguments, [string]$Directory) {
    $info = [Diagnostics.ProcessStartInfo]::new()
    $info.FileName = $File
    $info.WorkingDirectory = $Directory
    $info.UseShellExecute = $false
    $info.CreateNoWindow = $true
    $info.RedirectStandardOutput = $true
    $info.RedirectStandardError = $true
    $info.StandardOutputEncoding = [Text.Encoding]::UTF8
    $info.StandardErrorEncoding = [Text.Encoding]::UTF8
    foreach ($argument in $ToolArguments) { $info.ArgumentList.Add($argument) }
    $process = [Diagnostics.Process]::new()
    $process.StartInfo = $info
    $null = $process.Start()
    $stdoutTask = $process.StandardOutput.ReadToEndAsync()
    $stderrTask = $process.StandardError.ReadToEndAsync()
    $process.WaitForExit()
    $result = [pscustomobject]@{
        ExitCode = $process.ExitCode
        Stdout = $stdoutTask.GetAwaiter().GetResult()
        Stderr = $stderrTask.GetAwaiter().GetResult()
    }
    $process.Dispose()
    return $result
}
function Assert-Success($Result, [string]$Label) {
    if ($Result.ExitCode -ne 0) { throw "$Label failed ($($Result.ExitCode)): $($Result.Stdout)$($Result.Stderr)" }
}
function Normalize-Output([string]$Value) {
    return $Value.Replace(([string][char]13 + [char]10), [string][char]10).TrimEnd([char[]]@(13,10))
}
$previousToolchain = $env:GOTOOLCHAIN
$previousProxy = $env:GOPROXY
try {
    $env:GOTOOLCHAIN = 'local'
    $env:GOPROXY = 'off'
    $version = Invoke-Tool $goPath @('version') $courseRoot
    Assert-Success $version 'go version'
    Write-Output $version.Stdout.Trim()
    $runCount = 0
    foreach ($lesson in $manifest) {
        $directory = Join-Path $courseRoot $lesson.path
        $files = @(Get-ChildItem -LiteralPath $directory -Recurse -File -Filter '*.go' | ForEach-Object FullName)
        $formatted = Invoke-Tool $gofmtPath (@('-l') + $files) $directory
        Assert-Success $formatted 'gofmt'
        if ($formatted.Stdout.Trim()) { throw "EP$($lesson.episode): files need gofmt: $($formatted.Stdout)" }
        $test = Invoke-Tool $goPath @('test', '-count=1', './...') $directory
        Assert-Success $test "EP$($lesson.episode) go test"
        foreach ($case in $lesson.cases) {
            $run = Invoke-Tool $goPath @('run', $case.path) $directory
            Assert-Success $run "EP$($lesson.episode) $($case.path)"
            $expected = [string]::Join([string][char]10, [string[]]$case.expected)
            if ((Normalize-Output $run.Stdout) -cne $expected -or $run.Stderr.Length -gt 0) {
                throw "EP$($lesson.episode) $($case.path): output differs. Expected: $expected Actual: $($run.Stdout)$($run.Stderr)"
            }
            $runCount++
        }
        if ($lesson.PSObject.Properties.Name -contains 'negative') {
            $tempRoot = Join-Path ([IO.Path]::GetTempPath()) ('go-basic-negative-' + [guid]::NewGuid())
            New-Item -ItemType Directory -Path $tempRoot | Out-Null
            Copy-Item -LiteralPath (Join-Path $directory $lesson.negative.file) -Destination (Join-Path $tempRoot 'main.go')
            $negative = Invoke-Tool $goPath @('run', 'main.go') $tempRoot
            if ($negative.ExitCode -eq 0 -or -not $negative.Stderr.Contains($lesson.negative.contains)) {
                throw "EP$($lesson.episode): expected diagnostic was not found: $($negative.Stderr)"
            }
            # Keep this small temporary fixture so a failed run remains inspectable.
        }
        $kind = if ($lesson.unit) { 'unit tests + outputs' } else { 'compile + outputs (no unit tests)' }
        Write-Output "PASS EP$($lesson.episode): $kind"
    }
    Write-Output "PASS: $($manifest.Count) episodes; $runCount program outputs matched; gofmt clean."
} finally {
    $env:GOTOOLCHAIN = $previousToolchain
    $env:GOPROXY = $previousProxy
}
