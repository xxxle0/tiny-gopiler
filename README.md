# Introduction:
I want to do some exp to understand how the compiler can check the code syntax. I know it will build an AST (Abstract Syntax Tree) from the raw code then traverse the tree to check the syntax is correct or not? also generate anything from the Tree but yeah need to get my hand on dirty to find out.
#### What I cannot create, I do not understand
# What?
How the compiler separate the expression, statement, function?
We know that a file of code contains many statement and expression
## What is expression?
Expression is a snippet of code that return the value
For Example:
- 1 + 2 (Compute expression)
- "Hello World" (String literal)
- 2 (Number literal)

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
Input: **x + y**  
Output:  
```javascript
[
    { type: 'Identifier', value: 'x' },
    { type: 'Operand', value: '+' },
    { type: 'Identifier', value: 'y' },
]
```
- Input: **LexicalAnalysis(x)**  
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
    { type: 'Operand', value: '+' },
    { type: 'Identifier', value: 'y' },
]
```
Output:
```javascript 
[
    { type: 'Program', body: [
        {
            type: "BinaryExpression",
            params: [
                {
                    type: "Identifier",
                    Value: "x"
                },
                {
                    type: "Identifier",
                    value: "y",
                }
            ]
        },
    ]}
]
```
*Notes: The reason why LISP can be used for build a compiler is LISP syntax is much more easy to struct AST tree   
`(+ 1 2)` -> the operand always be placed at the head of expression*

### Transformation
This step takes the output from `Parsing` step then add/delete/update nodes to transform new structure AST is suitable to later usage

### Code Generation
Traverse the Transformed AST 
# General Architect:
![General Architect](https://github.com/xxxle0/tiny-gopiler/blob/master/Diagram.png?raw=true)
*https://whimsical.com/compiler-under-the-hood-D9jGFbG6JPfvhGLm2rEBDP*
# Folder structure:
# References:
- https://github.com/hazbo/the-super-tiny-compiler
- https://www.destroyallsoftware.com/screencasts/catalog/a-compiler-from-scratch
- https://www.twilio.com/blog/abstract-syntax-trees