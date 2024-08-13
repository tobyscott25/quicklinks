# Quick Links

## Development

```bash
docker compose up
```

## Production

```bash
docker build -t quick-links .
docker run -p 8080:8080 -v /path/to/redirects.yaml:/redirects.yaml quick-links
```
