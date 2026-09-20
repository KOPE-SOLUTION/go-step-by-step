#requires -Version 7.0
[CmdletBinding()]
param(
    [ValidateRange(1, 28)][int]$Episode = 1,
    [string]$DestinationRoot
)

# Kept for old links and commands. Practice files are now created by the learner.
Write-Output 'Practice now uses one shared file: practics/main.go.'
Write-Output 'Create practics and main.go once in VS Code, then replace the code for each episode.'
Write-Output 'EP 1-20: run go run main.go from practics.'
Write-Output 'EP 21 onward: follow the lesson for go.mod, the sensor package, and tests.'
Write-Output 'See docs/PRACTICE.md. No files were created or changed.'
