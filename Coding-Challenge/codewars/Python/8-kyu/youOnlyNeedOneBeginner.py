def check(seq, elem):
    try:
        if seq.index(elem) < len(seq):
            return True
    except:
        return False
# return elem in seq