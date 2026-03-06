# AI Assistant Instructions for Caracal Project

## ⚠️ CRITICAL PROTOCOL

**EVERY AI assistant interaction with the caracal project MUST:**
1. Read `docs/standards/README.md` FIRST
2. Read this file (`AI_ASSISTANT_INSTRUCTIONS.md`) SECOND
3. Read `golang_best_practices.md` THIRD
4. Only then proceed with any work

**Failure to follow this protocol is unacceptable.**

---

## Pre-Work Checklist

Before writing ANY code, adding ANY feature, or making ANY changes, verify:

- [ ] I have read `docs/standards/README.md`
- [ ] I have read `docs/standards/AI_ASSISTANT_INSTRUCTIONS.md` (this file)
- [ ] I have read `docs/standards/golang_best_practices.md`
- [ ] I understand all mandatory compliance requirements
- [ ] I am prepared to follow Go best practices
- [ ] I will validate all changes before submission

---

## Mandatory Requirements

### 1. Always Reference Standards

When responding to user requests:
- **Always** cite the relevant standard or guideline
- **Always** explain how the proposed solution aligns with standards
- **Always** validate against `golang_best_practices.md`

Example response format:
```
Per golang_best_practices.md - Error Handling section:
✅ All errors are wrapped with context
✅ Errors are checked immediately after operations
✅ No silent error ignoring
```

### 2. Code Quality Standards

Every code change must meet these criteria:

**Formatting**
- Must pass `go fmt`
- Must pass `go vet`
- Lines under 80 characters (where practical)
- Proper indentation and spacing

**Naming**
- Exported symbols: PascalCase
- Unexported symbols: camelCase
- Variables and functions: descriptive, concise names
- Constants: PascalCase (exported), camelCase (unexported)

**Error Handling**
- Never ignore errors silently
- Always wrap errors: `fmt.Errorf("context: %w", err)`
- Check errors immediately after operations
- Provide helpful error messages

**Documentation**
- All exported functions must have godoc comments
- All exported types must have godoc comments
- All exported packages must have package-level comments
- Comments must start with the symbol name
- Comments should be clear and complete

**Testing**
- All public functions must have tests
- Tests in `_test.go` files, same package
- Use table-driven tests for multiple cases
- Aim for >80% code coverage

### 3. Project Structure Compliance

Always follow the recommended structure:
```
caracal/
├── main.go              # Entry point
├── go.mod
├── go.sum
├── internal/            # Unexported packages
│   ├── auth/
│   └── database/
├── pkg/                 # Exported packages
│   ├── api/
│   └── models/
├── cmd/                 # CLI tools
│   └── cli/
├── docs/                # Documentation
│   └── standards/       # This standards folder
└── tests/               # Integration tests
```

### 4. Error Handling Pattern

Every error scenario must follow this pattern:

```go
// ❌ WRONG
result, _ := someFunction()

// ❌ WRONG
if err != nil {
    log.Fatal(err)
}

// ✅ CORRECT
if err != nil {
    return fmt.Errorf("failed to process data: %w", err)
}
```

### 5. Documentation Pattern

Every exported symbol must have godoc:

```go
// ❌ WRONG - No comment
func ProcessData(input string) (string, error) {
    // ...
}

// ✅ CORRECT
// ProcessData transforms input data according to business rules.
// It returns an error if the input format is invalid.
func ProcessData(input string) (string, error) {
    // ...
}
```

---

## Workflow

### When Starting Work

1. **Read All Standards** (non-negotiable)
2. **Understand the Request** - Ask clarifying questions if needed
3. **Plan the Implementation** - Consider project structure and best practices
4. **Write Code** - Follow all standards and best practices
5. **Validate** - Run tests, linters, formatters
6. **Review** - Ensure compliance with all guidelines
7. **Document** - Add necessary documentation and comments
8. **Submit** - Provide context showing standards compliance

### When Making Changes

```
1. Read relevant standards section ← ALWAYS FIRST
2. Gather context from existing code
3. Plan changes with standards in mind
4. Implement changes
5. Validate with tools:
   - go fmt ./...
   - go vet ./...
   - go test ./...
   - go test -cover ./...
6. Document changes
7. Explain compliance with standards
```

### When Reviewing Code

Use this checklist:

- [ ] Does it follow naming conventions?
- [ ] Is all error handling proper?
- [ ] Are there godoc comments for all exported symbols?
- [ ] Does it pass `go fmt`, `go vet`, `go test`?
- [ ] Is the code organized per project structure?
- [ ] Are there tests for public functions?
- [ ] Is documentation clear and complete?

---

## Communication Template

When responding about code work:

```
📋 STANDARDS REFERENCE
- [Section from golang_best_practices.md]
- [Section from README.md]

✅ COMPLIANCE CHECKLIST
- [x] Follows naming conventions
- [x] Proper error handling
- [x] Complete documentation
- [x] Tests included
- [x] Code formatted with gofmt
- [x] Passes go vet

📝 IMPLEMENTATION
[Description of what was done and why]

🔍 VALIDATION
[Tools run and results]

📚 CONTEXT
[Explanation of how this aligns with standards]
```

---

## Common Tasks with Standards

### Adding a New Function

```go
// FunctionName describes what the function does.
// It should include parameters, return values, and possible errors.
// Include examples if the function is complex.
func FunctionName(param1 string, param2 int) (string, error) {
    if param1 == "" {
        return "", fmt.Errorf("param1 cannot be empty")
    }
    
    // Implementation
    
    return result, nil
}
```

**Standards Applied:**
- ✅ godoc comment starting with function name
- ✅ PascalCase for exported function
- ✅ Error handling with context
- ✅ Clear description

### Adding a New Package

```go
// Package mypackage provides functionality for [purpose].
// 
// Example usage:
//  result, err := mypackage.DoSomething(input)
//  if err != nil {
//      log.Fatal(err)
//  }
package mypackage
```

**Standards Applied:**
- ✅ Package-level godoc comment
- ✅ Clear purpose description
- ✅ Usage example

### Adding Tests

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {"valid input", "test", "result", false},
        {"empty input", "", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := FunctionName(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("FunctionName() error = %v, wantErr %v", err, tt.wantErr)
            }
            if got != tt.want {
                t.Errorf("FunctionName() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

**Standards Applied:**
- ✅ Table-driven test pattern
- ✅ Clear test cases
- ✅ Proper error checking
- ✅ Descriptive test names

---

## Tools to Always Use

Before finalizing ANY code:

```bash
# Format code
go fmt ./...

# Run static analysis
go vet ./...

# Run all tests
go test ./...

# Check coverage
go test -cover ./...

# Clean dependencies
go mod tidy

# Run all at once
go fmt ./... && go vet ./... && go test -cover ./...
```

---

## Red Flags - NEVER DO THIS

❌ **Silent Error Ignoring**
```go
_ = doSomething()  // WRONG!
```

❌ **Unexported Symbols Without Lowercase Start**
```go
func MyFunction() {}  // Should be myFunction() if unexported
```

❌ **Missing Error Handling**
```go
result, _ := getValue()  // WRONG!
```

❌ **No Documentation on Exported Symbols**
```go
func DoSomething() {}  // Missing godoc comment
```

❌ **Poor Error Messages**
```go
return fmt.Errorf("error")  // Lacks context
```

❌ **No Tests for Public Functions**
```go
// Function has no _test.go file or test cases
```

---

## Escalation Protocol

If you encounter situations where standards cannot be followed:

1. **Document the situation clearly**
   - What is the requirement?
   - Why does the standard not apply?
   - What is the proposed alternative?

2. **Explain the business justification**
   - Why is this exception necessary?
   - What are the trade-offs?

3. **Propose a compliant alternative**
   - How can we maintain code quality?
   - How can we minimize deviation?

4. **Request explicit approval**
   - Before proceeding with the exception
   - Document the approval for future reference

---

## Response Template for AI Assistants

Use this template for all responses involving code work:

```
# Task: [Description of what was requested]

## 📚 Standards Reference
Reading from: docs/standards/README.md, golang_best_practices.md

**Relevant Guidelines:**
- [Guideline 1 from standards]
- [Guideline 2 from standards]

## ✅ Validation Checklist
- [x] Code formatted with gofmt
- [x] Passes go vet analysis
- [x] All errors properly handled
- [x] All exported symbols documented
- [x] Tests provided
- [x] Follows naming conventions

## 📝 Implementation

[Code changes or explanation]

## 🔍 How This Complies with Standards

[Explanation of compliance]

## 💻 Validation Commands Run

```bash
go fmt ./...
go vet ./...
go test -cover ./...
```

[Results of commands]
```

---

## Summary

As an AI assistant working on caracal:

1. **ALWAYS read standards first** - Non-negotiable
2. **ALWAYS validate against guidelines** - Before any code
3. **ALWAYS provide context** - Explain compliance
4. **ALWAYS run validation tools** - Prove quality
5. **ALWAYS document everything** - Exported symbols need godoc
6. **ALWAYS handle errors properly** - With context
7. **ALWAYS test your code** - At least for public functions

**These standards exist to maintain code quality, consistency, and collaboration.**

---

**Version:** 1.0  
**Last Updated:** March 6, 2026  
**Status:** ACTIVE - Mandatory for all AI assistants  
**Compliance:** Required - No exceptions without explicit approval

