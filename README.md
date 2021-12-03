# Polygon Search
This simple program uses Geospatial Partitioning to find the postal code of a specific geographical point in Germany.  I use this program to compare the performance of four popular programming languages. There are some points regarding this benchmark:

- 	I didn't use any third-party library for this benchmark, and all the implementations remain based on standard data structures and libraries available in each programming language. 
-   The benchmark result is based on sample data containing polygons that cover all the german region.
-   To keep the system as simple as possible, I am using a custom file format for importing the polygons data. you can download the original data from https://www.suche-postleitzahl.org/plz-karte-erstellen  
-   the algorithm contains three operation modes:
	-	**HIGH**
	-	**MEDIUM**
	-	**LOW**  

## Compiler version
-   The results are related to the following compiler versions : 
    -   **gcc 9.3.0**
    -   **clang 10.0.0**
    -   **jdk-17**
    -   **go version 1.17.1**
    -   **rustc 1.55.0**

## Test Data


## Build
### C++ gcc
    cd cppgeo
    mkdir __buid
    cmake ..
    make -j
### C++ clang
    uncomment set(CMAKE_C_COMPILER "clang") set(CMAKE_CXX_COMPILER "clang++") in the CMakeLists.txt file
    cd cppgeo
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

### RUST
    cd rust/polygon_search
    cargo build --release

## Usage 
For viewing the available options for each version, use -h switch

## Test Scenario
The results are based on 1000000 times lookup of the specified point and are in Microsecond.

### CASE 1
The first Scenario considers the location of the Freiburg central church.

**Freiburger Münster Lookup results (High and low accuracy)**

!["High accuracy lookup result"](https://github.com/mohsenatigh/polygon_search/blob/main/images/f1.png)
!["Low accuracy lookup result"](https://github.com/mohsenatigh/polygon_search/blob/main/images/f2.png)

| Language | High Accuracy  | Medium Accuracy | Low Accuracy |
| ------------ | ------------ | ------------ | ------------ |
| C++ (gcc)  |  1394029 μs | 354843 μs | 271316 μs |
| C++ (clang) |  1654293 μs | 384636 μs | 320392 μs |
|  JAVA |  1969147 μs |  595850 μs |  514363 μs |
|  GO | 2130709 μs | 719377 μs | 652770 μs |
|  RUST | 2200332 μs | 504338 μs | 433640 μs |

!["lookup results"](https://github.com/mohsenatigh/polygon_search/blob/main/images/1.png)

**Random points around Freiburger Münster**

| Language | High Accuracy  | Medium Accuracy | Low Accuracy |
| ------------ | ------------ | ------------ | ------------ |
| C++ (gcc)  |  1501170 μs | 355831 μs | 273120 μs |
| C++ (clang) |  1724351 μs | 387201 μs | 323969 μs |
|  JAVA |  2377246 μs |  630499 μs |  539974 μs |
|  GO | 2216277 μs | 725739 μs | 668134 μs |
|  RUST | 2283935 μs | 508129 μs | 443500 μs |

!["lookup results"](https://github.com/mohsenatigh/polygon_search/blob/main/images/2.png)

**A location in a big polygon (High and low accuracy)**
!["High accuracy lookup result"](https://github.com/mohsenatigh/polygon_search/blob/main/images/p2.png)
!["Low accuracy lookup result"](https://github.com/mohsenatigh/polygon_search/blob/main/images/p1.png)

| Language | High Accuracy  | Medium Accuracy | Low Accuracy |
| ------------ | ------------ | ------------ | ------------ |
| C++ (gcc)  |  5187784 μs | 378546 μs | 321742 μs |
| C++ (clang) |  5187583 μs | 380561 μs | 335955 μs |
|  JAVA |  7507826 μs |  566826 μs |  504545 μs |
|  GO | 6492508 μs | 708009 μs | 643355 μs |
|  RUST | 9399210 μs | 508060 μs | 430104 μs |

!["lookup results"](https://github.com/mohsenatigh/polygon_search/blob/main/images/3.png)

**A location in a big polygon ( Random Points )**

| Language | High Accuracy  | Medium Accuracy | Low Accuracy |
| ------------ | ------------ | ------------ | ------------ |
| C++ (gcc)  |  5246180 μs | 386938 μs | 326808 μs |
| C++ (clang) |  5261288 μs | 388207 μs | 336071 μs |
|  JAVA |  7867935 μs |  604901 μs |  574479 μs |
|  GO | 6531627 μs | 712063 μs | 650678 μs |
|  RUST | 9396629 μs | 506499 μs | 444764 μs |

!["lookup results"](https://github.com/mohsenatigh/polygon_search/blob/main/images/4.png)
