package aoc.year24;

import java.io.IOException;

public class Main {
    public static void main(String[] args) throws IOException {
        String input = InputReader.readFileAsString("inputtest.txt");
        final char[][] baseField = convertFieldToCharArray(input);
        String[] directions = {"^",">","v","<"};

        // PART 1
        String guardMap = input;
        char[][] guardField = convertFieldToCharArray(guardMap);
        Guard g = determineGuard(guardField, directions);
        guardField = runGuardThroughMap(g, guardField, directions);
        System.out.println(guardMap);
        System.out.println("Amount of X: " + amountOfX(guardField));

        // PART 2
        System.out.println("Finding ways for guard to get stuck in loops...");
        char[][] fieldToFillWithObstacles = deepCopy(baseField);
        int stuckCounter = 0;
        for(int i = 0; i < baseField.length; i++) {
            for(int j = 0; j < baseField[i].length; j++) {

                fieldToFillWithObstacles = deepCopy(baseField);

                // Set a # to the next spot where the guard has visited
                char charInFieldSlot = fieldToFillWithObstacles[i][j];
                if(guardField[i][j] == 'X') {
                    fieldToFillWithObstacles[i][j] = '#';
                } else {
                    continue;
                }

                // Generate a guard with the given map
                Guard g2 = determineGuard(fieldToFillWithObstacles, directions);

                // Run the guard through the map, if it comes back with isReachedEnd == false, it got stuck
                fieldToFillWithObstacles = runGuardThroughMap(g2, fieldToFillWithObstacles, directions);
                if(!g2.isReachedEnd()) {
                    System.out.println("Guard got stuck while setting obstacle at x: " + j + " y: " + i);
                    stuckCounter++;
                }

                // Put the char that was in the field back
                fieldToFillWithObstacles[i][j] = charInFieldSlot;
            }

        }
        System.out.println("Amount of times guard got stuck: " + stuckCounter);

    }

    public static char[][] deepCopy(char[][] original) {
        if (original == null) {
            return null;
        }
        char[][] copy = new char[original.length][];
        for (int i = 0; i < original.length; i++) {
            copy[i] = original[i].clone();
        }
        return copy;
    }

    public static char[][] convertFieldToCharArray(String field) {
        String[] rows = field.split("\n");
        char[][] grid = new char[rows.length][];
        for(int i = 0; i < rows.length; i++) {
            grid[i] = rows[i].toCharArray();
        }
        return grid;
    }

    public static void setObstacleOnMap(char[][] grid, int x, int y) {
        if(grid[y][x] == '.') {
            grid[y][x] = '#';
        }
    }

    public static char[][] runGuardThroughMap(Guard g, char[][] guardField, String[] directions) {
            guardField = setGuardInField(guardField, g);
            int currentAmountOfX = 0;
            int stuckCheck = guardField.length;
            int i = 0;
            while(!g.isReachedEnd()) {
                int newAmountOfX = amountOfX(guardField);
                g = moveOrChangeDirection(g, guardField, directions);
                guardField = setGuardInField(guardField, g);
    
                // Check if amount of X has changed, which would mean the guard has explored new areas
                if(currentAmountOfX == newAmountOfX) {
                    i++;
                } else {
                    // If guard has checked new area, update amount of area checked and reset stuck-counter
                    i = 0;
                    currentAmountOfX = newAmountOfX;
                }
    
                if(i >= stuckCheck) {
                    break;
                }
            }
            return guardField;
    }

    public static String gridToString(char[][] grid) {
        StringBuilder sb = new StringBuilder();
        for(int i = 0; i < grid.length; i++) {
            sb.append(grid[i]);
            if(i < grid.length - 1) {
                sb.append("\n");
            }
        }
        return sb.toString();
    }

    public static int amountOfX(char[][] field) {
        int amount = 0;
        for(int i = 0; i < field.length; i++) {
            for(int j = 0; j < field[i].length; j++) {
                if(field[i][j] == 'X') amount++;
            }
        }
        return amount;
    }

    public static Guard moveOrChangeDirection(Guard g, char[][] field, String[] directions) {
        for(int i = 0; i < directions.length; i++) {
            if(g.getX() == 0 || g.getY() == 0 || g.getX() == field[0].length-1 || g.getY() == field.length-1) {
                g.setReachedEnd(true);
                break;
            }
            if(g.getDirection().equals(directions[i])) {
                // Move up
                if(i == 0) {
                    // Check if (x,y-1) has # in it
                    char[] rowToCheck = field[g.getY()-1];
                    char slotToCheck = rowToCheck[g.getX()];
                    if(slotToCheck == '#') {
                        g.setDirection(directions[i+1]);
                    } else {
                        g.setY(g.getY()-1);
                    }
                }
                // Move right
                if(i == 1) {
                    // Check if (x+1,y) has # in it
                    char[] rowToCheck = field[g.getY()];
                    char slotToCheck = rowToCheck[g.getX()+1];
                    if(slotToCheck == '#') {
                        g.setDirection(directions[i+1]);
                    } else {
                        g.setX(g.getX()+1);
                    }
                }
                // Move down
                if(i == 2) {
                    // Check if (x,y+1) has # in it
                    char[] rowToCheck = field[g.getY()+1];
                    char slotToCheck = rowToCheck[g.getX()];
                    if(slotToCheck == '#') {
                        g.setDirection(directions[i+1]);
                    } else {
                        g.setY(g.getY()+1);
                    }
                }
                // Move left
                if(i == 3) {
                    // Check if (x-1,y) has # in it
                    char[] rowToCheck = field[g.getY()];
                    char slotToCheck = rowToCheck[g.getX()-1];
                    if(slotToCheck == '#') {
                        g.setDirection(directions[0]);
                    } else {
                        g.setX(g.getX()-1);
                    }
                }
            }
        }
        return g;
    }

    public static char[][] setGuardInField(char[][] field, Guard g) {
            field[g.getY()][g.getX()] = 'X';
            return field;
    }

    public static Guard determineGuard(char[][] field, String[] directions) {
        Guard guard = new Guard();
        for (int i = 0; i < field.length; i++) {
            for (String direction : directions) {
                for(int j = 0; j < field[i].length; j++) {
                    if(field[i][j] == direction.toCharArray()[0]) {
                        guard.setDirection(direction);
                        guard.setX(j);
                        guard.setY(i);
                    }
                }
            }
        }
        return guard;
    }
}