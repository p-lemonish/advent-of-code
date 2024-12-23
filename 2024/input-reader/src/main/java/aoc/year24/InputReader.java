package aoc.year24;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;

public class InputReader {
    public static String readFileAsString(String resourceName) throws IOException {
        InputStream inputStream = InputReader.class.getClassLoader().getResourceAsStream(resourceName);
        if (inputStream == null) {
            throw new IOException("Resource not found: " + resourceName);
        }
        StringBuilder result = new StringBuilder();
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream))) {
            String line;
            while ((line = reader.readLine()) != null) {
                result.append(line).append(System.lineSeparator());
            }
        }
        return result.toString().trim();
    }
}
