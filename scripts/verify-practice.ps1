#requires -Version 7.0
[CmdletBinding()]
param([string]$GoCommand = 'go')
$ErrorActionPreference = 'Stop'
$courseRoot = Split-Path -Parent $PSScriptRoot
$manifest = @(Get-Content -LiteralPath (Join-Path $PSScriptRoot 'phase01-checks.json') -Raw | ConvertFrom-Json)
$goPath = (Get-Command $GoCommand -ErrorAction Stop).Source
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
$previousWork = $env:GOWORK
$testRoot = Join-Path ([IO.Path]::GetTempPath()) ('go-course-practice-' + [guid]::NewGuid().ToString('N'))
$practiceRoot = Join-Path $testRoot 'practics'
New-Item -ItemType Directory -Path $practiceRoot | Out-Null
$encoding = [Text.UTF8Encoding]::new($false)
try {
    $env:GOTOOLCHAIN = 'local'
    $env:GOPROXY = 'off'
    $env:GOWORK = 'off'
    foreach ($lesson in $manifest) {
        $sourceRoot = Join-Path $courseRoot $lesson.path
        $main = [IO.File]::ReadAllText((Join-Path $sourceRoot 'main.go'))
        if ($lesson.episode -eq 12) {
            $init = Invoke-Tool $goPath @('mod', 'init', 'example.com/go-practice') $practiceRoot
            Assert-Success $init 'practice go mod init'
            $main = $main.Replace('example.com/go-course/phase01/ep12-packages-modules/sensor', 'example.com/go-practice/sensor')
            $sensorRoot = Join-Path $practiceRoot 'sensor'
            New-Item -ItemType Directory -Path $sensorRoot | Out-Null
            Copy-Item -LiteralPath (Join-Path $sourceRoot 'sensor/reading.go') -Destination (Join-Path $sensorRoot 'reading.go')
        }
        [IO.File]::WriteAllText((Join-Path $practiceRoot 'main.go'), $main, $encoding)
        if ($lesson.episode -ge 13) {
            Copy-Item -LiteralPath (Join-Path $sourceRoot 'main_test.go') -Destination (Join-Path $practiceRoot 'main_test.go')
            $test = Invoke-Tool $goPath @('test', '-count=1', '.') $practiceRoot
            Assert-Success $test "EP$($lesson.episode) practice tests"
        }
        $entry = if ($lesson.episode -le 11) { 'main.go' } else { '.' }
        $run = Invoke-Tool $goPath @('run', $entry) $practiceRoot
        Assert-Success $run "EP$($lesson.episode) practice run"
        $expected = [string]::Join([string][char]10, [string[]]$lesson.cases[0].expected)
        if ((Normalize-Output $run.Stdout) -cne $expected -or $run.Stderr.Length -gt 0) {
            throw "EP$($lesson.episode) practice output mismatch: $($run.Stdout)$($run.Stderr)"
        }
        Write-Output "PASS EP$($lesson.episode): shared practice workflow"
    }
    $negativeRoot = Join-Path $testRoot 'typo'
    New-Item -ItemType Directory -Path $negativeRoot | Out-Null
    $hello = [IO.File]::ReadAllText((Join-Path $courseRoot ($manifest[0].path + '/main.go')))
    [IO.File]::WriteAllText((Join-Path $negativeRoot 'main.go'), $hello.Replace('fmt.Println', 'fmt.PrintIn'), $encoding)
    $negative = Invoke-Tool $goPath @('run', 'main.go') $negativeRoot
    if ($negative.ExitCode -eq 0 -or -not $negative.Stderr.Contains('undefined: fmt.PrintIn')) {
        throw "Expected spelling diagnostic: $($negative.Stdout)$($negative.Stderr)"
    }
    Write-Output 'PASS: EP1 spelling mistake produced the expected diagnostic'
    $mutationRoot = Join-Path $testRoot 'threshold-mutation'
    New-Item -ItemType Directory -Path $mutationRoot | Out-Null
    $unitRoot = Join-Path $courseRoot ($manifest | Where-Object episode -EQ 13).path
    $unitMain = [IO.File]::ReadAllText((Join-Path $unitRoot 'main.go'))
    [IO.File]::WriteAllText((Join-Path $mutationRoot 'main.go'), $unitMain.Replace('celsius >= threshold', 'celsius > threshold'), $encoding)
    [IO.File]::WriteAllText((Join-Path $mutationRoot 'go.mod'), "module example.com/threshold-mutation`n`ngo 1.22.0`n", $encoding)
    Copy-Item -LiteralPath (Join-Path $unitRoot 'main_test.go') -Destination (Join-Path $mutationRoot 'main_test.go')
    $mutation = Invoke-Tool $goPath @('test', '-count=1', '-v', '.') $mutationRoot
    if ($mutation.ExitCode -eq 0 -or -not $mutation.Stdout.Contains('FAIL: TestStatus/equal')) {
        throw "Boundary test did not catch the intended bug: $($mutation.Stdout)$($mutation.Stderr)"
    }
    Write-Output 'PASS: EP13 boundary test caught >= changed to >'
    Write-Output "PASS: 14 practice episodes; original practics untouched. Temporary workspace: $testRoot"
} finally {
    $env:GOTOOLCHAIN = $previousToolchain
    $env:GOPROXY = $previousProxy
    $env:GOWORK = $previousWork
}
