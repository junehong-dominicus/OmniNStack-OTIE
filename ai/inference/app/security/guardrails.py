import re

def scan_for_injection(prompt: str) -> bool:
    # Basic regex to simulate prompt injection detection
    patterns = [r'ignore previous instructions', r'system prompt', r'bypass']
    for pattern in patterns:
        if re.search(pattern, prompt, re.IGNORECASE):
            return True
    return False

