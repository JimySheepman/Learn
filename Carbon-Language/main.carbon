package Geometry api;
import Math;

class Circle{
    var r: f32;
};

fn PrintTotalArea(circles: Slice(Circle)){
    var area: f32 = 0;
    for(c: Circle in circles){
        area += Math.Pi * c.r *c.r;
    }
    Print("Total area: {0}",area)
}

fn Main() -> i32{
    var circles: Array(Circle) = ({.r = 1.0},{.r = 2.0});

    PrintTotalArea(circles);
    return 0;
}