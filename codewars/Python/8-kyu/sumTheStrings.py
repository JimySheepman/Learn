def sum_str(a, b):
    if a == "" and b != "":
        return b 
    elif b == "" and a != "":
        return a 
    elif b == "" and a == "":
        return "0"
    else:
        return str(int(a)+int(b))
        
# return str(int(a or 0) + int(b or 0))