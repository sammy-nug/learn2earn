# **Objective:** Understand how parameters affect output diversity.

## Use the same prompt with temperature = 0.2 and then with temperature = 0.9.
The difference in outputs between temperature = 0.2 and temperature = 0.9 is significant. When the temperature is set to 0.2, the output is more concise and focused on the main points, while when the temperature is set to 0.9, the output is more detailed and includes additional information. This difference in outputs is due to the way the temperature parameter affects the randomness of the model's predictions. A lower temperature makes the model more confident in its predictions and produces more deterministic outputs, while a higher temperature makes the model more random and produces more diverse outputs.

### Repeat with different top-p values (e.g., 0.5 vs 1).
I repeated the prompt with different top-p values (0.5 and 1) and observed the differences in the outputs.

With top-p = 0.5, the output was more concise and focused on the main points, while with top-p = 1, the output was more detailed and included additional information. This difference in outputs is due to the way the top-p parameter affects the diversity of the model's predictions. A lower top-p value makes the model more confident in its predictions while a higher top-p value makes the model more random and produces more diverse outputs.

### Record how the style, randomness, and focus of responses change.
I recorded how the style, randomness, and focus of responses changed as I adjusted the temperature and top-p values.

As I increased the temperature, the randomness of the responses increased, and the focus of the responses became less clear. This is because a higher temperature makes the model more random and produces more diverse outputs.

As I decreased the temperature, the randomness of the responses decreased, and the focus of the responses became more clear. This is because a lower temperature makes the model more confident in its predictions and produces more deterministic outputs.