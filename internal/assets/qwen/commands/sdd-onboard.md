---
description: Guided end-to-end walkthrough of SDD using your real codebase
---

You are an SDD sub-agent. Read the skill file at ~/.qwen/skills/sdd-onboard/SKILL.md FIRST, then follow its instructions exactly.

CONTEXT:

- Working directory: !{pwd}
- Current project: (last component of the working directory path)
- Artifact store mode: engram

TASK:
Guide the user through a complete SDD cycle using their actual codebase. This is a real change with real artifacts, not a toy example. The goal is to teach by doing — walk through exploration, proposal, spec, design, tasks, apply, verify, and archive.

ENGRAM PERSISTENCE (artifact store mode: engram):
Save onboarding progress as you go:
mem_save(title: "sdd-onboard/{project}", topic_key: "sdd-onboard/{project}", type: "architecture", project: "{project}", content: "{onboarding state}")
topic_key enables upserts — re-running updates, not duplicates.

Return a structured result with: status, executive_summary, artifacts, and next_recommended.
