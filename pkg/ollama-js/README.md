# ollama-js-demo

Ollama JavaScript library: [ollama-js](https://github.com/ollama/ollama-js)

### Tools example

```bash
$ yarn                        # Install project dependencies
$ yarn global add typescript  # Install TypeScript, Than `echo $PATH` to confirm PATH contains the `yarn global bin` path
$ tsc tools.ts && node tools.js # Build and Run to see results
```

`tools.ts` flow explain:

```js
// Initialize conversation with a user query
let messages = [{ role: 'user', content: 'What is the flight time from New York (NYC) to Los Angeles (LAX)?' }];

// First API call: Send a messages and function description to the model
const response = await ollama.chat({
    // xx
})

// Add the model's response to the conversation history
messages.push(response.message);

// Check if the model decided to use the provided function
if (!response.message.tool_calls || response.message.tool_calls.length === 0) {
    // xx
}

// Process function calls made by the model
if (response.message.tool_calls) {
    // xx
}

// Second API call: Get final response from the model
const finalResponse = await ollama.chat({
    // xx
})
console.log(finalResponse.message.content);
```