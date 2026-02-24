# Objective: Identify and explain bias in AI systems.

---


# Part A: Build Your Foundation (no AI first)

## Find or simulate a short example of model bias (e.g., gender, nationality, profession).
### 1. Gender bias
#### Example: Gendered language in job descriptions
* Description: A job description for a software engineer may use gendered language such as "female software engineer" or "male software engineer" instead of "software engineer."
* Bias: This type of bias can lead to underrepresentation of women in the field of software engineering and can perpetuate gender stereotypes.

## Explain in your own words why that bias may exist.
* The use of gendered language in job descriptions can be seen as a form of bias because it can limit the pool of potential candidates who are qualified for the position. For example, a job description that specifically mentions "female software engineer" may discourage men from applying for the position, even if they are qualified.

## Propose one mitigation (data balancing, prompt phrasing, model choice).
* One way to mitigate this type of bias is to use a more inclusive language in job descriptions. For example, instead of using gendered language, the job description could simply state "software engineer" and provide a brief description of the required skills and qualifications.

---

# Part B: Strategic AI Use
## Ask AI to deepen your understanding:
"I found this bias example: [Gender bias]. I think it exists because: [it can limit the pool of potential candidates who are qualified for the position. For example, a job description that specifically mentions "female software engineer" may discourage men from applying for the position, even if they are qualified.]. My proposed mitigation: [One way to mitigate this type of bias is to use a more inclusive language in job descriptions. For example, instead of using gendered language, the job description could simply state "software engineer" and provide a brief description of the required skills and qualifications.]. What am I missing? What underlying causes haven't I considered?"

## AI response:
## What am I missing? What underlying causes haven't I considered?

Your example and mitigation are good, but they mainly focus on surface-level wording bias. There are deeper structural causes that should also be considered.

---

### 1. Data Bias (Training Data Problem)

AI models learn from large datasets collected from the internet, books, forums, and job postings.

If historical data reflects:
- More men in tech roles
- More women in caregiving roles
- Gendered language in job ads

The model absorbs these patterns and may reproduce them.

The root issue is not just wording — it is historical inequality embedded in training data.

---

### 2. Representation Imbalance

If the dataset contains more examples of male software engineers than female software engineers, the model may statistically associate the profession with men more often.

This is not intentional bias. It is pattern learning from imbalanced representation.

---

### 3. Societal and Cultural Norms

Bias does not start with AI. It originates from:
- Media portrayals
- Workplace demographics
- Educational access
- Cultural stereotypes

AI systems reflect society. If society contains bias, models may mirror it unless corrected.

---

### 4. Algorithmic Reinforcement (Feedback Loops)

Even small biases can grow stronger over time.

For example:
- Biased job recommendations → fewer women apply  
- Fewer applicants → fewer women hired  
- Hiring data feeds back into future systems  

This creates a feedback loop that reinforces inequality.

---

### 5. Implicit vs Explicit Bias

The example focuses on explicit wording like:
 "female software engineer"

However, bias can also be implicit, such as:
- Using words like “aggressive,” “dominant,” or “rockstar” in job ads  
- Associating leadership traits more with men  
- Recommending higher salaries to one gender  

Implicit bias is harder to detect but equally important.

---

### 6. Structural Hiring Bias (Beyond Language)

Even with neutral language:
- Resume screening algorithms may favor certain demographics  
- Historical hiring data may skew predictions  
- Promotion data may reflect past discrimination  

Mitigation may require:
- Dataset auditing  
- Bias testing across demographic groups  
- Fairness-aware training techniques  
- Human oversight  

---

## Stronger Framing

Instead of saying:

Bias exists because language can discourage candidates.

You could say:

Bias exists because AI systems learn from historically imbalanced data and societal stereotypes, which can unintentionally reinforce unequal representation unless actively mitigated.

---
# Then explore edge cases:

What happens when this bias appears in healthcare, hiring, or criminal justice? Here are my predictions: [Healthcare: AI systems may misdiagnose or overdiagnose certain conditions based on historical data. 
Hiring: AI systems may recommend higher salaries to one gender based on historical hiring data. 
Criminal Justice: AI systems may predict reoffending rates based on historical data leading to unfair sentencing]. What are the real-world implications?"

## Real-World Implications of Bias in Critical Systems

Bias in AI systems is not just a technical issue — it can directly affect people’s health, income, freedom, and life outcomes.

---

### 🏥 Healthcare

**Your prediction:**  
AI systems may misdiagnose or overdiagnose certain conditions based on historical data.

**Real-world implications:**

- Delayed or incorrect treatment for certain demographic groups  
- Higher mortality rates due to underdiagnosis  
- Over-treatment of some groups and under-treatment of others  
- Reduced trust in medical systems  
- Widening existing health disparities  

If historical data underrepresents certain populations, the AI may perform worse for them. For example, diagnostic systems trained mostly on one demographic group may fail to detect diseases accurately in others.

---

### 💼 Hiring

**Your prediction:**  
AI systems may recommend higher salaries to one gender based on historical hiring data.

**Real-world implications:**

- Reinforcement of wage gaps  
- Fewer job opportunities for underrepresented groups  
- Biased resume filtering that excludes qualified candidates  
- Reduced workplace diversity  
- Long-term economic inequality  

If past hiring favored certain groups, AI systems trained on that data may continue to prefer similar profiles, creating a cycle of exclusion.

---

### ⚖️ Criminal Justice

**Your prediction:**  
AI systems may predict reoffending rates based on historical data leading to unfair sentencing.

**Real-world implications:**

- Harsher sentencing for certain communities  
- Denial of parole based on biased risk assessments  
- Increased incarceration rates for specific demographic groups  
- Loss of public trust in the justice system  
- Institutional reinforcement of systemic inequality  

If historical arrest or policing data reflects biased enforcement patterns, predictive systems may incorrectly label some groups as “higher risk,” even if the underlying data is skewed.

---

## Broader Consequences

Across all three domains, biased AI systems can:

- Scale discrimination faster than human decision-making  
- Make bias appear “objective” because it is algorithmic  
- Reduce accountability if decisions are blamed on the model  
- Amplify historical inequalities instead of correcting them  

Because these systems operate at scale, even small biases can affect thousands or millions of people.

---

## Key Insight

When bias appears in high-stakes systems like healthcare, hiring, or criminal justice, it does not just create unfair outputs — it shapes life opportunities, access to resources, and long-term societal equality.

That is why auditing, transparency, and human oversight are essential.

---

# Part C: Critical Reflection

## What % did you complete before using AI?

I completed 40% of the course before using AI.

## Did AI replace your thinking or amplify it?

AI amplified my thinking. It helped me to understand the concepts and ideas.

## Could you explain this to someone else without AI?
Yes, I can explain the concepts and ideas without AI.

## What did you contribute that AI couldn't?
I contributed my understanding which is based on my knowledge and experience that is different from AI.

---


