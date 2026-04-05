---
description: Start a new SDD change — runs exploration then creates a proposal
---

Follow the SDD orchestrator workflow for starting a new change named "{{args}}".

WORKFLOW:

1. Let the **sdd-explore** agent investigate the codebase for this change
2. Present the exploration summary to the user
3. Let the **sdd-propose** agent create a proposal based on the exploration
4. Present the proposal summary and ask the user if they want to continue with specs and design

CONTEXT:

- Working directory: !`echo -n "$(pwd)"`
- Current project: !`echo -n "$(basename $(pwd))"`
- Change name: $ARGUMENTS
- Artifact store mode: engram

ENGRAM NOTE:
Sub-agents handle persistence automatically. Each phase saves its artifact to engram with topic_key "sdd/$ARGUMENTS/{type}".

Read the orchestrator instructions to coordinate this workflow. Do NOT execute phase work inline — delegate to sub-agents.
