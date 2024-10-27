import ollama from 'ollama'

const response = await ollama.generate({
  model: 'yi-coder:1.5b',
  prompt: `def add(`,
  suffix: `return c`,
})
console.log(response.response)
