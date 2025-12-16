# Check for kubectl
if (-not (Get-Command kubectl -ErrorAction SilentlyContinue)) {
    Write-Host "kubectl not found. Please ensure it is installed and in your PATH." -ForegroundColor Red
    exit 1
}

# Create Namespace
Write-Host "Creating namespace sparrow-vps..." -ForegroundColor Cyan
kubectl create namespace sparrow-vps
if ($LASTEXITCODE -ne 0 -and $LASTEXITCODE -ne 1) {
    # 1 might be "already exists" depending on kubectl version, but usually it's non-zero
    # Actually kubectl create ns returns 1 if it exists. We should check if it exists instead.
}
# Safer way:
kubectl get namespace sparrow-vps > $null 2>&1
if ($LASTEXITCODE -ne 0) {
    kubectl create namespace sparrow-vps
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Failed to create namespace." -ForegroundColor Red
        exit 1
    }
}

# Build Frontend Image
Write-Host "Building frontend Docker image..." -ForegroundColor Cyan
docker build -t sparrow-frontend ./frontend
if ($LASTEXITCODE -ne 0) {
    Write-Host "Docker build failed." -ForegroundColor Red
    exit 1
}

# Build Dummy Metrics Image
Write-Host "Building dummy metrics Docker image..." -ForegroundColor Cyan
docker build -t sparrow-dummy-metrics ./dummy-metrics
if ($LASTEXITCODE -ne 0) {
    Write-Host "Dummy metrics build failed." -ForegroundColor Red
    exit 1
}

# Function to deploy resources
function Deploy-Resource {
    param (
        [string]$ResourcePath
    )
    Write-Host "Applying $ResourcePath" -ForegroundColor Cyan
    kubectl apply -f $ResourcePath
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Failed to apply $ResourcePath" -ForegroundColor Red
        exit 1
    }
}

$BasePath = ".\kubernetes"
$RepoPath = "$BasePath\repo-service"
$ContPath = "$BasePath\container-service"
$FronPath = "$BasePath\frontend"
$DepPath = "$BasePath\deploy-service"

# Deploy Resources
Deploy-Resource "$BasePath\sparrow-pv.yml"
Deploy-Resource "$BasePath\repo-data-pvc.yml"

Deploy-Resource "$RepoPath\repo-service-deployment.yml"
Deploy-Resource "$RepoPath\repo-service-svc.yml"

Deploy-Resource "$ContPath\container-service-deployment.yml"
Deploy-Resource "$ContPath\container-service-svc.yml"

Deploy-Resource "$FronPath\frontend-deployment.yml"
Deploy-Resource "$FronPath\frontend-service.yml"
Deploy-Resource "$FronPath\frontend-ingress.yml"

Deploy-Resource "$DepPath\deploy-service-deployment.yml"
Deploy-Resource "$DepPath\deploy-service-svc.yml"

# Deploy Monitoring Stack
$MetricsPath = "$BasePath\metrics"
Write-Host "Deploying Monitoring Stack..." -ForegroundColor Cyan
Deploy-Resource "$MetricsPath\monitoring-namespace.yaml"
Deploy-Resource "$MetricsPath\prometheus-rbac.yaml"
Deploy-Resource "$MetricsPath\prometheus-config.yaml"
Deploy-Resource "$MetricsPath\prometheus-deployment.yaml"
Deploy-Resource "$MetricsPath\prometheus-ingress.yaml"
Deploy-Resource "$MetricsPath\grafana-datasources.yaml"
Deploy-Resource "$MetricsPath\grafana-dashboard-provider.yaml"
Deploy-Resource "$MetricsPath\grafana-dashboards.yaml"
Deploy-Resource "$MetricsPath\grafana-deployment.yaml"
Deploy-Resource "$MetricsPath\grafana-ingress.yaml"
Deploy-Resource "$MetricsPath\dummy-metrics.yaml"

Write-Host "All resources applied. Note: Host file (C:\Windows\System32\drivers\etc\hosts) modification is manual on Windows. Map 'sparrow-vps.local' to your cluster IP (usually 127.0.0.1 for Docker Desktop)." -ForegroundColor Yellow
