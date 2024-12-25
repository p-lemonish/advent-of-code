package aoc.year24;

import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Main {
    public static void main(String[] args) throws IOException {
        String input = InputReader.readFileAsString("input.txt");
        HashMap<Long, List<Long>> equations = parseInputIntoEquations(input);

        // Generate all expressions for all entries in the equations HashMap
        long totalCalibrationResult = 0; // PART 1 answer
        for (Map.Entry<Long, List<Long>> entry : equations.entrySet()) {
            Long key = entry.getKey();
            List<Long> value = entry.getValue();
            List<String> expressions = generateExpressions(value);
            for (String expression : expressions) {
                long evaluation = evaluateExpression(expression);

                // Found a solution to the expression that matches the given key, break and go to next entry
                if(key.equals(evaluation)) {
                    totalCalibrationResult += evaluation;
                    break;
                } 
            }
        }
        System.out.println("Total calibration result: " + totalCalibrationResult);
    }

    public static long evaluateExpression(String expression) {
        String[] tokens = expression.split("(?=[+*])|(?<=[+*])"); // Crazy regex hack to get the numbers and their operators out
        long total = Long.parseLong(tokens[0]);
        for(int i = 1; i < tokens.length; i += 2) {
            String operator = tokens[i];
            long number = Long.parseLong(tokens[i+1]);

            if(operator.equals("+")) total += number;
            if(operator.equals("*")) total *= number;
        }
        return total;
    }

    // Recursively generate expressions with + and * between numbers
    public static List<String> generateExpressions(List<Long> numbers) {
        List<String> expressions = new ArrayList<>();
        generateExpressionsHelper(numbers, 1, "" + numbers.get(0), expressions);
        return expressions;
    }
    
    // Helper method for recursion
    public static void generateExpressionsHelper(List<Long> numbers, int i, String current, List<String> expressions) {
        if(i == numbers.size()) {
            expressions.add(current);
            return;
        }

        // Add +
        generateExpressionsHelper(numbers, i+1, current + "+" + numbers.get(i), expressions);

        // Add *
        generateExpressionsHelper(numbers, i+1, current + "*" + numbers.get(i), expressions);
    }
        
    // Split left and right sides of the equations longo HashMap, LHS is the total value, RHS is the list of numbers that require operators
    public static HashMap<Long, List<Long>> parseInputIntoEquations(String input) {
        String[] rows = input.split("\n");
        HashMap<Long, List<Long>> equations = new HashMap<>();
        for (String row : rows) {
            String[] leftAndRightSide = row.split(":");
            Long leftSide = Long.parseLong(leftAndRightSide[0]);
            String rightSideStr = leftAndRightSide[1].trim(); // Trim spaces from front- and backside
            leftAndRightSide[1] = rightSideStr;
            String[] rightSideStrArr = leftAndRightSide[1].split(" ");
            List<Long> rightSide = new ArrayList<>();
            for (String num : rightSideStrArr) {
                rightSide.add(Long.parseLong(num));
            }
            equations.put(leftSide, rightSide);
        }
        return equations;
    }
}