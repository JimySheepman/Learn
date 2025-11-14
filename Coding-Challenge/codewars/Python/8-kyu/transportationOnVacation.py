def rental_car_cost(d):
    if d < 3:
        result = d*40
    elif d >= 3 and d < 7:
        result = d*40-20
    else:
        result = d*40-50
    return result
