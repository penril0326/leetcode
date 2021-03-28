# Solution

### Description
Add two numbers without using '+' or any arithmetic operator.


### Level
Hard

### Explain
Because computer is alway use binary so we do this same.

Considerate two numbers 85 and 67.  

1. Add two numbers and ignore carry, you will get 42.
1. Add then only care carry, you will get 110.
1. Add two results you will get final answer 152.

Now we considerate with binary:
1. If we add two binary numbers and ignore carry, it would be **XOR**.
2. If we add two binary numbers and only care about carry, then you will find it would be **AND**.
3. Finally, recurse util no carry.