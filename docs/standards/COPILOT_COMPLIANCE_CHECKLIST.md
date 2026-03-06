# GitHub Copilot: Caracal Project Standards Checklist

## 🤖 AI Assistant Compliance Protocol

**This checklist MUST be completed before EVERY interaction on the caracal project.**

---

## ✅ PRE-INTERACTION CHECKLIST

### Step 1: Read Standards (Required)
- [ ] Read `docs/standards/README.md` 
- [ ] Read `docs/standards/AI_ASSISTANT_INSTRUCTIONS.md`
- [ ] Read `docs/standards/golang_best_practices.md`
- [ ] Read `docs/standards/LICENSES.md`
- [ ] Keep `docs/standards/LICENSE_QUICK_REFERENCE.md` available

### Step 2: Understand License Requirements
- [ ] Approved licenses: MIT, Apache 2.0, BSD-3-Clause, BSD-2-Clause, ISC, Unlicense
- [ ] Forbidden licenses: GPLv2, GPLv3, AGPL, LGPL, SSPL, Proprietary
- [ ] Pre-approved packages memorized from recommended list
- [ ] License verification process understood

### Step 3: Verify Project Structure
- [ ] Understand: main.go, go.mod, go.sum at root
- [ ] Understand: config/ package exists
- [ ] Understand: docs/standards/ contains all policies
- [ ] Understand: internal/ for unexported, pkg/ for exported

### Step 4: Ready to Work
- [ ] All standards understood
- [ ] License requirements clear
- [ ] Verification procedures known
- [ ] Red flags identified
- [ ] Ready to ensure compliance

---

## 📋 PRE-CODE CHECKLIST

Before writing ANY code:
- [ ] Understand the requirement completely
- [ ] Check if similar code exists
- [ ] Plan to follow Go best practices
- [ ] Identify any dependencies needed
- [ ] Prepare to verify dependency licenses

---

## 🔐 DEPENDENCY CHECKLIST

Before adding ANY dependency:

**License Verification:**
- [ ] Package has accessible LICENSE file
- [ ] License is MIT, Apache 2.0, BSD, ISC, or Unlicense
- [ ] NOT GPL, AGPL, LGPL, SSPL, or proprietary
- [ ] Check transitive dependencies with `go mod graph`

**Package Verification:**
- [ ] Package is from pre-approved list (preferred)
- [ ] Package is actively maintained
- [ ] No known security vulnerabilities
- [ ] Package solves the exact problem

**Approved to Add:**
- [ ] Ready to run: `go get [package]`
- [ ] Ready to run: `go mod tidy`
- [ ] Ready to document in go.mod

---

## 💻 CODE QUALITY CHECKLIST

Before finalizing ANY code:

**Formatting & Linting:**
- [ ] Code passes: `go fmt ./...`
- [ ] Code passes: `go vet ./...`
- [ ] Code passes: `go test ./...`
- [ ] Code passes: `go test -cover ./...`

**Documentation:**
- [ ] All exported functions have godoc comments
- [ ] All exported types have godoc comments
- [ ] All exported packages have package comments
- [ ] Comments are clear and complete

**Error Handling:**
- [ ] All errors are checked
- [ ] Errors are wrapped with context: `fmt.Errorf("message: %w", err)`
- [ ] No silent error ignoring with `_`
- [ ] Error messages are helpful

**Testing:**
- [ ] Public functions have tests
- [ ] Tests are in `_test.go` files
- [ ] Table-driven test pattern used
- [ ] Coverage is >80% where practical

---

## 🚫 RED FLAGS - STOP IF ANY APPLY

❌ **STOP IF ANY OF THESE:**
- GPL-licensed package being added
- License is unknown or unclear
- License is proprietary or experimental
- No LICENSE file is accessible
- Not in pre-approved packages list and license not verified
- "It probably won't matter" attitude
- Skipping license verification
- Using `_ = doSomething()` to ignore errors
- No godoc comments on exported symbols
- No tests for public functions

**IF RED FLAG FOUND:**
1. Stop all work
2. Document the issue
3. Report to team lead
4. Find alternative solution
5. Restart with compliance

---

## ✨ COMPLIANCE CONFIRMATION

### Before Submitting Code

Confirm:
- [ ] All standards have been followed
- [ ] All licenses have been verified
- [ ] All errors are handled properly
- [ ] All code is documented
- [ ] All tests pass
- [ ] All tools pass (fmt, vet, test)
- [ ] Ready for production

### Response Format

Include in every response about code:
```
📚 STANDARDS REFERENCE
- [Relevant section from standards]

✅ COMPLIANCE CHECKLIST
- [x] All standards followed
- [x] Code formatted
- [x] Tests passing
- [x] Documentation complete

🔐 LICENSE VERIFICATION (if applicable)
- [x] Licenses verified
- [x] No GPL found
- [x] Using approved packages

📝 IMPLEMENTATION
[Description]
```

---

## 🎯 KEY REMINDERS

**ALWAYS:**
1. ✅ Read standards first
2. ✅ Verify licenses before adding dependencies
3. ✅ Use pre-approved packages
4. ✅ Follow Go best practices
5. ✅ Handle errors properly
6. ✅ Document everything
7. ✅ Test thoroughly
8. ✅ Run validation tools
9. ✅ Escalate violations
10. ✅ Check for GPL

**NEVER:**
1. ❌ Skip license verification
2. ❌ Add GPL-licensed packages
3. ❌ Ignore errors silently
4. ❌ Skip documentation
5. ❌ Skip testing
6. ❌ Assume licenses are safe
7. ❌ Use proprietary packages
8. ❌ Bypass standards
9. ❌ Commit without validation
10. ❌ Use "it won't matter" excuse

---

## 📞 ESCALATION CONTACTS

If encountering issues:

**GPL or Proprietary License Found:**
- Document the package and license
- Report to team lead immediately
- Find alternative from approved list
- Do NOT proceed with non-compliant package

**Unclear License:**
- Research on https://opensource.org/licenses/
- Check package README and LICENSE file
- Ask team lead for clarification
- Do NOT proceed without clarity

**Not in Pre-Approved List:**
- Verify license is on approved list
- If yes, document and proceed
- If no, find alternative or ask team lead
- Update pre-approved list if appropriate

---

## 📊 QUICK STATS

**Approved Licenses:** 5 primary + 4 secondary
- ✅ MIT (Most Common)
- ✅ Apache 2.0 (Enterprise)
- ✅ BSD-3-Clause (Academic)
- ✅ BSD-2-Clause (Minimal)
- ✅ ISC (Simple)

**Forbidden Licenses:** 6
- ❌ GPLv2, GPLv3, AGPL, LGPL, SSPL, Proprietary

**Pre-Approved Packages:** 30+
- Logging, Config, Web, CLI, Testing, Database, Validation, etc.

**Standard Documents:** 5 files, 1,483 lines
- README.md, AI_ASSISTANT_INSTRUCTIONS.md, golang_best_practices.md, LICENSES.md, LICENSE_QUICK_REFERENCE.md

---

## ✅ SIGN-OFF

**I, GitHub Copilot, acknowledge:**

1. ✅ I have read and understand all standards documents
2. ✅ I understand the non-negotiable license requirements
3. ✅ I will verify licenses before adding ANY dependency
4. ✅ I will NEVER add GPL, AGPL, LGPL, SSPL, or proprietary packages
5. ✅ I will ALWAYS use pre-approved packages when available
6. ✅ I will ALWAYS escalate license violations immediately
7. ✅ I will ALWAYS follow Go best practices
8. ✅ I will ALWAYS ensure code quality and test coverage
9. ✅ I will ALWAYS document my work thoroughly
10. ✅ I will ALWAYS comply with these standards

**Compliance Status:** ✅ ACTIVE - READY FOR WORK

---

**Version:** 1.0  
**Last Updated:** March 6, 2026  
**Status:** MANDATORY FOR ALL AI INTERACTIONS  
**Enforced By:** GitHub Copilot & Development Team

*This checklist must be completed before every interaction on the caracal project.*

