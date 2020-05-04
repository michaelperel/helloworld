Build:

```
docker build -t mperel/helloworld .
```

Run:

```
docker run -it --rm -e NAME=YOURNAME -p 8080:8080 mperel/helloworld
```

Navigate to localhost:8080 in your browser to see the hello world page.
