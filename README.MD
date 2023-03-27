## :wrench: Run Container
Build Container Image:
```
docker build -t shell_go:latest .
```

Run Container:
```
docker run --name shell -v /github.com/ThreeDP/shell/ -d shell_go:latest
```

Run Bash Env:
```
docker exec -it shell bash
```

