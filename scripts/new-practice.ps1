#requires -Version 7.0
[CmdletBinding()]
param(
    [Parameter(Mandatory = $true)][ValidateRange(1, 28)][int]$Episode,
    [string]$DestinationRoot
)
$ErrorActionPreference = 'Stop'
$courseRoot = Split-Path -Parent $PSScriptRoot
$manifest = Get-Content -LiteralPath (Join-Path $PSScriptRoot 'phase01-checks.json') -Raw | ConvertFrom-Json
$lesson = $manifest | Where-Object episode -EQ $Episode
if (-not $lesson) { throw "Episode $Episode is not available." }
if (-not $DestinationRoot) { $DestinationRoot = Join-Path $courseRoot 'practics/phase-01' }
$target = [IO.Path]::GetFullPath((Join-Path $DestinationRoot $lesson.practice))
if (Test-Path -LiteralPath $target) {
    throw "Practice folder already exists. Your work was not changed: $target"
}
$source = Join-Path $courseRoot $lesson.path
foreach ($item in $lesson.copies) {
    if (-not (Test-Path -LiteralPath (Join-Path $source $item))) { throw "Missing source: $item" }
}
New-Item -ItemType Directory -Path $target -ErrorAction Stop | Out-Null
foreach ($item in $lesson.copies) {
    Copy-Item -LiteralPath (Join-Path $source $item) -Destination (Join-Path $target $item) -Recurse -ErrorAction Stop
}
Write-Output "Practice ready: $target"
Write-Output "Open that folder, then run: go run ./$($lesson.entry)"
