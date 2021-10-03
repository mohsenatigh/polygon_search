# Polygon Search
This program uses an equal space portioning technique for solving the problem of locating polygons in a huge polygon list. As it is a relatively simple algorithm, it was a suitable candidate to compare and benchmark the functionality of different programming languages.  There are some points regarding this benchmark:

- 	I didn't use any third-party library for this benchmark, and all the implementations remain based on standard data structures and libraries available in each programming language. 
-   The benchmark result is based on sample data containing 86000 polygons that cover all the german regions.
-   To keep the system as simple as possible, I am using a custom file format for importing the polygons data. 
-   the algorithm contains three operation modes:
	-	**HIGH**: which returns the exact polygon containing the query point
	-	**MEDIUM**:  which returns the polygon(s) with relatively high accuracy with few hundred-meter fault, which is mainly based on the region and data
	-	**LOW**:   for the German region, the fault rate is about 1.5 to 2 kilometers
-   The memory usage will be reduced (by the magnitude of 10 or 15 times ) if we disable the high accuracy mode.
-   The complexity is O(1) for 46% of the regions and O(M) for others  (The maximum M for the current sample data is 13). 
-   I considered the Java JIT  functionality by adding a warmup phase before starting the actual test
-   The medium accuracy version is usually enough for most of the usage,

##Build
### C++
    cd cpp
    mkdir __buid
    cmake ..
    make -j 

### JAVA 
	cd java
	export PATH=$PATH:{JDKPATH}
    export PATH=$PATH:{GRADLEPATH}/bin
    gradle build

### GO
	cd go
    go build

## Usage 
For viewing the available options for each version, use -h switch

## Test Scenario
The test scenario contains four different test cases that cover most of the aspects of the algorithms. The results are based on 1000 times lookup of the specified point and are in Microsecond.

### CASE 1
The first scenario considers a location in the middle of a very sparse area. the complexity of the search is O(1) for all the three search mode

!["lookup location"](https://github.com/mohsenatigh/polygon_search/blob/main/images/1.png)

the following table is the lookup results

| Language | High Accuracy  | Medium Accuracy | Low Accuracy |
| ------------ | ------------ | ------------ | ------------ |
| C++  |  17 μs | 17 μs | 17 μs |
|  JAVA |  95 μs |  95 μs |  95 μs |
|  GO | 70 μs | 70 μs | 70 μs |

!["lookup results"](https://github.com/mohsenatigh/polygon_search/blob/main/images/1_chart.png)

### CASE 2

The second case considers the "Freiburger Münster." it is a location with an average complexity.The following image visualizes the search results in the three levels of accuracy

!["lookup location"](https://github.com/mohsenatigh/polygon_search/blob/main/images/2.gif)

the following table is the lookup results

| Language | High Accuracy  | Medium Accuracy | Low Accuracy |
| ------------ | ------------ | ------------ | ------------ |
| C++  |  62 μs | 20 μs | 17 μs |
|  JAVA |  111 μs |  102 μs |  95 μs |
|  GO | 127 μs | 79 μs | 70 μs |

!["lookup results"](https://github.com/mohsenatigh/polygon_search/blob/main/images/2_chart.png)

## Case 3

Case 3 is the lookup of a location in the densest area of the map (center of Nuremberg)

!["lookup location"](https://github.com/mohsenatigh/polygon_search/blob/main/images/3.gif)

the following table is the lookup results

| Language | High Accuracy  | Medium Accuracy | Low Accuracy |
| ------------ | ------------ | ------------ | ------------ |
| C++  |  76 μs | 45 μs | 17 μs |
|  JAVA |  113 μs |  105 μs |  95 μs |
|  GO | 200 μs | 163 μs | 70 μs |

!["lookup results"](https://github.com/mohsenatigh/polygon_search/blob/main/images/3_chart.png)
