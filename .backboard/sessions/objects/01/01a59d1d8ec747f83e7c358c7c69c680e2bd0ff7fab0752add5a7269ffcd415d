# 🎯 Objective: Extract Structured Data from Unstructured Input

---

## 1️⃣ Choose an unstructured text (e.g., "John Doe, age 29, lives in Paris and works as a software engineer.").

**Input Text:**  
John Doe, age 29, lives in Paris and works as a software engineer.

---

## 2️⃣ Create a template prompt:

> Extract the following fields from the text: Name, Age, Location, Occupation.  
> Return the output in JSON format.

---

## 3️⃣ Validate the output and ensure consistency across multiple inputs.

**Expected Output (Validated)**
```json
{
  "Name": "John Doe",
  "Age": 29,
  "Location": "Paris",
  "Occupation": "Software Engineer"
}
```

# Validation Across Multiple Inputs

## Input:
Maria Lopez, 34 years old, is a data analyst based in Madrid.

## Output:
```json
{
  "Name": "Maria Lopez",
  "Age": 34,
  "Location": "Madrid",
  "Occupation": "Data Analyst"
}
```
# Test Case 2

## Input:
David Chen is a 42-year-old teacher living in Toronto.

## Output:
```json
{
  "Name": "David Chen",
  "Age": 42,
  "Location": "Toronto",
  "Occupation": "Teacher"
}
```

##  Consistency Check

Across all examples:

- The same four fields are extracted.  
- Field names remain identical: `Name`, `Age`, `Location`, `Occupation`.  
- JSON structure is uniform.  
- Age is consistently represented as a number.  
- Occupations are standardized in title case.  