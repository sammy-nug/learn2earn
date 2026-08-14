# 🎯 Objective: Create Prompts That Produce Structured Machine-Readable Output

---

## 1️⃣ Write a prompt that asks for unstructured information (e.g., “List three cities with their populations.”).

**Prompt**
> List three cities with their populations.
  

###  Example Output (Unstructured)

- Tokyo has a population of about 37 million people.  
- New York City has around 8.5 million residents.  
- Paris has approximately 2.1 million inhabitants.  

---

## 2️⃣ Rewrite the prompt to enforce structured JSON output:

> List three real cities with their populations.  
> Return the result in JSON format using this structure:  
> `[{"city": "string", "population": number}]`  
>  
 

---


# Test with multiple inputs and verify the JSON stays valid.

## Test Case 1

### Prompt
> List three real cities with their populations.  
> Return the result in JSON using this format:  
> `[{"city": "string", "population": number}]`

### Output

```json
[
  {"city": "Tokyo", "population": 37000000},
  {"city": "New York City", "population": 8500000},
  {"city": "Paris", "population": 2100000}
]
```

## Test Case 2
### Prompt

> List three different real cities with their populations.
> Return strictly valid JSON using this format:
> [{"city": "string", "population": number}]
```json
[
  {"city": "Lagos", "population": 15000000},
  {"city": "Berlin", "population": 3800000},
  {"city": "Sydney", "population": 5300000}
]
```
