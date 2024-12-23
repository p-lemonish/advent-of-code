package aoc.year24;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.HashMap;
import java.util.List;
import java.util.stream.Collectors;

public class Main {
    public static void main(String[] args) throws IOException {
        String input = InputReader.readFileAsString("input.txt");
        String[] rulesAndUpdates = input.split(("\n\n"));
        String rules = rulesAndUpdates[0];
        String updates = rulesAndUpdates[1];
        HashMap<Integer, List<Integer>> rulesMap = new HashMap<>();

        // Put all rules into hashmap
        String[] rulesArray = rules.split("\n");
        for (String rule : rulesArray) {
            mapRules(rule, rulesMap);
        }

        // PART 1
        // Parse an update and subject it to given rules
        String[] updatesArray = updates.split("\n");
        List<String> validUpdates = new ArrayList<>();
        List<String> invalidUpdates = new ArrayList<>();
        for (String update : updatesArray) {
            if(isValid(update, rulesMap).isEmpty()) {
                validUpdates.add(update);
            } else {
                invalidUpdates.add(update);
            }
        }
        int sum = sumOfMiddlePageNumbers(validUpdates);
        System.out.println("Sum of middle page numbers: " + sum);

        // PART 2
        List<String> fixedUpdates = new ArrayList<>();
        for (String invalidUpdate : invalidUpdates) {
            List<Integer> indexes = isValid(invalidUpdate, rulesMap);
            if(!indexes.isEmpty()) {
                List<Integer> invalidUpdateIntegers = stringToIntegerList(invalidUpdate);
                List<Integer> fixedUpdateIntegers = fixInvalidUpdate(invalidUpdateIntegers, rulesMap, indexes.get(0), indexes.get(1));
                String fixedUpdate = integerListToString(fixedUpdateIntegers);
                fixedUpdates.add(fixedUpdate);
            }
        }
        int sumOfFixedMiddlePageNumbers = sumOfMiddlePageNumbers(fixedUpdates);
        System.out.println("Sum of fixed middle page numbers: " + sumOfFixedMiddlePageNumbers);

    }

    public static int sumOfMiddlePageNumbers(List<String> validUpdates) {
        int sum = 0;
        for (String update : validUpdates) {
            List<Integer> integers = stringToIntegerList(update);
            int halfway = (int) Math.floor(integers.size()/2.0);
            sum += integers.get(halfway);
        }
        return sum;
    }

    public static boolean valuesContainValue(Integer value, Collection<List<Integer>> values) {
        for (List<Integer> valueList : values) {
            if(valueList.contains(value)) return true;
        }
        return false;
    }

    public static List<Integer> stringToIntegerList(String string) {
        List<Integer> integerList = Arrays.stream(string.split(","))
                                        .map(Integer::parseInt).toList();
        return integerList;
    }

    public static String integerListToString(List<Integer> integerList) {
        return integerList.stream()
                        .map(String::valueOf)
                        .collect(Collectors.joining(",")); 
    }

    // Recursively fixes the update until isValid returns an empty list, then returns the fixedUpdate
    public static List<Integer> fixInvalidUpdate(List<Integer> updateIntegers, HashMap<Integer, List<Integer>> rulesMap, int indexOfRulesValue, int indexOfUpdateInteger) {
        List<Integer> fixedUpdate = new ArrayList<>(updateIntegers);
        fixedUpdate.set(indexOfRulesValue, updateIntegers.get(indexOfUpdateInteger));
        fixedUpdate.set(indexOfUpdateInteger, updateIntegers.get(indexOfRulesValue));
        String fixedUpdateString = integerListToString(fixedUpdate);
        List<Integer> indexes = isValid(fixedUpdateString, rulesMap);
        if(!indexes.isEmpty()) {
            fixedUpdate = fixInvalidUpdate(fixedUpdate, rulesMap, indexes.get(0), indexes.get(1));
        }
        return fixedUpdate;
    }

    // Return empty list if isValid == true, return list of two indexes which cause isValid to be false
    public static List<Integer> isValid(String update, HashMap<Integer, List<Integer>> rulesMap) {
        List<Integer> updateIntegerList = stringToIntegerList(update);
        List<Integer> invalidIndexes = new ArrayList<>();

        for (Integer updateInteger : updateIntegerList) {

            // If updateInteger is the key, it must come before the value in the update
            if(rulesMap.containsKey(updateInteger)) {
                List<Integer> ruleValues = rulesMap.get(updateInteger);

                // Go through each value and check if key (updateInteger) comes before the values in an update row
                for (Integer rulesValue : ruleValues) {

                    // If rulesValue is not in updateIntegerList => go next
                    if(!updateIntegerList.contains(rulesValue)) {
                        continue;
                    }

                    // If the updateInteger comes after the value of it's rule => not valid
                    int indexOfRulesValue = updateIntegerList.indexOf(rulesValue);
                    int indexOfUpdateInteger = updateIntegerList.indexOf(updateInteger);
                    if(indexOfRulesValue < indexOfUpdateInteger) {
                        invalidIndexes.add(indexOfRulesValue);
                        invalidIndexes.add(indexOfUpdateInteger);
                        return invalidIndexes;
                    }
                }
            } else {
                // If updateInteger is not as a key in rulesMap and also is not the last integer in the update, then return isValid false
                int indexOfUpdateInteger = updateIntegerList.indexOf(updateInteger);
                int lastIndex = updateIntegerList.size()-1;
                if(indexOfUpdateInteger != lastIndex) {
                    invalidIndexes.add(indexOfUpdateInteger);
                    invalidIndexes.add(lastIndex);
                    return invalidIndexes;
                }
            }
        }
        return invalidIndexes;
    }

    public static void mapRules(String rules, HashMap<Integer, List<Integer>> rulesMap) {
        String[] rulesSplit = rules.split("\\|");
        
        // Put value into List<Integer> if key has a List connected to it, else create a new ArrayList
        rulesMap.computeIfAbsent(Integer.parseInt(rulesSplit[0]), k -> new ArrayList<>()).add(Integer.parseInt(rulesSplit[1]));
    }
}