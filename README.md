# Kubernetes Audit Webhook

## How to Use
1. Build an Docker image
```bash
docker build -t k8s-audit-webhook:1.0 .
```

2. Launch a Docker container
```bash
docker run -d -p 8080:8080  --name webhook k8s-audit-webhook:1.0
```

## Validation
* Send a sample events
```bash
curl -vvv -sX POST http://localhost:8080/audits -d@events.json -H 'Content-Type:application/json'
```
* Check logs
```bash
docker logs webhook
```