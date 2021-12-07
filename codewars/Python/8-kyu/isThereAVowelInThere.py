def is_vow(inp):
    return [chr(n) if chr(n) in "aeiou" else n for n in inp]
