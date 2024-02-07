# gittufscanner (POC)

## Quick Overview

This Proof of Concept (POC) project safeguards the integrity of git tags in the top 1000 critical repositories, ensuring they remain unchanged and secure. It's a proactive shield against unauthorized tag alterations that could disrupt many systems.

## Why It Matters

Critical repositories are the backbone of numerous systems. Unauthorized changes to their git tags can cause significant issues. This POC prevents such risks by monitoring and locking down tag changes.

## How It Works

1. **Clone Repositories:** We clone the critical repositories to monitor them locally.
2. **Integrate Gittuf:** We use the gittuf tool to monitor tag changes.
3. **Lock Tags:** Initially, we trust new tags. Then, we lock each tag to prevent unauthorized changes, recording their state securely.
4. **Monitor Changes:** Regular checks are made for unauthorized tag changes. If found, they're flagged as security risks.

## Key Points

- **Security:** Quickly identifies and flags unauthorized tag changes.
- **Assumptions:** Initially trusts all projects until they're monitored.
- **Challenges:** Uses shell scripts for operations due to the lack of an API in gittuf and faces limitations in detecting newly created and altered tags.

## Bottom Line

This POC is a critical step towards ensuring the security and integrity of important repositories by monitoring and protecting git tags from unauthorized changes.