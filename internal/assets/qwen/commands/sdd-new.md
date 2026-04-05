---
description: Start a new SDD change — runs exploration then creates a proposal
---

Change name: {{args}}
Working directory: !{pwd}

Follow the SDD orchestrator workflow for starting a new change named "{{args}}".

WORKFLOW:

1. Let the **sdd-explore** agent investigate the codebase for this change
2. Present the exploration summary to the user
3. Let the **sdd-propose** agent create a proposal based on the exploration
4. Present the proposal summary and ask the user if they want to continue with specs and design

ENGRAM NOTE:
Sub-agents handle persistence automatically. Each phase saves its artifact to engram with topic_key "sdd/{{args}}/{type}".

Do NOT execute phase work inline — delegate to the agents above.
