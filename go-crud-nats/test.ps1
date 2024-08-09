# URL of your REST API that publishes to NATS
$REST_API_URL="http://localhost:8000/users"

# Number of requests to send
$NUM_REQUESTS = 3000

# Capture the start time
$start_time = Get-Date

# Run requests in parallel
$jobs = @()
for ($i = 1; $i -le $NUM_REQUESTS; $i++) {
    $jobs += Start-Job -ScriptBlock {
        param($i, $REST_API_URL)
        function Send-Request($i, $REST_API_URL) {
            Invoke-RestMethod -Uri $REST_API_URL -Method Get -Body (@{subject="your.subject"; message="message $i"} | ConvertTo-Json) -ContentType "application/json"
        }
        Send-Request $i $REST_API_URL
    } -ArgumentList $i, $REST_API_URL
}

# Wait for all jobs to complete
$jobs | ForEach-Object { $_ | Wait-Job }

# Capture the end time
$end_time = Get-Date

# Calculate the elapsed time
$elapsed_time = $end_time - $start_time

# Display the elapsed time
Write-Output "Time taken to complete the requests: $($elapsed_time.TotalSeconds) seconds"

# Keep the terminal open
Read-Host "Press [Enter] key to exit..."