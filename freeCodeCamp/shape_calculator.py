class Rectangle:
    def __init__(self, width, height):
        self.width = width
        self.height = height

    def __repr__(self):
        return 'Rectangle(width=' + str(self.width) + ', height=' + str(self.height) + ')'

    def set_width(self, width):
        self.width = width

    def set_height(self, height):
        self.height = height

    def get_area(self):
        return self.width * self.height

    def get_perimeter(self):
        return 2 * self.width + 2 * self.height

    def get_diagonal(self):
        return (self.width ** 2 + self.height ** 2) ** 0.5

    def get_picture(self):
        if self.width > 50 or self.height > 50: return 'Too big for picture.'
        pic = ''
        for x in range(0, self.height):
            pic += '*'*self.width + '\n'

        return pic

    def get_amount_inside(self, shape):
        if shape.width > self.width or shape.height > shape.height: return 0
        
        return (self.width // shape.width) * (self.height // shape.height)


class Square(Rectangle):
    def __init__(self, side):
        self.width = side
        self.height = side

    def __repr__(self):
        return 'Square(side=' + str(self.width) + ')'

    def set_side(self, side):
        self.width = side
        self.height = side