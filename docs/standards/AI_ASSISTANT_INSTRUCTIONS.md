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

## 🔐 DEPENDENCY & LICENSE MANAGEMENT (NON-NEGOTIABLE)

### License Requirements - MANDATORY COMPLIANCE

**CRITICAL: All dependencies added to the caracal project MUST comply with these license requirements.**

**Acceptable Licenses (Permissive - Required):**
- ✅ **MIT License** - Most commonly used, highly permissive
- ✅ **Apache License 2.0** - Very permissive, patent protection included
- ✅ **BSD Licenses** (2-Clause, 3-Clause) - Permissive, very flexible
- ✅ **ISC License** - Permissive, minimal restrictions
- ✅ **Unlicense** - Public domain equivalent, most permissive
- ✅ **CC0** - Public domain dedication

**⛔ NOT ACCEPTABLE:**
- ❌ GPL (v2, v3, AGPL) - Copyleft, requires sharing source code
- ❌ LGPL - Copyleft derivative restrictions
- ❌ SSPL - Server-side public license, restrictive
- ❌ Proprietary licenses - Licensing fees or restrictions
- ❌ Custom/Unknown licenses - Must be reviewed before use
- ❌ Experimental/Beta licenses - Not production-ready

### How to Check Licenses

Before adding ANY dependency:

```bash
# Check go.mod for all dependencies
go mod graph

# View dependency information
go list -json ./...

# Check individual package license
go get -u -v [package]

# Search for known issues
curl https://raw.githubusercontent.com/golang/go/master/doc/go1.N.html | grep -i license
```

Use online tools:
- https://licenses.nuget.org/ (general reference)
- https://opensource.org/licenses/ (official license database)
- https://github.com/search?q=license:MIT (GitHub advanced search)

### Pre-Dependency Checklist

**BEFORE using any new Go package, verify:**

- [ ] Package has a clear LICENSE file
- [ ] License is MIT, Apache 2.0, BSD, ISC, or Unlicense
- [ ] License file is accessible in the repository
- [ ] No GPLv2, GPLv3, AGPL, or proprietary licenses
- [ ] Package is actively maintained
- [ ] Package has no known security vulnerabilities
- [ ] Package dependency chain is vetted (check transitional dependencies)

### Recommended Packages by Category

**Web/HTTP Frameworks:**
- ✅ `net/http` (stdlib) - MIT (no external dep needed)
- ✅ `github.com/gin-gonic/gin` - MIT
- ✅ `github.com/labstack/echo` - MIT
- ✅ `github.com/gorilla/mux` - BSD-3-Clause

**Logging:**
- ✅ `github.com/sirupsen/logrus` - MIT
- ✅ `github.com/uber-go/zap` - MIT
- ✅ `log` (stdlib) - No license needed

**JSON/Serialization:**
- ✅ `encoding/json` (stdlib) - Built-in
- ✅ `github.com/golang/protobuf` - BSD-3-Clause

**Testing:**
- ✅ `testing` (stdlib) - Built-in
- ✅ `github.com/stretchr/testify` - MIT

**Database:**
- ✅ `database/sql` (stdlib) - Built-in
- ✅ `github.com/lib/pq` (PostgreSQL) - BSD-2-Clause
- ✅ `github.com/go-sql-driver/mysql` - Mozilla Public License 2.0

**Configuration:**
- ✅ `github.com/spf13/viper` - MIT
- ✅ `github.com/joho/godotenv` - MIT

**Validation:**
- ✅ `github.com/go-playground/validator` - MIT
- ✅ `regexp` (stdlib) - Built-in

### Adding Dependencies

When adding a new dependency:

```bash
# 1. Add the dependency
go get github.com/package/name@latest

# 2. Tidy the module file
go mod tidy

# 3. Verify the license in go.mod
cat go.mod

# 4. Check the actual LICENSE file in vendor or module cache
find $GOPATH/pkg/mod -name "LICENSE" -path "*package*"

# 5. Add to your dependency tracking (if using)
# Update docs/dependencies.md with license info
```

### Dependency Documentation

**REQUIRED:** Maintain a `docs/DEPENDENCIES.md` file listing all major dependencies and their licenses:

Example format:
```markdown
# Dependencies and Licenses

| Package | License | Purpose |
|---------|---------|---------|
| github.com/sirupsen/logrus | MIT | Structured logging |
| github.com/spf13/viper | MIT | Configuration management |
```

### Validation Commands

Add these to your CI/CD pipeline:

```bash
# Check for GPL licenses (will fail if found)
go mod graph | grep -i "gpl" && exit 1

# Verify all dependencies are accessible
go mod verify

# List all dependencies with their modules
go mod graph

# Update and tidy
go mod tidy
```

### What to Do If You Find GPL

If you discover a GPL-licensed package in use:

1. **STOP immediately** - Do not commit or use it
2. **Document the issue** - Note where it was used
3. **Find alternatives** - Use the "Recommended Packages" section above
4. **Report** - Inform the team lead
5. **Replace** - Switch to a permissive-licensed alternative

### Escalation for License Questions

If you're unsure about a license:

1. Research on https://opensource.org/licenses/
2. Check the package README for license info
3. Look at the actual LICENSE file
4. Ask the team lead or legal/compliance contact
5. **DO NOT proceed without clarity**

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

❌ **Using GPL-Licensed Packages**
```go
import "github.com/gpl-licensed/package"  // FORBIDDEN!
```

❌ **Adding Dependencies Without License Check**
```bash
go get github.com/unknown/package  # WRONG! Must verify license first
```

❌ **Ignoring License Warnings**
- Do NOT ignore GPL, AGPL, or proprietary licenses
- Do NOT use "it probably won't matter" as a reason
- Do NOT skip the license verification step

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
8. **ALWAYS verify licenses** - MIT/Apache 2.0/BSD/ISC only, NEVER GPL
9. **ALWAYS use recommended packages** - From the approved list
10. **ALWAYS check transitive dependencies** - Ensure entire dependency tree is compliant

### License Compliance is NON-NEGOTIABLE

**Every single dependency added to caracal must:**
- ✅ Have MIT, Apache 2.0, BSD, ISC, or Unlicense
- ✅ Be from the approved recommended packages list
- ✅ Have been license-verified before adding
- ✅ Be documented in go.mod with clear licensing notes
- ❌ NEVER be GPL, AGPL, LGPL, SSPL, or proprietary
- ❌ NEVER be added without license verification

**These standards exist to maintain code quality, consistency, collaboration, and legal compliance.**

---

**Version:** 1.0  
**Last Updated:** March 6, 2026  
**Status:** ACTIVE - Mandatory for all AI assistants  
**Compliance:** Required - No exceptions without explicit approval

