
## Ollama development.

#### Get the required libraries and build the native LLM code:

```bash
$ go generate ./...
```

![go_generate_start](ollama_go_generate_start.jpg)
![go_generate_finish](ollama_go_generate_finish.jpg)

#### Run ollama (Or build at first):

```bash
$ go run main.go [serve]
```

```bash
$ go run main.go [ps|xxx]
```

#### More detail

https://github.com/ollama/ollama/blob/main/docs/development.md