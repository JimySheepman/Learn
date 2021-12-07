def fillable(stock, merch, n):
    if n >= 1:
        if merch in stock:
            if n <= stock[merch]:
                return True
            else:
                return False
        else:
            return False
    else:
        return 0
