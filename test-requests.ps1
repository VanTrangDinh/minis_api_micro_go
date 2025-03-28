# Cấu hình headers
$headers = @{
    "Content-Type" = "application/json"
}

# Hàm gửi request và xử lý lỗi
function Send-Request {
    param (
        [string]$Uri,
        [string]$Method,
        $Body,
        $AuthToken
    )
    
    $requestHeaders = $headers.Clone()
    if ($AuthToken) {
        $requestHeaders["Authorization"] = "Bearer $AuthToken"
    }
    
    try {
        if ($Body) {
            $response = Invoke-WebRequest -Uri $Uri -Method $Method -Headers $requestHeaders -Body $Body -ErrorAction Stop
        } else {
            $response = Invoke-WebRequest -Uri $Uri -Method $Method -Headers $requestHeaders -ErrorAction Stop
        }
        Write-Host "Success: $Method $Uri - Status: $($response.StatusCode)" -ForegroundColor Green
        return $response
    }
    catch {
        Write-Host "Error: $Method $Uri - Status: $($_.Exception.Response.StatusCode)" -ForegroundColor Red
        return $null
    }
}

# Hàm trích xuất token từ response
function Get-Token {
    param (
        $Response
    )
    if ($Response) {
        $content = $Response.Content | ConvertFrom-Json
        return $content.token
    }
    return $null
}

# Vòng lặp gửi request
for ($i = 1; $i -le 20; $i++) {
    Write-Host "`nTesting iteration $i" -ForegroundColor Cyan
    
    # Register request
    $registerBody = @{
        username = "testuser$i"
        password = "password123"
        email = "test$i@example.com"
    } | ConvertTo-Json
    $registerResponse = Send-Request -Uri "http://localhost:8081/register" -Method "Post" -Body $registerBody
    Start-Sleep -Milliseconds 200

    # Login request
    $loginBody = @{
        username = "testuser$i"
        password = "password123"
    } | ConvertTo-Json
    $loginResponse = Send-Request -Uri "http://localhost:8081/login" -Method "Post" -Body $loginBody
    $token = Get-Token -Response $loginResponse
    Start-Sleep -Milliseconds 200

    if ($token) {
        # Test protected endpoints
        Send-Request -Uri "http://localhost:8081/api/users/me" -Method "Get" -AuthToken $token
        Start-Sleep -Milliseconds 200

        # Test some invalid requests to generate 4xx errors
        Send-Request -Uri "http://localhost:8081/api/invalid" -Method "Get" -AuthToken $token
        Start-Sleep -Milliseconds 200

        # Test with invalid token to generate 401 errors
        Send-Request -Uri "http://localhost:8081/api/users/me" -Method "Get" -AuthToken "invalid_token"
        Start-Sleep -Milliseconds 200
    }

    # Random delay between iterations
    $delay = Get-Random -Minimum 500 -Maximum 2000
    Start-Sleep -Milliseconds $delay
}

Write-Host "`nTest completed!" -ForegroundColor Green 