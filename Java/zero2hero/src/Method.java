public class Method {
    int key;

    Method(){}

    Method(int key){
        this.key = key;
    }
    public  int minFuction(int n1,int n2 ){
        int min;
        if (n1>n2){
            min = n2;
        }else {
            min = n1;
        }
        return min;
    }

    public void methodRankPoints(double points){
        if (points>=202.5){
            System.out.println("Rank:A1");
        } else if (points>=122.4) {
            System.out.println("Rank:A2");
        }else {
            System.out.println("Rank:A1");
        }
    }

    public void swapFunction(int a,int b){
        int c;
        c = a;
        a=b;
        b=c;
        System.out.println("After swapping(Inside), a = "+a+" b= "+b);
    }

    // method overloding
    public double minFunction(double n1,double n2){
        double min;
        if (n1>n2){
            min = n2;
        }else {
            min = n1;
        }
        return min;
    }
}
