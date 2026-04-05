---
description: Continue the next SDD phase in the dependency chain
---

Change name: {{args}}
Working directory: !{pwd}

Follow the SDD orchestrator workflow to continue the active change.

WORKFLOW:
1. Check which artifacts already exist for this change (proposal, specs, design, tasks)
2. Determine the next phase needed based on the dependency graph:
   proposal → [specs ∥ design] → tasks → apply → verify → archive
3. Let the appropriate agent(s) handle the next phase: **sdd-propose**, **sdd-spec**, **sdd-design**, **sdd-tasks**, **sdd-apply**, **sdd-verify**, or **sdd-archive**
4. Present the result and ask the user to proceed

ENGRAM NOTE:
To check which artifacts exist: mem_search(query: "sdd/{{args}}/", project: "{project}") lists all artifacts for this change.
Sub-agents handle persistence automatically with topic_key "sdd/{{args}}/{type}".

Do NOT execute phase work inline — delegate to the agents above.
