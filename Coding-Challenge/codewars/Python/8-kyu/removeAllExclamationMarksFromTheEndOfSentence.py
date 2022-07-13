def remove(s):
    while True:
        if s[-1] != '!':
            return s
        else:
            s = s[:-1]
#     return s.rstrip("!")
