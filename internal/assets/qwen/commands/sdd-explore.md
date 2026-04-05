---
description: Explore and investigate an idea or feature — reads codebase and compares approaches
---


CONTEXT:

- Working directory: !{pwd}
- Current project: (last component of the working directory path)
- Topic to explore: {{args}}
- Artifact store mode: engram

TASK:
Explore the topic "{{args}}" in this codebase. Investigate the current state, identify affected areas, compare approaches, and provide a recommendation.

ENGRAM PERSISTENCE (artifact store mode: engram):
Read project context (optional):
mem_search(query: "sdd-init/{project}", project: "{project}") → if found, mem_get_observation(id) for full content
Save exploration:
mem_save(title: "sdd/{{args}}/explore", topic_key: "sdd/{{args}}/explore", type: "architecture", project: "{project}", content: "{exploration}")

This is an exploration only — do NOT create any files or modify code. Just research and return your analysis.

Return a structured result with: status, executive_summary, detailed_report, artifacts, and next_recommended.
