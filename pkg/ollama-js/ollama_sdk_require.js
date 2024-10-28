const ollama = require('ollama');

// 动态导入 ES Module
import('ollama').then((module) => {
  const ollama = module.default;

  // 设置一个 1 秒后的超时回调，用于中止请求
  setTimeout(() => {
    console.log('\nAborting request...\n');
    ollama.abort();
  }, 1000); // 1000 milliseconds = 1 second

  ollama.generate({
    model: 'qwen2.5:0.5b',
    prompt: 'why the sky is blue?',
    stream: true,
  }).then(
    async (stream) => {
      for await (const chunk of stream) {
        process.stdout.write(chunk.response);
      }
    }
  ).catch(
    (error) => {
      if (error.name === 'AbortError') {
        console.log('The request has been aborted');
      } else {
        console.error('An error occurred:', error);
      }
    }
  );
}).catch((error) => {
  console.error('Failed to import ollama:', error);
});