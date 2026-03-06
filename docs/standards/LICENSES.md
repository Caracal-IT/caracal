# License Policy for Caracal Project

## 🔐 LICENSING STANDARDS (NON-NEGOTIABLE)

**The caracal project uses ONLY permissive, open-source licenses. All dependencies MUST comply.**

---

## ✅ APPROVED LICENSES

### Tier 1 - Highly Recommended (Use These First)

| License          | Restrictiveness        | Best For                      | Example Projects               |
|------------------|------------------------|-------------------------------|--------------------------------|
| **MIT**          | Minimal                | General use, most packages    | Lodash, Rails, Vue.js          |
| **Apache 2.0**   | Low with patent clause | Enterprise, patent protection | Kubernetes, Android            |
| **BSD-3-Clause** | Low                    | Academic, enterprise          | Django, NumPy                  |
| **BSD-2-Clause** | Very Low               | Simple permissive use         | Go standard library components |
| **ISC**          | Minimal                | Lightweight, simple           | Node.js, npm                   |

### Tier 2 - Acceptable (If Tier 1 Not Available)

| License       | Notes                      |
|---------------|----------------------------|
| **Unlicense** | Public domain equivalent   |
| **CC0 1.0**   | Public domain dedication   |
| **WTFPL**     | Very permissive (informal) |
| **Zlib**      | Permissive for data        |

---

## ❌ FORBIDDEN LICENSES

### Copyleft Licenses (NEVER Use)

| License   | Why Forbidden                   | Risk                             |
|-----------|---------------------------------|----------------------------------|
| **GPLv2** | Requires source sharing         | Must open-source our code        |
| **GPLv3** | Copyleft with more restrictions | Must open-source our code        |
| **AGPL**  | Network copyleft                | Must open-source for network use |
| **LGPL**  | Weak copyleft                   | Dependency restrictions          |
| **SSPL**  | Restrictive, controversial      | License violations               |

### Proprietary Licenses (NEVER Use)

- Commercial licenses with fees
- Proprietary software licenses
- Custom restrictive licenses
- Licenses requiring payment or authorization

### Experimental/Unstable (NEVER Use)

- Pre-release licenses
- Deprecated licenses
- Custom-made licenses
- Licenses with "experimental" status

---

## 📋 APPROVED PACKAGES BY CATEGORY

### Web & HTTP

**✅ Approved**
- `net/http` (stdlib) - Built-in, no license concern
- `github.com/gin-gonic/gin` - MIT
- `github.com/labstack/echo` - MIT
- `github.com/gorilla/mux` - BSD-3-Clause
- `github.com/gorilla/websocket` - BSD-2-Clause
- `github.com/valyala/fasthttp` - MIT

**❌ Blocked**
- `beego` - Contains Apache 2.0 (if subcomponent uses GPL)
- Any package with unclear licensing

### Logging & Monitoring

**✅ Approved**
- `log` (stdlib) - Built-in
- `github.com/sirupsen/logrus` - MIT
- `go.uber.org/zap` - MIT
- `github.com/rs/zerolog` - MIT

**❌ Blocked**
- Packages with GPL dependencies

### Database & Storage

**✅ Approved**
- `database/sql` (stdlib) - Built-in
- `github.com/lib/pq` - BSD-2-Clause (PostgreSQL)
- `github.com/go-sql-driver/mysql` - Mozilla Public License 2.0
- `github.com/mattn/go-sqlite3` - MIT
- `github.com/mongodb/mongo-go-driver` - Apache 2.0

**❌ Blocked**
- GPL-licensed database drivers

### JSON & Serialization

**✅ Approved**
- `encoding/json` (stdlib) - Built-in
- `encoding/xml` (stdlib) - Built-in
- `github.com/golang/protobuf` - BSD-3-Clause
- `github.com/json-iterator/go` - MIT

**❌ Blocked**
- Custom serialization with GPL

### Configuration Management

**✅ Approved**
- `github.com/spf13/viper` - MIT
- `github.com/joho/godotenv` - MIT
- `github.com/kelseyhightower/envconfig` - MIT

**❌ Blocked**
- Configuration tools with GPL dependencies

### Testing & Mocking

**✅ Approved**
- `testing` (stdlib) - Built-in
- `github.com/stretchr/testify` - MIT
- `github.com/stretchr/objx` - MIT

**❌ Blocked**
- GPL-licensed testing frameworks

### Data Validation

**✅ Approved**
- `github.com/go-playground/validator` - MIT
- `github.com/asaskevich/govalidator` - MIT

**❌ Blocked**
- Validation libraries with GPL dependencies

### CLI & Command-Line Tools

**✅ Approved**
- `flag` (stdlib) - Built-in
- `github.com/spf13/cobra` - Apache 2.0
- `github.com/urfave/cli` - MIT

**❌ Blocked**
- CLI frameworks with GPL dependencies

### Cryptography & Security

**✅ Approved**
- `crypto` (stdlib) - Built-in
- `golang.org/x/crypto` - BSD-3-Clause
- `golang.org/x/sys` - BSD-3-Clause

**❌ Blocked**
- Crypto libraries with GPL or proprietary licenses

### Utilities & Helpers

**✅ Approved**
- `github.com/fatih/color` - MIT
- `github.com/pkg/errors` - BSD-2-Clause
- `github.com/google/uuid` - BSD-3-Clause

**❌ Blocked**
- Utility packages with GPL dependencies

---

## 🔍 HOW TO VERIFY LICENSES

### Step 1: Check Repository

```bash
# Navigate to GitHub repository
# Look for LICENSE file in root directory
# Read the license text
```

### Step 2: Use Go Mod Tools

```bash
# Get dependency info
go list -json github.com/package/name

# Check go.mod
cat go.mod

# Verify all modules
go mod verify
```

### Step 3: Check Online Databases

- https://opensource.org/licenses/
- https://github.com/search?q=license:MIT
- https://choosealicense.com/

### Step 4: Inspect Transitional Dependencies

```bash
# See complete dependency tree
go mod graph

# Check for GPL anywhere in the tree
go mod graph | grep -i "gpl"
```

### Step 5: When in Doubt

**DO NOT add the package.** Instead:
1. Document the uncertainty
2. Request team lead review
3. Find an alternative package
4. Wait for explicit approval

---

## 📝 DEPENDENCY TRACKING

### Required Documentation

Every significant dependency must be documented in `docs/DEPENDENCIES.md`:

```markdown
| Package                    | Version | License | Purpose | Added   | Notes                       |
|----------------------------|---------|---------|---------|---------|-----------------------------|
| github.com/sirupsen/logrus | v1.9.0  | MIT     | Logging | 2024-01 | Actively maintained         |
| github.com/spf13/viper     | v1.15.0 | MIT     | Config  | 2024-01 | Recommended by Go community |
```


---

## 🚨 AUDIT PROCEDURES

### Monthly License Audit

```bash
# Export all dependencies
go list -m all > /tmp/deps.txt

# Check for GPL
cat /tmp/deps.txt | grep -i "gpl" && echo "GPL FOUND - ALERT!" || echo "✅ No GPL"

# Verify licenses match policy
# For each package, verify LICENSE file matches approved list
```

### CI/CD License Checks

Include in your pipeline:

```yaml
# Pseudo-code for CI/CD
- name: Check Licenses
  run: |
    go mod graph | grep -E "gpl|proprietary" && exit 1 || echo "✅ Licenses OK"
    go mod verify
```

---

## 📧 ESCALATION PROCEDURE

### If You Find a License Violation

**Immediate Action Required:**

1. **STOP all work** on that dependency
2. **Document the issue:**
   - Package name and version
   - License found
   - Why it's problematic
3. **Report to team lead** immediately
4. **Find alternative** package from approved list
5. **Remove the problematic package** from go.mod
6. **Replace with approved alternative**

### If You're Unsure

**DO NOT GUESS. Instead:**

1. Research the license thoroughly
2. Check multiple sources
3. Read the actual LICENSE file
4. Ask for clarification from team lead
5. Wait for explicit approval before proceeding

---

## 🎯 BEST PRACTICES

### Do's ✅

- ✅ Always check licenses BEFORE adding a dependency
- ✅ Prefer stdlib packages when available
- ✅ Use packages from the approved list
- ✅ Document all major dependencies
- ✅ Run periodic license audits
- ✅ Keep go.mod clean with `go mod tidy`
- ✅ Review transitive dependencies

### Don'ts ❌

- ❌ Don't add packages with unclear licenses
- ❌ Don't assume a license is permissive
- ❌ Don't ignore GPL licenses "just this once"
- ❌ Don't skip license verification
- ❌ Don't add GPL-licensed packages
- ❌ Don't use proprietary packages
- ❌ Don't commit go.mod changes without license review

---

## 📚 RESOURCES

### Official License Information

- https://opensource.org/licenses/ - Official open source licenses
- https://choosealicense.com/ - License picker and guide
- https://tldrlegal.com/ - License summaries in plain English
- https://spdx.org/licenses/ - SPDX license list

### Go-Specific Resources

- https://golang.org/doc/effective_go#package_names
- https://pkg.go.dev/ - Go package documentation
- https://go.googlesource.com/go/+/master/LICENSE

### License Compliance Tools

- `go-licenses` - Tool to check Go project licenses
- GitHub's license detection - Built into GitHub
- Black Duck / WhiteSource - Commercial compliance tools

---

## ✋ FINAL REMINDER

**This is not a suggestion. This is policy.**

Every dependency added to caracal:
- ✅ MUST have MIT, Apache 2.0, BSD, ISC, or Unlicense
- ✅ MUST be from the approved packages list
- ✅ MUST be verified before adding
- ✅ MUST be documented
- ❌ MUST NOT be GPL, AGPL, LGPL, SSPL, or proprietary
- ❌ MUST NOT bypass this verification process

**Non-compliance is a blocking issue.**

---

**Version:** 1.0  
**Last Updated:** March 6, 2026  
**Status:** ACTIVE - MANDATORY POLICY  
**Enforced By:** AI Assistants and Development Team

