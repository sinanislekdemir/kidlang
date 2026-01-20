# KidLang Test Suite

This directory contains comprehensive tests for KidLang features.

## Running Tests

Execute the test suite with:

```bash
./test.sh
```

The script will:
1. Build the KidLang interpreter
2. Run all `.kid` test files
3. Compare output against `.expected` files
4. Report pass/fail status with colors
5. Return the number of failed tests as exit code

## Test Files

Each test consists of:
- `XX_testname.kid` - The test script
- `XX_testname.expected` - Expected output
- `XX_testname.setup` (optional) - Setup script to run before test
- `XX_testname.input` (optional) - Input to pipe to the program (for `ask` tests)

### Current Tests

1. **01_basic_print** - Basic print functionality
2. **02_box_variables** - Box variable operations
3. **03_math_operations** - Arithmetic operations (+, -, *, /, %)
4. **04_math_functions** - Math functions (sqrt, abs, sqr) with negative numbers
5. **05_string_operations** - String concatenation, repetition, indexing
6. **06_conditionals** - If/then/end conditionals with comparisons
7. **07_goto_labels** - Labels and goto statements
8. **08_stack_basic** - Basic stack operations with numeric and string keys
9. **09_stack_file_read** - Reading file into stack with auto-typing
10. **10_stack_file_write** - Writing stack to file
11. **11_file_operations** - Basic file read/write operations
12. **12_logical_operators** - Logical AND/OR operators
13. **13_turkish** - Turkish language support (TR)
14. **14_german** - German language support (DE)
15. **15_finnish** - Finnish language support (FI)
16. **16_complex_and** - Multiple AND conditions in single statement
17. **17_complex_or** - Multiple OR conditions in single statement
18. **18_mixed_and_or** - Mixed AND/OR with operator precedence
19. **19_type_comparisons** - String to number comparisons
20. **20_numeric_strings** - Numeric string operations and comparisons
21. **21_not_operator** - NOT operator (!=) with complex conditions
22. **22_mixed_types** - Integer, float, and string type mixing
23. **23_edge_cases** - Zero and negative number edge cases
24. **24_boolean_comparisons** - Boolean values in conditions
25. **25_long_chains** - Very long condition chains (5+ conditions)
26. **26_case_sensitivity** - Case-insensitive string comparisons
27. **27_ask_input** - User input with ask command (numeric)
28. **28_ask_text** - User input with ask command (text)
29. **29_ask_calculation** - Multiple ask inputs with calculation
30. **30_readline** - Reading file line by line
31. **31_seek** - Seeking to specific line in file
32. **32_sleep** - Sleep/delay command
33. **33_advanced_math** - Trigonometric and logarithmic functions
34. **34_random_now** - Random number and timestamp functions
35. **35_xor_operator** - XOR bitwise operator
36. **36_nested_conditionals** - Nested if/then/end blocks
37. **37_division** - Integer division
38. **38_float_division** - Floating point division
39. **39_empty_stack** - Stack operations from empty state
40. **40_multiple_labels** - Multiple goto labels and jumps
41. **41_exec_command** - Execute shell commands
42. **42_newline_char** - Newline special variable
43. **43_all_comparisons** - All comparison operators (=, !=, <, >, <=, >=)
44. **45_implicit_print** - Implicit printing (text without command)

## Test Coverage

**Current Status:** 47/47 tests passing (100% coverage) ✓

All major KidLang features are tested including:
- Core commands (print, ask, exec, sleep)
- Variables (box) and stacks  
- File operations (open, close, read, write, readline, seek)
- Math operations and all functions
- String operations
- All comparison operators
- Conditionals with complex AND/OR chains
- Labels and goto
- Special variables (random, now, \n)
- Implicit printing
- Multi-language support (EN, TR, DE, FI)
- **README examples validated**

## Adding New Tests

To add a new test:

1. Create `XX_testname.kid` with your test code
2. Create `XX_testname.expected` with the expected output
3. (Optional) Create `XX_testname.setup` if setup is needed
4. Run `./test.sh` to verify

## Exit Codes

- `0` - All tests passed
- `N` - N tests failed (1-29)
