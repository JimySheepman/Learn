def calculator(x, y, op):
    if type(x) == type(1) and type(y) == type(1):
        if op == '+':
            return x+y
        elif op == '-':
            return x-y
        elif op == '*':
            return x*y
        elif op == '/':
            return x/y
        else:
            return "unknown value"
    else:
        return "unknown value"
# return eval(f'{x}{op}{y}') if type(x) == type(y) == int and str(op) in '+-*/' else 'unknown value'