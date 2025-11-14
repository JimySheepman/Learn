import java.util.Random;
import java.util.Scanner;

public class MineSweeper {
    Scanner scanner = new Scanner(System.in);

    public void run() {
        System.out.print("Enter the row count: ");
        int row = scanner.nextInt();

        System.out.print("Enter the column count: ");
        int column = scanner.nextInt();

        int numberOfMines = (row * column) / 4;
        createTable(row, column, numberOfMines);
    }

    public void createTable(int row, int column, int numberOfMines) {
        String[][] matrix = new String[row][column];
        int[] arr;
        for (int i = 0; i < row; i++) {
            for (int j = 0; j < column; j++) {
                matrix[i][j] = "-";
            }
        }
        for (int i = 0; i < numberOfMines; i++) {
            do {
                arr = mine(row, column);

            }
            while (matrix[arr[0]][arr[1]].equals("*"));
            matrix[arr[0]][arr[1]] = "*";
        }
        System.out.println("Locations of Mines");

        for (int i = 0; i < row; i++) {
            for (int j = 0; j < column; j++) {
                System.out.print(matrix[i][j] + " ");
            }
            System.out.print("\n");
        }
        line();
        play(matrix, row, column);
    }

    public int[] mine(int row, int column) {
        Random randomLocation = new Random();
        int[] locationsOfMines = {randomLocation.nextInt(row), randomLocation.nextInt(column)};
        return locationsOfMines;
    }

    public void line() {
        System.out.println("==========================");
    }

    public void play(String[][] matrix, int row, int column) {
        System.out.println("Welcome to the MineSweeper Game!");
        String[][] hMatrix = new String[row][column];
        for (int i = 0; i < row; i++) {
            for (int j = 0; j < column; j++) {
                hMatrix[i][j] = "-";
            }
        }

        boolean status = true;
        int count = 0;
        int rw = 0;
        int cl = 0;
        int info = 0;
        do {
            for (int i = 0; i < row; i++) {
                for (int j = 0; j < column; j++) {
                    System.out.print(hMatrix[i][j] + " ");
                }
                System.out.print("\n");
            }
            do {
                System.out.print("Enter the row count: ");
                rw = scanner.nextInt();

                System.out.print("Enter the column count: ");
                cl = scanner.nextInt();

                if (rw >= row || cl >= column) {
                    System.out.println("Do not enter a number greater than the number of rows or columns. Please try again.");
                }
            }
            while (rw >= row || cl >= column);

            if (matrix[rw][cl].equals("*")) {
                System.out.println("Game Over!");
                status = false;
            }
            else {
                info = 0;
                if (rw - 1 >= 0) {
                    if (matrix[rw - 1][cl].equals("*")) {
                        info++;
                    }
                }
                if ((rw - 1 >= 0) && (cl - 1 >= 0)) {
                    if (matrix[rw - 1][cl - 1].equals("*")) {
                        info++;
                    }
                }
                if (cl - 1 >= 0) {
                    if (matrix[rw][cl - 1].equals("*")) {
                        info++;
                    }
                }
                if (cl + 1 < column) {
                    if (matrix[rw][cl + 1].equals("*")) {
                        info++;
                    }
                }
                if ((cl + 1 < column) && (rw + 1 < row)) {
                    if (matrix[rw + 1][cl + 1].equals("*")) {
                        info++;
                    }
                }
                if (rw + 1 < row) {
                    if (matrix[rw + 1][cl].equals("*")) {
                        info++;
                    }
                }
                if ((rw - 1 >= 0) && (cl + 1 < column)) {
                    if (matrix[rw - 1][cl + 1].equals("*")) {
                        info++;
                    }
                }
                if ((rw + 1 < row) && (cl - 1 >= 0)) {
                    if (matrix[rw + 1][cl - 1].equals("*")) {
                        info++;
                    }
                }
                hMatrix[rw][cl] = String.valueOf(info);
                count++;
            }
            line();
        }
        while (status && count < (row * column) - ((row * column) / 4));

        if (status) {
            System.out.println("Congratulations! You won the Game!");
            for (int i = 0; i < row; i++) {
                for (int j = 0; j < column; j++) {
                    if (hMatrix[i][j].equals("-")) {
                        System.out.print("* ");
                    }
                    else {
                        System.out.print(hMatrix[i][j] + " ");
                    }
                }
                System.out.print("\n");
            }
        }
    }
}