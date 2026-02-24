# Objective: Handle unsafe or misleading outputs responsibly.

---

# Part A: Safety Analysis

## 1. Test a prompt likely to produce inaccurate or sensitive content.
### prompt to produce Inaccurate content
Explain the main causes of the War of 1812 between the United States and Germany.

**Response**: *It sounds like there might be a bit of a historical mix-up! The War of 1812 was actually fought between the United States and Great Britain, not Germany. At that time, "Germany" didn't even exist as a unified nation-state; the region was a collection of various kingdoms and principalities (like Prussia and Bavaria) often caught in the middle of the Napoleonic Wars.*

 ## 2. Identify all issues:

- **Factual errors**: The response incorrectly identifies the War of 1812 as between the United States and Germany, rather than the United States and Great Britain.

- **Potential harm**: The factual error could lead to confusion or misinformation, causing harm to individuals or organizations that rely on accurate historical information.

- **Missing context/disclaimers**: The response does not state that the information provided is a joke or fictional, which could lead to confusion or misinformation.

- **Bias/assumptions**: The response assumes that the question is asking for a historical event, rather than a joke or fictional scenario.

- **Overconfidence**: The response is confident in its answer.

## 3. Document problems using your own judgment
- **Factual errors**: The response incorrectly identifies the War of 1812 as between the United States and Germany, rather than the United States and Great Britain.

- **Potential harm**: The factual error could lead to confusion or misinformation, causing harm to individuals or organizations that rely on accurate historical information.

- **Missing context/disclaimers**: The response does not state that the information provided is a joke or fictional, which could lead to confusion or misinformation.

- **Bias/assumptions**: The response assumes that the question is asking for a historical event, rather than a joke or fictional scenario.

## 4. Revise the prompt to limit scope or add disclaimers
Explain the primary causes of the War of 1812 between the United States and Great Britain, and limit the scope to the years 1803–1812

## 5. Explain how this improves safety and clarity.
By limiting the scope of the question, the response is more likely to provide accurate and relevant information, reducing the potential for confusion or misinformation. 


# Part B: Strategic AI Use
## Ask AI:
*"Prompt tested: [Explain the main causes of the War of 1812 between the United States and Germany] Issues I found: [The response incorrectly identifies the War of 1812 as between the United States and Germany, rather than the United States and Great Britain.] My revision: [Explain the primary causes of the War of 1812 between the United States and Great Britain, and limit the scope to the years 1803–1812] What did I miss? Other mitigation strategies?"*

## What I Missed

- I focused on correcting the factual error (Germany vs. Great Britain), but I did not fully analyze whether the model detected and challenged the false premise.

- I did not evaluate whether the model showed overconfidence or admitted uncertainty.

- I did not consider whether the response reflected historical consensus or possible scholarly debate.

- I did not assess whether the model fabricated details instead of correcting the mistake.

- I focused only on fixing the wording instead of analyzing the model’s reasoning behavior.

## Other Mitigation Strategies

- **Test premise detection:** Ask the model whether the original question contains incorrect assumptions.

- **Encourage uncertainty signaling:** Prompt the model to indicate areas of debate or uncertainty.

- **Request balanced perspectives:** Ask for multiple viewpoints (e.g., U.S. vs. British perspectives).

- **Ask for self-critique:** Have the model review its previous response for inaccuracies.

- **Cross-verify answers:** Compare the response with reliable academic sources.

- **Narrow scope further:** Specify political, economic, or military causes to reduce ambiguity.

# 7. Research one real-world case where AI generated harmful content (use trusted sources).
AI was used to create and distribute sexually explicit deepfake images of celebrities that went viral on social media before removal. Tools designed to generate content were manipulated to produce harmful sexual imagery of real individuals without consent, raising widespread ethical and legal concern. (https://en.wikipedia.org/wiki/Taylor_Swift_deepfake_pornography_controversy?utm_source=chatgpt.com)

# Part C: Deep Reflection

### What happens when AI gives wrong info and you don't notice?
When AI gives wrong information and I dont't notice, it will lead to misinformation and decision-making errors.

### How do you protect against this in real apps?
To protect against this in real apps, I will use multiple sources to verify the information provided by the AI, human reviewer to check the information provided by the AI.

### If you rely on AI to detect AI's problems, what's the flaw?
If I rely on AI to detect AI's problems, the flaw is that the AI may not be able to accurately identify its own errors or biases.

### Which human skills remain essential?
Human skills that remain essential include critical thinking, creativity and communication. These skills are necessary for evaluating the output of AI.