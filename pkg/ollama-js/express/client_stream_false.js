const axios = require('axios');

async function fetchGeneratedData() {
  try {
    const response = await axios.post('http://localhost:3000/api/gen', {
      model: 'qwen2.5:0.5b',
      prompt: 'why the sky blue?',
      stream: false, // 根据需要设置是否为流式响应
    });

    // 处理响应数据
    console.log('生成的数据:', response.data.response);
  } catch (error) {
    console.error('请求失败:', error);
  }
}

// 调用函数
fetchGeneratedData();