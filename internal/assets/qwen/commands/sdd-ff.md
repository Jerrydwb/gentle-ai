---
description: Fast-forward all SDD planning phases — proposal through tasks
---

Follow the SDD orchestrator workflow to fast-forward all planning phases for change "{{args}}".

WORKFLOW:
Run these sub-agents in sequence:

1. sdd-propose — create the proposal
2. sdd-spec — write specifications
3. sdd-design — create technical design
4. sdd-tasks — break down into implementation tasks

Present a combined summary after ALL phases complete (not between each one).

CONTEXT:

- Working directory: !{pwd}
- Current project: (last component of the working directory path)
- Change name: {{args}}
- Artifact store mode: engram

ENGRAM NOTE:
Sub-agents handle persistence automatically. Each phase saves its artifact to engram with topic_key "sdd/{{args}}/{type}" where type is: proposal, spec, design, tasks.

Read the orchestrator instructions to coordinate this workflow. Do NOT execute phase work inline — delegate to sub-agents.
