
# Ollama development.

https://github.com/ollama/ollama/blob/main/docs/development.md

## Commands overview

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

## Serve ollama in docker (Without llama.cpp local generate；to make LlamaServer available)

```
# Serve, than other command will not show "Error: could not connect to ollama app, is it running?"
$ docker run -d -v /home/ubuntu/.ollama:/root/.ollama -p 11434:11434 --name ollama_deploy ollama/ollama

# Serve with OLLAMA_ORIGINS=* (to fix CORS error from remote)
$ docker run -e OLLAMA_ORIGINS=* -d -v /home/ubuntu/.ollama:/root/.ollama -p 11434:11434 --name ollama_deploy ollama/ollama

# Serve with --privileged (to fix permission denied error on Windows), here `-v ollama:` is volume.
$ docker run --privileged -d -v ollama:/root/.ollama -p 11434:11434 --name ollama_deploy ollama/ollama

# Test running
$ curl http://localhost:11434  # Ollama is running

# List models (2 ways)
$ go run main.go run qwen2:0.5b                         # With codebase
$ docker exec -it ollama_deploy ollama run qwen2:0.5b   # Without codebase
```

### docker basic knowledge

```
1.
Install see Docs: https://docs.docker.com/engine/install/ubuntu/
Download: https://developer.aliyun.com/mirror/docker-ce/

2.
$ [sudo useradd -d /home/ubuntu -s /bin/bash ubuntu]
$ sudo usermod -aG docker ubuntu

3.
$ sudo vi /etc/docker/daemon.json
{
  "registry-mirrors": [
    "https://docker.1panel.live/",
    "https://hub.rat.dev"
  ],
}

4.
$ sudo systemctl restart docker
```

## Use ollama cmd (When LlamaServer available):

```bash
$ go run main.go [list|ps]
```

```bash
$ go run main.go [stop] [gemma2:2b]
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
PARAMETER temperature 1
SYSTEM """
你是一位名字叫Jack的编程助手
"""
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

## REST API of ollama

Code is *server/routes.go*
Docs is *docs/api.md* contains advanced detail.

Overview:
```
1. `model` param is required

  like `{ "model": "llama3.1" }` 

2. `prompt` param for /api/generate，/api/embeddings 

  like `{ "prompt": "Hello world" }`

2. `steam` param

  like `{ "stream": true }` or `{ "stream": false }`

3. `options` param

  like `{ "options": {"seed":123, "temperature": 0.2} }`, same of `parameter` in modelfile. 

  for /api/generate，/api/chat，/api/embed，/api/embeddings

  default options see src code \ollama\api\types.go:DefaultOptions()

4. `image` param

  like `{ "image": ["this_is_image_base64_string"] }`

  for multimodal /api/chat

5. `message` param 

  like `{ "message": {"role": "user", "content": "hello"} }`

   `messages` param, for /api/chat with history, 
   
  like `{ "messages": [{"role": "user", "content": "hello"},{"role": "assistant", "content": "hi"},{"role": "user", "content": "how are you?"}] }`

  for /api/chat

6. `tools` param, Requires `stream` to be set to `false`
  （see docs/api.md）

    js usage example: [ollama/ollama-js](https://github.com/ollama/ollama-js)/examples/tools/tools.ts
```

options explain:

```
【Options 结构体】
// Predict options used at runtime

NumKeep (int):4
  保留的候选序列数量。在生成文本时，保留一定数量的最佳候选序列。
  // set a minimal num_keep to avoid issues on context shifts

Seed (int):-1
  随机数生成器的种子。用于初始化随机数生成器，确保结果可复现。

NumPredict (int):-1
  预测时生成的候选序列数量。每次预测时生成的候选序列数量。

TopK (int):40
  选择概率最高的前 K 个候选词。用于限制候选词汇的数量，提高生成质量。

TopP (float32):0.9
  核采样（Nucleus Sampling）阈值。选择累积概率最高的词汇，直到达到 TopP 阈值。

MinP (float32):
  最小概率阈值。低于此阈值的词汇将被忽略。

TFSZ (float32):1.0
  温度强制采样（Temperature Forced Sampling）的 Z 参数。用于调整采样过程中的温度。

TypicalP (float32):1.0
  典型采样（Typical Sampling）的概率阈值。选择典型分布中的词汇。

RepeatLastN (int):64
  计算重复惩罚时考虑的最后 N 个词汇。用于避免生成重复的词汇。

Temperature (float32):0.8
  采样温度，控制随机性。温度越高，生成的结果越随机；温度越低，生成的结果越确定。

RepeatPenalty (float32):1.1
  重复词汇的惩罚系数。对于重复出现的词汇增加惩罚，降低其概率。

PresencePenalty (float32):0.0
  存在惩罚系数。对于已经出现过的词汇增加惩罚，降低其再次出现的概率。

FrequencyPenalty (float32):0.0
  频率惩罚系数。对于频繁出现的词汇增加惩罚，降低其再次出现的概率。

Mirostat (int):0
  使用 Mirostat 算法的版本。Mirostat 是一种动态调整温度的方法，用于改进生成质量。

MirostatTau (float32):5.0
  Mirostat 算法中的 τ 参数。用于控制 Mirostat 算法的灵敏度。

MirostatEta (float32):0.1
  Mirostat 算法中的 η 参数。用于控制 Mirostat 算法的学习速率。

PenalizeNewline (bool):true
  是否惩罚换行符。如果为 true，则减少换行符出现的概率。

Stop ([]string):
  停止生成的字符串列表。当生成的文本包含这些字符串之一时，停止生成。


【Runner 结构体】
// Runner options which must be set when the model is loaded into memory

NumCtx (int):2048
  上下文窗口大小。模型处理文本时考虑的最大上下文长度。

NumBatch (int):512
  批处理大小。每次处理的样本数量。

NumGPU (int):-1
  使用的 GPU 数量。指定使用的 GPU 设备数量。
  // -1 here indicates that NumGPU should be set dynamically

MainGPU (int):
  主 GPU 的 ID。指定主 GPU 的设备 ID。

LowVRAM (bool):false
  是否启用低显存模式。如果为 true，则优化显存使用。

F16KV (bool):true
  是否使用半精度浮点数存储键值对。如果为 true，则使用 FP16 存储键值对，节省显存。

LogitsAll (bool):
  是否计算所有位置的 logits。如果为 true，则计算整个序列的 logits。

VocabOnly (bool):
  是否仅加载词汇表。如果为 true，则只加载词汇表而不加载其他数据。

UseMMap (*bool):false
  是否使用内存映射文件。如果为 true，则使用内存映射文件加载数据。

UseMLock (bool):false
  是否锁定内存以防止交换到磁盘。如果为 true，则锁定内存，提高性能。

NumThread (int):0
  使用的线程数。指定并行处理时使用的线程数量。
  // 0 let the runtime decide
```

## Serve ollama without docker (After llama.cpp local generate):

```bash
$ go run main.go [serve]
```
![go_run_main.go_serve](ollama_go_run_main.go_serve.jpg)

## Build Native LLM code with required libraries (llama.cpp local generate):

```bash
$ go generate ./...
```

![go_generate_start](ollama_go_generate_start.jpg)
![go_generate_finish](ollama_go_generate_finish.jpg)
