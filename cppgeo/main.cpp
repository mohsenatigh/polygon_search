#include "GeoMatrix.hpp"
#include "Polygon.hpp"
#include "Rectangle.hpp"
#include <chrono>
#include <fstream>
#include <functional>
#include <iostream>
#include <sstream>
#include <string>
#include <thread>
#include <unistd.h>

using Matrix = GeoMatrix<double, uint64_t, 1000>;
//---------------------------------------------------------------------------------------
bool loadFile(const std::string &fileName, Matrix &matrix) {

  std::ifstream file(fileName);
  if (!file.is_open()) {
    return false;
  }

  auto readPoint = [](const std::string &in) {
    std::stringstream str(in);
    std::string x, y;
    getline(str, x, ',');
    getline(str, y, ',');
    return Point<double>(std::stod(x), std::stod(y));
  };

  std::string line;
  std::string part;

  while (getline(file, line)) {
    std::stringstream str(line);
    std::string id;
    std::vector<Point<double>> pointsList;

    if (!getline(str, id, '|')) {
      continue;
    }

    for (int i = 0; getline(str, part, '|'); i++) {
      auto point = readPoint(part);
      pointsList.push_back(point);
    }

    matrix.add(Polygon<double>(pointsList), std::stol(id));
  }

  file.close();
  return true;
}

//---------------------------------------------------------------------------------------
void RunFunc(std::function<void()> func, const std::string &title) {

  std::cout << "Run " << title << " in ";

  auto start = std::chrono::high_resolution_clock::now();
  func();
  auto end = std::chrono::high_resolution_clock::now();
  std::cout << std::chrono::duration_cast<std::chrono::microseconds>(end -
                                                                     start)
                   .count()
            << " Micro Second \n"
            << std::endl;
}
//---------------------------------------------------------------------------------------

int main(int argc, char *argv[]) {

  Matrix rMat;

  int c;
  const std::string help = "-f data file\n"
                           "-a lat\n"
                           "-o long\n"
                           "-H search with high accuracy\n"
                           "-L search with low accuracy\n"
                           "-h show help\n";

  std::string dataFile = "data.csv";
  double lat = 7.847786843776704;
  double lng = 47.99549694289439;
  uint32_t cCount = 1000;
  Matrix::GeoMatrixAccuracy acc=Matrix::MEDIUM;
  std::vector<uint64_t> out;

  while ((c = getopt(argc, argv, "c:f:a:o:hHL")) != -1) {
    switch (c) {
    case 'f':
      dataFile = std::string(optarg);
      break;
    case 'a':
      lat = std::stod(optarg);
      break;
    case 'o':
      lng = std::stod(optarg);
      break;
    case 'c':
      cCount = std::stoi(optarg);
      break;
    case 'H':
      acc=Matrix::HIGH;
      break;
    case 'L':
      acc=Matrix::LOW;
      break;
    case 'h':
      std::cout << help << "\n";
      exit(0);
      break;
    }
  }

  
  out.reserve(32000);
  Point<double> queryPoint(lat, lng);

  RunFunc(
      [&]() {
        if (!loadFile(dataFile, rMat)) {
          std::cout << "load DB failed  \n";
          exit(0);
        };
      },
      "Load DB");

  RunFunc([&]() { rMat.build(); }, "Build DB");

  out = rMat.query(queryPoint,acc,-1);
  for (auto &i : out) {
    std::cout << "Find Polygon with ID : " << i << std::endl;
  }

  uint64_t count = 0;
  RunFunc(
      [&]() {
        for (uint32_t i = 0; i < cCount; i++) {
          out.clear();
          rMat.query(queryPoint, out,acc,1);
          count += out.size();
        }
      },
      "Performance Test");

  std::cout << "select " << count << " polygon \n";
}