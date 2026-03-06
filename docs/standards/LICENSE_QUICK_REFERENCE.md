# Quick Reference: License Compliance Checklist

**Print this. Use this. EVERY TIME.**

---

## Before Adding ANY Dependency

### ✅ VERIFICATION CHECKLIST

- [ ] I have read `docs/standards/LICENSES.md`
- [ ] I have identified the package and its license
- [ ] License is **MIT**, **Apache 2.0**, **BSD**, **ISC**, or **Unlicense**
- [ ] License is NOT GPL, AGPL, LGPL, SSPL, or proprietary
- [ ] Package is from the approved packages list (if available)
- [ ] I checked transitional dependencies with `go mod graph`
- [ ] Package is actively maintained
- [ ] No known security vulnerabilities

### ❌ IMMEDIATE REJECTION IF

- [ ] License is GPL (v2, v3, AGPL) - STOP
- [ ] License is LGPL or SSPL - STOP
- [ ] License is proprietary or unknown - STOP
- [ ] No LICENSE file is accessible - STOP
- [ ] License is unclear or experimental - STOP

---

## Quick License Check

```bash
# 1. Get the package
go get github.com/owner/package@latest

# 2. Find the license
go list -json github.com/owner/package | grep -i license

# 3. Or navigate to GitHub and look for LICENSE file
# 4. Read the license text
# 5. Cross-reference with approved licenses in docs/standards/LICENSES.md
```

---

## Approved Licenses (Use These)

| License | Risk | Use Case |
|---------|------|----------|
| 🟢 MIT | None | General purpose, most flexible |
| 🟢 Apache 2.0 | None | Enterprise, has patent clause |
| 🟢 BSD-3-Clause | None | Academic, corporate software |
| 🟢 BSD-2-Clause | None | Minimal restrictions |
| 🟢 ISC | None | Simple, lightweight |
| 🟡 Unlicense | Low | Public domain, uncommon |

---

## Forbidden Licenses (NEVER Use)

| License | Risk | Reason |
|---------|------|--------|
| 🔴 GPLv2 | CRITICAL | Copyleft - must open-source |
| 🔴 GPLv3 | CRITICAL | Copyleft - must open-source |
| 🔴 AGPL | CRITICAL | Network copyleft - must open-source |
| 🔴 LGPL | HIGH | Weak copyleft - dependencies restricted |
| 🔴 SSPL | HIGH | Restrictive - unclear enforcement |
| 🔴 Proprietary | CRITICAL | Licensing fees or authorization required |
| 🔴 Unknown | CRITICAL | Cannot verify compliance |

---

## Recommended Packages (Pre-Approved)

### Most Common Categories

**Logging:**
- ✅ github.com/sirupsen/logrus - MIT
- ✅ go.uber.org/zap - MIT

**Config:**
- ✅ github.com/spf13/viper - MIT
- ✅ github.com/joho/godotenv - MIT

**Web/HTTP:**
- ✅ github.com/gin-gonic/gin - MIT
- ✅ github.com/gorilla/mux - BSD-3-Clause

**Testing:**
- ✅ github.com/stretchr/testify - MIT

**Database:**
- ✅ github.com/lib/pq - BSD-2-Clause (PostgreSQL)
- ✅ github.com/go-sql-driver/mysql - Mozilla Public License

**CLI:**
- ✅ github.com/spf13/cobra - Apache 2.0
- ✅ github.com/urfave/cli - MIT

**More options:** See `docs/standards/LICENSES.md`

---

## If You Find a GPL License

**DO THIS IMMEDIATELY:**

1. **STOP** - Do not add the package
2. **DOCUMENT** - Note what you found
3. **REPORT** - Tell the team lead
4. **FIND ALTERNATIVE** - Use approved packages list
5. **REPLACE** - Update go.mod to use compliant package
6. **VERIFY** - Run license check again

---

## When in Doubt

**STOP. DO NOT PROCEED.**

1. Research the license on https://opensource.org/licenses/
2. Read the actual LICENSE file
3. Ask the team lead for clarification
4. Wait for explicit approval
5. Only then proceed

---

## Red Flags 🚩

❌ "It probably won't matter"  
❌ "Let's use it anyway"  
❌ "We can figure it out later"  
❌ "Just this one GPL package"  
❌ "Unknown license but it looks safe"  

**ANY of these = WRONG. Stop and escalate.**

---

## Tools to Help

```bash
# Check all dependencies
go mod graph

# Verify package info
go list -json github.com/package/name

# Tidy up
go mod tidy

# Find packages by license
# Search "MIT" or "Apache 2.0" on GitHub
# Filter by license on pkg.go.dev
```

---

## Remember

✅ **This is NON-NEGOTIABLE**  
✅ **Every single dependency matters**  
✅ **Verify BEFORE adding**  
✅ **Use recommended packages when available**  
✅ **Check transitive dependencies**  
✅ **Escalate when unsure**  
✅ **Never use GPL or proprietary licenses**  

---

**Version:** 1.0  
**Last Updated:** March 6, 2026  
**Status:** MANDATORY - Print this and keep it handy

