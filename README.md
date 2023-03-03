# Layoffs

## Running
### Container
At project root, run:

```bash
docker build -f Dockerfile -t layoffs-api:latest .

docker -p 3000:3000 run layoffs-api:latest
```