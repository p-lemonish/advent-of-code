package aoc.year24;

public class Guard {
    private int x;
    private int y;
    private String direction;
    public boolean isReachedEnd() {
        return reachedEnd;
    }
    public void setReachedEnd(boolean reachedEnd) {
        this.reachedEnd = reachedEnd;
    }
    private boolean reachedEnd;
    public Guard() {}
    public int getX() {
        return x;
    }
    public void setX(int x) {
        this.x = x;
    }
    @Override
    public String toString() {
        return "Guard [x=" + x + ", y=" + y + ", direction=" + direction + ", reachedEnd=" + reachedEnd + "]";
    }
    public int getY() {
        return y;
    }
    public void setY(int y) {
        this.y = y;
    }
    public String getDirection() {
        return direction;
    }
    public void setDirection(String direction) {
        this.direction = direction;
    }
}
