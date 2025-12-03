

### 中文 [README.md](README.md)English [README.md](README.md)

## Learning Path (Recommended Order)

Follow this order to learn, which aligns with the beginner's understanding logic and avoids knowledge gaps:

001-demo-HelloWorld (Program Entry & Basic Printing)

002-demo-Variable (Variable Declaration & Usage)

003-demo-Scope (Variable Scope & Memory Storage)

004-demo-NumberTypes (Numeric Types & Value Ranges)

005-demo-Bool (Boolean Type & Logical Operations)

006-demo-String (String Operations & Encoding Features)

007-demo-TypeConversion (Explicit Type Conversion Rules)

008-demo-ControlFlow (Conditional Branch Statements)

009-demo-Loop (Loops & Flow Control)

010-demo-Function (Function Definition & Closures)

011-demo-Pointer (Pointers & Memory Addresses)

012-demo-Const (Constants & Enumerators)

## Core Content of Each Demo

| No.  | Directory Name | Core Knowledge Points                                        | Learning Focus                                               |
| ---- | -------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 001  | HelloWorld     | Role of main package/function, fmt package printing (Print/Println/Printf) | Program entry rules, formatted output syntax                 |
| 002  | Variable       | Variable declaration (var/short declaration/batch declaration), blank identifier _, zero-value feature | Selection of declaration scenarios, ignoring return values with blank identifier |
| 003  | Scope          | Scope of local/global/formal parameters, variable shadowing rules, basic escape analysis | Scope boundary judgment, stack/heap memory differences       |
| 004  | NumberTypes    | Integers (signed/unsigned), floating-point numbers (precision issues), basic complex numbers | Value ranges of numeric types, avoiding floating-point precision loss |
| 005  | Bool           | Boolean values (true/false), short-circuit behavior of logical operators | Short-circuit evaluation of &&/                              |
| 006  | String         | String definition (""/``), indexing/slicing, concatenation performance comparison | UTF-8 encoding features, efficient concatenation with Builder |
| 007  | TypeConversion | Explicit conversion rules, numeric precision loss, string-numeric conversion (strconv package) | Conversion restrictions for same underlying type, strconv.Atoi/FormatInt |
| 008  | ControlFlow    | if/else if/else, switch (multi-value/no expression/fallthrough), error handling paradigm | Flexible use of switch, best practices for error handling with if |
| 009  | Loop           | for loops (basic/infinite/for-range), break/continue/goto    | Replacing while/do-while with for, appropriate use cases for goto |
| 010  | Function       | Function declaration, multiple return values, anonymous functions, closures | Variable capture in closures, error handling with multiple return values |
| 011  | Pointer        | Pointer declaration (&/*), modifying original variables via pointers, new function | Difference between pass-by-pointer and pass-by-value, memory address operations |
| 012  | Const          | Constant declaration, iota enumerator, compile-time value feature | iota auto-increment rules, differences between constants and variables |

## Execution & Debugging

### 1. Run a Single Demo

Navigate to the target demo directory and execute the go run command:

```
# Example: Run HelloWorld
cd 001-demo-HelloWorld
go run hello_world.go
```

### 2. Common Issue Resolution

- **"undefined: xxx" error**: Check the completeness of package imports and case sensitivity of variable/function names (Go is case-sensitive);

- **Type conversion error**: Ensure the converted types have the same underlying type (e.g., bool cannot be directly converted to int);

- **Abnormal closure variable values**: Capture the current value of loop variables via function parameters to avoid late binding issues.

## Learning Recommendations

**Practice First**: For each demo, first understand the comments, then modify the code (e.g., adjust loop conditions, variable types) and observe the result changes;

**Focus on Key Points**: Prioritize mastering easily confused knowledge points such as variable scope, pointer passing, closures, and string concatenation performance;

## Supporting Resources

- Not available for now