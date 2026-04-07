# register_schemas.ps1
# Registers Protobuf schemas for OTIE with the Local Schema Registry

$REGISTRY_URL = "http://localhost:8081"

function Register-ProtobufSchema {
    param (
        [string]$Subject,
        [string]$FilePath
    )

    Write-Host "📡 Registering schema for subject: $Subject from $FilePath..." -ForegroundColor Cyan

    if (-not (Test-Path $FilePath)) {
        Write-Error "❌ File not found: $FilePath"
        return
    }

    # Read the proto file content
    $schemaContent = Get-Content -Raw $FilePath

    # Properly construct JSON payload using PowerShell's ConvertTo-Json
    $payloadObj = @{
        schema = $schemaContent
        schemaType = "PROTOBUF"
    }
    $payload = $payloadObj | ConvertTo-Json -Compress

    try {
        $response = Invoke-RestMethod -Uri "$REGISTRY_URL/subjects/$Subject/versions" `
                                     -Method Post `
                                     -Body $payload `
                                     -ContentType "application/vnd.schemaregistry.v1+json"
        
        Write-Host "✅ Successfully registered! Version: $($response.version)" -ForegroundColor Green
    }
    catch {
        Write-Host "❌ Registration failed for $Subject" -ForegroundColor Red
        if ($_.Exception.Response) {
            $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
            $errBody = $reader.ReadToEnd()
            Write-Host "🔍 Error details: $errBody" -ForegroundColor Yellow
        } else {
            Write-Host "🔍 Error: $($_.Exception.Message)" -ForegroundColor Yellow
        }
    }
}

# --- Main ---

# 1. Register Collector (Raw Events)
Register-ProtobufSchema -Subject "raw-events-value" -FilePath "api/v1/collector.proto"

# 2. Register Processor (Processed Features)
Register-ProtobufSchema -Subject "processed-events-value" -FilePath "api/v1/processor/feature.proto"

Write-Host "🎉 All schemas registered!" -ForegroundColor Green
