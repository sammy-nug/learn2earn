# 🎯 Objective: Practice Designing Prompts That Mimic Function Calls

---

## 1️⃣ Create a prompt: “Simulate a weather API response for Paris today.”

> Simulate a weather API response for Paris today. 

---

## 2️⃣ Define an expected JSON schema, 
**Improved Prompt with Defined JSON Schema**

> Simulate a weather API response for Paris today.  
> Return JSON using the following schema:  
>  
> ```json
> {
>   "location": "string",
>   "temperature_celsius": number,
>   "condition": "string"
> }
> ```

---

# Test the model with different cities and compare outputs

## Test Case 1: Paris

### Prompt
> Simulate a weather API response for Paris today.  
> Return valid JSON using the defined schema.

### Output

```json
{
  "location": "Paris",
  "temperature_celsius": 18,
  "condition": "Cloudy"
}
```
## Test Case 2: Tokyo

### Prompt

> Simulate a weather API response for Tokyo today.
> Return valid JSON using the defined schema.

### Output
```json
[
  {"city": "Lagos", "population": 15000000},
  {"city": "Berlin", "population": 3800000},
  {"city": "Sydney", "population": 5300000}
]
```
## Compare outputs.

### Across multiple inputs:

- JSON syntax remains valid  
- Keys remain consistent (`location`, `temperature_celsius`, `condition`)  
- Data types remain correct  
- No additional commentary appears  
- Structure does not drift  