#pragma once
#include "Point.hpp"
#include "Rectangle.hpp"
#include <iostream>
#include <vector>

template <typename TYPE> class Polygon {
private:
  using Points = std::vector<Point<TYPE>>;

  Points points_;

  mutable Rectangle<TYPE> bBox_;
  mutable bool havebBox_ = false;

public:
  Polygon(const Points &points) : points_(points) {}
  Polygon() {}
  //---------------------------------------------------------------------------------------
  const Rectangle<TYPE> &getBoundingBox() const {

    if (havebBox_) {
      return bBox_;
    }

    Point<TYPE> start = {std::numeric_limits<TYPE>::max(),
                         std::numeric_limits<TYPE>::max()};

    Point<TYPE> end = {std::numeric_limits<TYPE>::min(),
                       std::numeric_limits<TYPE>::min()};

    for (const auto &i : points_) {
      if (i.getX() < start.getX()) {
        start.setX(i.getX());
      }

      if (i.getY() < start.getY()) {
        start.setY(i.getY());
      }

      if (i.getX() > end.getX()) {
        end.setX(i.getX());
      }

      if (i.getY() > end.getY()) {
        end.setY(i.getY());
      }
    }

    bBox_.setStart(start);
    bBox_.setEnd(end);
    havebBox_ = true;
    return bBox_;
  }
  //---------------------------------------------------------------------------------------
  bool pointIsInside(const Point<TYPE> &point) const {
    int count = 0;

    if (havebBox_ && !bBox_.pointInside(point)) {
      return false;
    }

    size_t size = points_.size();
    for (size_t i = 1; i < size; i++) {
      const Point<TYPE> &p1 = points_[i - 1];
      const Point<TYPE> &p2 = points_[i];
      if (p1.getX() < point.getX() && p2.getX() < point.getX()) {
        continue;
      }
      if (p1.getY() <= point.getY() && p2.getY() >= point.getY()) {
        count++;
      } else if (p1.getY() >= point.getY() && p2.getY() <= point.getY()) {
        count++;
      }
    }

    if (count % 2) {
      return true;
    }

    return false;
  }
  //---------------------------------------------------------------------------------------
  const Points &getPoints() { return points_; }
  //---------------------------------------------------------------------------------------
};