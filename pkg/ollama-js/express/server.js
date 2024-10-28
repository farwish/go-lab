const express = require('express');
const axios = require('axios');
const app = express();
const port = 3000;

app.use(express.json()); // 解析 JSON 请求体

// 定义 API 路由
app.post('/api/gen', async (req, res) => {
  try {
    const { model, prompt, stream } = req.body;

    // 向 ollama generate 接口发送请求
    const response = await axios.post('http://127.0.0.1:11434/api/generate', {
      model,
      prompt,
      stream,
    }, {
      responseType: 'json', // 设置响应类型为 JSON
      headers: {
        'Content-Type': 'application/json',
      },
      timeout: 60000, // 设置请求超时时间（60秒）
    });

    // 将响应数据返回给客户端
    res.json(response.data);
  } catch (error) {
    console.error('请求失败:', error);
    res.status(500).json({ error: '请求失败' });
  }
});

// 启动服务器
app.listen(port, () => {
  console.log(`服务器运行在 http://localhost:${port}`);
});