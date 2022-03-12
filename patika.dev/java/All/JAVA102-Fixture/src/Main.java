import java.util.ArrayList;
import java.util.Random;

public class Main {

    public static void main(String[] args) {
        Random rand = new Random();
        ArrayList<Teams> teams = new ArrayList<Teams>();
        ArrayList<Teams> tempTeams = new ArrayList<Teams>();
        ArrayList<Weeks> weeks = new ArrayList<Weeks>();
        Teams homeTeam;
        Teams awayTeam;

        teams.add(new Teams("Real Madrid"));
        teams.add(new Teams("FC Barcelona"));
        teams.add(new Teams("Atlético de Madrid"));
        teams.add(new Teams("Sevilla"));
        teams.add(new Teams("Real Sociedad"));
        teams.add(new Teams("Villarreal"));

        if (teams.size() % 2 == 1) {
            teams.add(new Teams("BAY"));
        }

        int numberOfGames = teams.size() - 1;
        int week = 1;
        while (numberOfGames > 0) {
            tempTeams.addAll(teams);
            Weeks wk = new Weeks(week);
            while (tempTeams.size() > 0) {
                int random = rand.nextInt(tempTeams.size());
                homeTeam = tempTeams.get(random);
                tempTeams.remove(homeTeam);
                random = rand.nextInt(tempTeams.size());
                awayTeam = tempTeams.get(random);
                if (!homeTeam.opponent.contains(awayTeam.name)) {
                    homeTeam.opponent.add(awayTeam.name);
                    tempTeams.remove(awayTeam);
                    wk.homeTeam.add(homeTeam.name);
                    wk.awayTeam.add(awayTeam.name);
                }
                else {
                    wk = new Weeks(week);
                    tempTeams.clear();
                    tempTeams.addAll(teams);
                }
            }
            numberOfGames--;
            week++;
            weeks.add(wk);
        }

        for(Weeks wk : weeks){
            System.out.println("(1st Half of the Season) WEEK: " + wk.number);
            for (int i = 0; i < wk.awayTeam.size(); i++) {
                System.out.println(wk.homeTeam.get(i) + " vs " + wk.awayTeam.get(i));
            }
            System.out.println();
        }

        System.out.println("---------------------------------------------------------------\n");

        for(Weeks wk : weeks){
            System.out.println("(2nd Half of the Season) WEEK: " + (wk.number + numberOfGames));
            for (int i = 0; i < wk.awayTeam.size(); i++) {
                System.out.println(wk.awayTeam.get(i) + " vs " + wk.homeTeam.get(i));
            }
            System.out.println();
        }
    }
}