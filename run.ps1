
param(
    [Parameter(Mandatory=$false)]
    [ValidateSet("dev", "stage", "prod", "build", "clean", "migrate", "reset")]
    [string]$Command = "help"
)

function Show-Help {
    Write-Host "Available commands:"
    Write-Host "  .\run.ps1 dev      - Run in development mode"
    Write-Host "  .\run.ps1 stage    - Run in staging mode"
    Write-Host "  .\run.ps1 prod     - Run in production mode"
    Write-Host "  .\run.ps1 migrate  - Run database migrations"
    Write-Host "  .\run.ps1 reset    - Reset database (drop all tables and migrate)"
    Write-Host "  .\run.ps1 build    - Build the application"
    Write-Host "  .\run.ps1 clean    - Remove binary files"
    Write-Host "  .\run.ps1 help     - Show this help message"
}

function Run-Development {
    Write-Host "Running in development mode..." -ForegroundColor Green
    go run main.go --env=development
}

function Run-Staging {
    Write-Host "Running in staging mode..." -ForegroundColor Yellow
    go run main.go --env=staging
}

function Run-Production {
    Write-Host "Running in production mode..." -ForegroundColor Red
    go run main.go --env=production
}

function Run-Migrations {
    Write-Host "Running database migrations..." -ForegroundColor Cyan
    go run main.go --env=development --migrate
}

function Reset-Database {
    Write-Host "Resetting database..." -ForegroundColor Magenta
    $confirmation = Read-Host "Are you sure you want to reset the database? (yes/no)"
    if ($confirmation -eq "yes" -or $confirmation -eq "y") {
        go run main.go --env=development --reset
    } else {
        Write-Host "Database reset cancelled." -ForegroundColor Gray
    }
}

function Build-Application {
    Write-Host "Building application..." -ForegroundColor Cyan
    go build -o main.exe .
}

function Clean-Up {
    Write-Host "Cleaning up..." -ForegroundColor Gray
    if (Test-Path "main.exe") { Remove-Item "main.exe" -Force }
    Write-Host "Cleanup completed!" -ForegroundColor Green
}

switch ($Command) {
    "dev" { Run-Development }
    "stage" { Run-Staging }
    "prod" { Run-Production }
    "migrate" { Run-Migrations }
    "reset" { Reset-Database }
    "build" { Build-Application }
    "clean" { Clean-Up }
    "help" { Show-Help }
    default { Show-Help }
}
