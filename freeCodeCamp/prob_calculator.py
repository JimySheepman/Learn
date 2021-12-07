import copy
import random

class Hat:
    def __init__(self, **kwargs):
        self.contents = []
        for k, v in kwargs.items():
            for i in range(v):
                self.contents.append(k)
        self.contents_backup = copy.deepcopy(self.contents)
    
    def draw(self, num):
        self.drawn = []
        for j in range(num):
            if self.contents == []:
                self.contents = copy.deepcopy(self.contents_backup)
            col = random.choice(self.contents)
            self.contents.remove(col)
            self.drawn.append(col)
        return self.drawn

def success(hat, expected_balls, num_balls_drawn):
    hat.contents = copy.deepcopy(hat.contents_backup)
    draw_result = hat.draw(num_balls_drawn)
    for colour, number in expected_balls.items():
        for i in range(number):
            try:
                draw_result.remove(colour)
            except:
                return 0
    return 1

def experiment(hat, expected_balls, num_balls_drawn, num_experiments):
    succ = 0
    for i in range(num_experiments):
        succ += success(hat, expected_balls, num_balls_drawn)
    prob = succ / num_experiments
    return prob