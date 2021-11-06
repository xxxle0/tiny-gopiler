# Introduction:
I want to do some exp to understand how the compiler can check the code syntax. I know it will AST (Abstract Syntax Tree) from the code then traverse the tree to check the syntax is correct or not?

My quote: ***You Understand It <-> You Build It.**
# What?
How the compiler separate the expression, statement, function?
We know that a file of code contains many statement and expression
## What is expression?
Expression is a snippet of code that return the value
For Example:
- 1 + 2 (Compute expression)
- "Hello World" (String expression)
Fun fact: IIFE (Invoke Immediately Function Expression) of Javascript is an expression
## What is statement?
Statement is a snippet of code that performs an action
For Example:
- If else statement
- Loop statement
## What is Compiler?
Compiler is a program which transform the code from A type to B type. We have many compilers for many purposes for example:
C++ compiler (transform the codebase from C++/C to Assembly then Machine Code) 

# How? 
## How the compiler works?
**Parsing** -> **Transformation** -> **Code Generation**
### Parsing
This step is breaking down the raw codebase to AST
contains 2 steps:
- Lexical Analysis (you can call it Tokenize step)
- Syntactic Analysis

#### Lexical Analysis:
Input: Raw Code  
Process: Tokenize  
Output: Tokenizer or Lexer  
Example:  
Input: x + y  
Output:  
```javascript
[
    { type: 'Identifier', value: 'x' },
    { type: 'Operand', value: '+' },
    { type: 'Identifier', value: 'y' },
]
```
- Input: LexicalAnalysis(x)  
- Output:  
```javascript
[
    { type: 'Identifier', value: 'LexicalAnalysis' },
    { type: 'Punctuator', value: '(' },
    { type: 'Identifier', value: 'x' },
    { type: 'Punctuator', value: ')' },
]
```
#### Syntactic Analysis
Input: Tokenizer or Lexer  
Process: Parse  
Output: Abstract Syntax Tree (AST)  
Example:  
Input:   
```javascript
[
    { type: 'Identifier', value: 'x' },
    { type: 'operand', value: '+' },
    { type: 'Identifier', value: 'y' },
]
```
Output:
```javascript 
[
    { type: 'Program', body: [...]}
]
```
### Transformation
// TODO

### Code Generation
// TODO
# General Architect:
![General Architect](https://github.com/xxxle0/tiny-popiler/blob/master/Diagram.png?raw=true)
# Folder structure:
# References:
- https://github.com/hazbo/the-super-tiny-compiler
- https://www.destroyallsoftware.com/screencasts/catalog/a-compiler-from-scratch
- https://www.twilio.com/blog/abstract-syntax-trees