#pragma once
#include <cstddef>
//---------------------------------------------------------------------------------------
template <typename TYPE> class Point {
private:
  TYPE x_;
  TYPE y_;

public:
  Point(TYPE x = 0, TYPE y = 0) : x_(x), y_(y) {}
  //---------------------------------------------------------------------------------------
  TYPE getX() const { return x_; }
  //---------------------------------------------------------------------------------------
  TYPE getY() const { return y_; }
  //---------------------------------------------------------------------------------------
  void setX(const TYPE x) { x_ = x; }
  //---------------------------------------------------------------------------------------
  void setY(const TYPE y) { y_ = y; }
  //---------------------------------------------------------------------------------------
  bool operator==(const Point<TYPE> &other) const {
    return (x_ == other.x_ && y_ == other.y_);
  }
};

//---------------------------------------------------------------------------------------
template <typename TYPE> class PointHash {
public:
  std::size_t operator()(const Point<TYPE> &point) const {
    size_t hash=(point.getX() << 32) | point.getY();
    return hash;
  }
};
