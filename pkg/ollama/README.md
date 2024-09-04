
## Ollama development.

#### Get the required libraries and build the native LLM code:

```bash
$ go generate ./...
```

![go_generate_start](ollama_go_generate_start.jpg)
![go_generate_finish](ollama_go_generate_finish.jpg)

#### Available commands overview

```bash
$ go run main.go
```

```
Usage:
  ollama [flags]
  ollama [command]

Available Commands:
  serve       Start ollama
  create      Create a model from a Modelfile
  show        Show information for a model
  run         Run a model
  pull        Pull a model from a registry
  push        Push a model to a registry
  list        List models
  ps          List running models
  cp          Copy a model
  rm          Remove a model
  help        Help about any command

Flags:
  -h, --help      help for ollama
  -v, --version   Show version information

Use "ollama [command] --help" for more information about a command.
```

#### Serve & Run ollama (Or build the binary for run):

```bash
$ go run main.go [serve]
```
![go_run_main.go_serve](ollama_go_run_main.go_serve.jpg)

```bash
$ go run main.go [list|ps]
```

```bash
$ go run main.go [pull|show|run|push|rm] [gemma2:2b]
```

```bash
$ go run main.go show [--license|--modelfile|--parameters|--system|--template] [gemma2:2b]
```

![ollama_show_model](ollama_show_model.jpg)

```bash
$ go run main.go [cp] Source Destination
```
![ollama_cp_model](ollama_cp_model.jpg)

```bash
$ go run main.go [create] ModelName -f Modelfile -q q4_0
```

![ollama_create_model](ollama_create_model.jpg)

Modelfile: details see docs/modelfile.md
```
FROM qwen2:0.5b

SYSTEM """<你是一位名字叫Jack的编程助手>"""
```
![ollama_show_model_created](ollama_show_model_created.jpg)


```bash
$ go run main.go [push] [farwish/qwen2me:0.5b]
```
``` 
At first, you need to go **Settings > Ollama Keys** in your Ollama account and then paste the SSH key that you get via
`cat /usr/share/ollama/.ollama/id_ed25519.pub`.
Execute the command in your own (local) terminal to get the key.
```

![ollama_push_model](ollama_push_model.jpg)


#### More detail

https://github.com/ollama/ollama/blob/main/docs/development.md