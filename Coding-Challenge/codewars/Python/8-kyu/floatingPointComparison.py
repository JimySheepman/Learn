def approx_equals(a, b):
    if a-b >= 0.001 or a-b <= -0.001:
        return False
    elif b-a >= 0.001 or b-a <= -0.001:
        return False
    else:
        return True
