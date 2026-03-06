# Standards and Guidelines

⚠️ **CRITICAL: AI assistants MUST read and reference this file before EVERY interaction with the caracal project.**

## Overview

This folder contains the standards and best practices that **MUST** be followed throughout the caracal project. All developers, including AI assistants, **MUST** reference these documents when making code changes, adding features, or reviewing code.

**This README.md file is the PRIMARY CONTEXT for all AI chat interactions on this project.**

## Purpose

These standards ensure:
- **Consistency** across the codebase
- **Quality** and maintainability of code
- **Best practices** aligned with Go community standards
- **Efficient collaboration** between team members and AI assistants

## Documents in This Folder

### 1. golang_best_practices.md
Comprehensive guide covering:
- Code style and formatting conventions
- Naming conventions for packages, variables, functions, and constants
- Error handling patterns
- Documentation standards (godoc)
- Code organization and project structure
- Concurrency patterns
- Testing conventions
- Performance optimization tips
- Dependency management
- Common Go patterns
- Useful tools and commands

## AI Assistant Instructions

### Context Requirement
**ALL** documents in this standards folder must be read and understood before:
- Writing new code
- Modifying existing code
- Adding features or functionality
- Reviewing code changes
- Making architectural decisions

### Mandatory Compliance
The following must be adhered to at all times:

1. **Proper English**
   - AI assistants MUST use proper English (grammar, spelling, punctuation, and professional wording) in all outputs.
   - This applies to chat responses, code comments, documentation, PR descriptions, and commit messages.

2. **Code Format**
   - All Go code must be formatted using `gofmt`
   - Variable and function names must follow the naming conventions specified in golang_best_practices.md
   - Exported symbols use PascalCase, unexported use camelCase

2. **Error Handling**
   - Never ignore errors silently
   - Always wrap errors with context using `fmt.Errorf("message: %w", err)`
   - Check for errors immediately after operations

3. **Documentation**
   - All exported functions, types, and packages must have godoc comments
   - Comments must start with the name of the item being documented
   - Comments should be clear and concise

4. **Code Organization**
   - Follow the project structure guidelines
   - Use `internal/` for unexported packages
   - Use `pkg/` for exported packages
   - Keep related functionality in the same package

5. **Testing**
   - Write tests for all public functions
   - Place tests in `_test.go` files in the same package
   - Use table-driven tests for multiple test cases

6. **Project Structure**
   - Follow the recommended directory layout
   - Organize code by feature or functionality
   - Keep the project clean and maintainable

### Before Making Changes

AI assistants must:
1. ✅ **Read this file (docs/standards/README.md) before EVERY interaction**
2. ✅ Read and understand all standards documents
3. ✅ Check that proposed changes comply with all guidelines
4. ✅ Ensure code follows Go best practices
5. ✅ Validate that error handling is proper
6. ✅ Verify that documentation is complete
7. ✅ Run tests and linters before finalizing changes

### AI Usage Protocol

**This is NON-NEGOTIABLE:**
- Every AI chat session on this project MUST start by reading `docs/standards/README.md`
- Every code change MUST be validated against these standards
- If an AI assistant does not reference these standards, remind them to do so
- All code must comply with `golang_best_practices.md` guidelines
- No exceptions without explicit approval and documentation

### Common Commands Reference

```bash
# Format code
go fmt ./...

# Run static analysis
go vet ./...

# Run tests
go test ./...

# Check test coverage
go test -cover ./...

# Clean up dependencies
go mod tidy
```

## Escalation and Exceptions

If you believe an exception to these standards is necessary:
1. Document the reason clearly
2. Explain why the standard cannot be followed
3. Propose an alternative that maintains code quality
4. Request approval before implementation

## Version History

- **v1.0** (March 6, 2026) - Initial standards documentation created
  - Added golang_best_practices.md
  - Established AI assistant compliance requirements

## Contact and Questions

For questions about these standards or to propose updates, please contact the team lead or create a discussion in the project repository.

---

**Last Updated:** March 6, 2026  
**Status:** Active - Must be followed for all new code and changes

