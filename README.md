# gofibonacci
Simple go program to calculate the nth Fibonacci number using arbitrary precision integers

Uses Dijkstra's algorithm as described by http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/Fibonacci/fibFormula.html#section3

Usage: Takes a single command line parameter (the index of the number in the Fibonacci sequence to be calculated)

I created a multi-threaded version of this function just to see how efficient it is. It takes around an order of magnitude longer than the single-threaded version. 