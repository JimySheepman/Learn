package first_ten_classes;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class RegexMatches {


    public void printAll(){
        example01();
        example02();
        example03();
        example04();
        example05();
    }

    public void example01(){
        String line = "This order was placed for QT3000! ok?";
        String pattern = "(.*)(\\d+)(.*)";

        Pattern r = Pattern.compile(pattern);
        Matcher m =r.matcher(line);

        if (m.find( )){
            System.out.println("found value: "+m.group(0));
            System.out.println("found value: "+m.group(1));
            System.out.println("found value: "+m.group(2));
        }else {
            System.out.println("no match");
        }
    }
    public void example02(){
        String input = "cat cat cat cattie cat";
        String regex = "\\bcat\\b";
        Pattern p = Pattern.compile(regex);
        Matcher m = p.matcher(input);
        int count = 0;

        while (m.find()){
            count++;
            System.out.println("match number: "+count);
            System.out.println("start(): "+m.start());
            System.out.println("end(): "+m.end());
        }

    }
    public void example03(){
        String regex = "foo";
        String input = "fooooooooooooooooo";
        Pattern pattern;
        Matcher matcher;

        pattern = Pattern.compile(regex);
        matcher = pattern.matcher(input);

        System.out.println("Current REGEX is: "+regex);
        System.out.println("Current INPUT is: "+input);

        System.out.println("lookingAt(): "+matcher.lookingAt());
        System.out.println("matches(): "+matcher.matches());
    }
    public void example04(){
        String regex = "dog";
        String input = "The dog says meow. "+ "All dog say meow.";
        String replace = "cat";

        Pattern p = Pattern.compile(regex);
        Matcher m = p.matcher(input);
        input = m.replaceAll(replace);
        System.out.println(input);
    }

    public void example05(){
        String regex = "a*b";
        String input = "aabfooaabfooabfoob";
        String replace = "-";
        Pattern p = Pattern.compile(regex);
        Matcher m = p.matcher(input);
        StringBuffer sb = new StringBuffer();
        while (m.find()){
            m.appendReplacement(sb,replace);
        }
        m.appendTail(sb);
        System.out.println(sb.toString());
    }
}
