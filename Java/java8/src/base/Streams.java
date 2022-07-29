package base;

import java.util.Arrays;
import java.util.IntSummaryStatistics;
import java.util.List;
import java.util.Random;
import java.util.stream.Collectors;

public class Streams {
    private void genetingStreams(){
        List<String> strings = Arrays.asList("abc","","bc","efg","abcd","","jkl");
        List<String> filtered = strings
                .stream()
                .filter(string -> !string.isEmpty())
                .collect(Collectors.toList());
        System.out.println(filtered);
    }

    private void forEachs(){
        Random random = new Random();
        random.ints()
              .limit(10)
              .forEach(System.out::println);
    }


    private void maps(){
        List<Integer> numbers = Arrays.asList(3,2,2,3,7,3,5);

        List<Integer> squaresList = numbers
                .stream()
                .map(i -> i*i)
                .distinct()
                .collect(Collectors.toList());
        System.out.println(squaresList);
    }

    private void filters(){
        List<String> strings = Arrays.asList("abc","","bc","efg","abcd","","jkl");
        int count = (int) strings
                .stream()
                .filter(string -> string.isEmpty())
                .count();
    }

    private void limits(){
        Random random = new Random();
        random.ints()
                .limit(10)
                .forEach(System.out::println);
    }

    private void sorteds(){
        Random random = new Random();
        random.ints()
                .limit(10)
                .sorted()
                .forEach(System.out::println);
    }

    private void parallelProcessing(){
        List<String> strings = Arrays.asList("abc","","bc","efg","abcd","","jkl");
        long count = strings.parallelStream()
                .filter(string -> string.isEmpty())
                .count();
    }

    private void collectors(){
        List<String> strings = Arrays.asList("abc","","bc","efg","abcd","","jkl");
        List<String> filtered = strings
                .stream()
                .filter(string -> !string.isEmpty())
                .collect(Collectors.toList());

        System.out.println("Filtered List: "+filtered);
        String mergedString = strings.
                stream()
                .filter(string->!string.isEmpty())
                .collect(Collectors.joining(", "));
        System.out.println("Merged String: "+ mergedString);
    }

    private void statistics(){
        List numbers = Arrays.asList(3,2,2,3,7,3,5);

        IntSummaryStatistics stats = numbers
                .stream()
                .mapToInt((x)-> x)
                .summaryStatistics();

        System.out.println("Highest number in List : " + stats.getMax());
        System.out.println("Lowest number in List : " + stats.getMin());
        System.out.println("Sum of all numbers : " + stats.getSum());
        System.out.println("Average of all numbers : " + stats.getAverage());
    }

    public void printAll(){
        genetingStreams();
        forEachs();
        maps();
        filters();
        limits();
        sorteds();
        parallelProcessing();
        collectors();
        statistics();

        StreamExample example = new StreamExample();
        example.printAll();
    }
}
