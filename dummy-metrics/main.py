from prometheus_client import start_http_server, Gauge, Counter, Histogram, Summary
import random
import time

# Metrics
# Build Metrics
BUILD_REQUESTS = Counter('build_requests_total', 'Total number of build requests')
BUILD_DURATION = Histogram('build_request_duration_seconds', 'Duration of build requests')

# Preview Metrics
PREVIEW_REQUESTS = Counter('preview_requests_total', 'Total number of preview requests')
PREVIEW_DURATION = Histogram('preview_request_duration_seconds', 'Duration of preview requests')

# Deployment Metrics
DEPLOYMENTS = Counter('deployments_total', 'Total number of deployments', ['status']) # status="success" or "failure"
DEPLOYMENT_DURATION = Histogram('deployment_duration_seconds', 'Duration of deployments')

def update_metrics():
    # Simulate Build Requests
    if random.random() > 0.7:
        BUILD_REQUESTS.inc()
        BUILD_DURATION.observe(random.uniform(5, 120)) # Builds take 5-120s

    # Simulate Preview Requests
    if random.random() > 0.6:
        PREVIEW_REQUESTS.inc()
        PREVIEW_DURATION.observe(random.uniform(1, 10)) # Previews are fast

    # Simulate Deployments
    if random.random() > 0.8:
        status = 'success' if random.random() > 0.1 else 'failure'
        DEPLOYMENTS.labels(status=status).inc()
        if status == 'success':
            DEPLOYMENT_DURATION.observe(random.uniform(10, 60))

if __name__ == '__main__':
    # Start up the server to expose the metrics.
    start_http_server(8000)
    print("Serving metrics on port 8000")
    while True:
        update_metrics()
        time.sleep(1)
