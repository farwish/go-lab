const axios = require('axios');
const { Readable } = require('stream');

async function fetchStreamData() {
  try {
    const response = await axios({
        method: 'POST',
        url: 'http://127.0.0.1:11434/api/generate',
        data: {
            model: 'qwen2.5:0.5b',
            prompt: 'Who is the president of the United States?',
            stream: true,
        },
        responseType:'stream', // 设置响应类型为流
        headers: {
            'Content-Type': 'application/json',
        },
        timeout: 60000,
    });

    // 将 Axios 响应转换为 Node.js 可读流
    const stream = Readable.from(response.data);

    // 读取流数据
    const readStream = async () => {
        try {
          for await (const chunk of stream) {
            processChunk(chunk);
          }
        } catch (error) {
          console.error('读取流时发生错误:', error);
        }
    };

    // 开始读取流
    readStream();
  } catch (error) {
    console.error('请求失败:', error);
  }
}

// 处理每一块数据的函数
function processChunk(chunk) {
    const textDecoder = new TextDecoder('utf-8');
    const text = textDecoder.decode(chunk);
    console.log(text); // 输出每一块数据
}

// 调用函数
fetchStreamData();