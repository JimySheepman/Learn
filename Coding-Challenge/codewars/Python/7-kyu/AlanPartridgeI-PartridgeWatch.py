def part(arr):
    flag = False
    count = 0
    key_list = ["Partridge", "PearTree", "Chat", "Dan",
                "Toblerone", "Lynn", "AlphaPapa", "Nomad"]
    for i in key_list:
        for j in arr:
            if i == j:
                flag = True
                count += 1
    if flag:
        return "Mine's a Pint"+''.join(("!" for c in range(count)))
    else:
        return "Lynn, I've pierced my foot on a spike!!"


# l = ["Partridge", "PearTree", "Chat", "Dan", "Toblerone", "Lynn", "AlphaPapa", "Nomad"]
# s = len([i for i in arr if i in l])
# return "Mine's a Pint"+"!"*s if s>0 else 'Lynn, I\'ve pierced my foot on a spike!!'
