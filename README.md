# Simple fictitious scripting language (FSL) interpeter

## Project description
This is a sample of test code which contains basic interpeter.
It should have possibility to read and process json input data in format represented in the following sample:
```json
{
  "var1":1,
  "var2":2,
  
  "init": [
    {"cmd" : "#setup" }
  ],
  
  "setup": [
    {"cmd":"update", "id": "var1", "value":3.5},
    {"cmd":"print", "value": "#var1"},
    {"cmd":"#sum", "id": "var1", "value1":"#var1", "value2":"#var2"},
    {"cmd":"print", "value": "#var1"},
    {"cmd":"create", "id": "var3", "value":5},
    {"cmd":"delete", "id": "var1"},
    {"cmd":"#printAll"}
  ],
  
  "sum": [
      {"cmd":"add", "id": "$id", "operand1":"$value1", "operand2":"$value2"}
  ],

  "printAll":
  [
    {"cmd":"print", "value": "#var1"},
    {"cmd":"print", "value": "#var2"},
    {"cmd":"print", "value": "#var3"}
  ]
}
```
Expected response is:
```
3.5
5.5
undefined
2
5
```

Other requirements:
- The finished project must support receiving multiple FSL scripts. 
- Functions and variables must persist between FSL scripts. 
- Resolve conflicts by overwriting existing variables or functions.
- The system will create a representation of the script processed. 
- The init function is immediately called after each script is processed.
- The input is a JSON object of named variables and named functions
- Variables are defined as a key value pair.
- References to variables are preceded by a hash mark (#).
- A function is an array of command objects.
- An attribute called “cmd” is required and will define which operation to perform.
- All parameters passed to a function are referenced by a $.
- Function calls are defined in the “cmd” attribute by preceding the function name with a hash mark (#).

## Build & run
Installed Golang of any version is required. The project was written with Go v1.18. Before first run vendored packages should be populated. It can be done by running the following command: `go mod tidy` in the project root directory.

Use `make test` to run unit tests.

Use `make` to run code with test sample
