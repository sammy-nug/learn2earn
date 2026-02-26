# 🎯 Objective: Combine Multiple Prompts to Simulate a Tool-Using Agent

---


#  Step 1: Ask for the User’s Query

### Prompt
> What would you like to know about the weather?

### Example User Input
> What’s the temperature in Paris today?

### Output of Step 1
```json
{
  "intent": "get_weather",
  "location": "Paris"
}
```
# Step 2: Simulate Weather API Call

### Prompt

Simulate a weather API response using this schema:
```json
{
  "location": "string",
  "temperature_celsius": number,
  "condition": "string"
}
```
**Simulated API Output**
```json
{
  "location": "Paris",
  "temperature_celsius": 18,
  "condition": "Cloudy"
}
```
# Step 3 prompt: Format the final answer back to the user.

**Prompt**

Using the API response, generate a natural language reply for the user.

# Document the chained process and outputs.

## Chained Flow Summary

| Step   | Purpose               | Output Type             |
| ------ | --------------------- | ----------------------- |
| Step 1 | Interpret user query  | Structured intent JSON  |
| Step 2 | Simulate tool call    | Weather API JSON        |
| Step 3 | Format final response | Natural language answer |
