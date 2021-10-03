#pragma once
#include "Point.hpp"
#include <limits>
#include <vector>
template <typename TYPE> class Rectangle {
private:
  Point<TYPE> start_;
  Point<TYPE> end_;

public:
  //-----------------------------------------------------------------------------------
  Rectangle() {}
  Rectangle(const TYPE startX, const TYPE startY, const TYPE endX,
            const TYPE endY) {
    start_ = Point(startX, startY);
    end_ = Point(endX, endY);
  }

  Rectangle(const Point<TYPE> &start, const Point<TYPE> &end)
      : start_(start), end_(end) {}
  //-----------------------------------------------------------------------------------
  void setStart(const Point<TYPE> &start) { start_ = start; }
  //-----------------------------------------------------------------------------------
  void setEnd(const Point<TYPE> &end) { end_ = end; }
  //-----------------------------------------------------------------------------------
  const Point<TYPE> &getStart() const { return start_; }
  //-----------------------------------------------------------------------------------
  const Point<TYPE> &getEnd() const { return end_; }
  //-----------------------------------------------------------------------------------
  bool pointInside(const Point<TYPE> &point) const {

    auto checkV = [](TYPE start, TYPE end, TYPE val) -> bool {
      if (val >= start && val <= end) {
        return true;
      }
      if (val <= start && val >= end) {
        return true;
      }
      return false;
    };

    return (checkV(start_.getX(), end_.getX(), point.getX()) &&
            checkV(start_.getY(), end_.getY(), point.getY()));
  }
  //-----------------------------------------------------------------------------------
  bool rectInside(const Rectangle<TYPE> &rect) const {
    return (pointInside(rect.getStart()) && pointInside(rect.getEnd()));
  }
  //-----------------------------------------------------------------------------------
  bool rectOverlap(const Rectangle<TYPE> &rect) const {
    return (pointInside(rect.getStart()) || pointInside(rect.getEnd()));
  }
  //-----------------------------------------------------------------------------------
  Point<TYPE> getLength() const {
    TYPE xL = end_.getX() - start_.getX();
    TYPE yL = end_.getY() - start_.getY();
    return Point<TYPE>(xL, yL);
  }
  //-----------------------------------------------------------------------------------
  bool operator==(const Rectangle<TYPE> &other) const {
    Point<TYPE> myLen = getLength();
    Point<TYPE> otherLen = other.getLength();
    return (myLen == otherLen);
  }
  //-----------------------------------------------------------------------------------
  bool operator>(const Rectangle<TYPE> &other) const {
    Point<TYPE> myLen = getLength();
    Point<TYPE> otherLen = other.getLength();
    return ((myLen.getX() * myLen.getY()) >
            (otherLen.getY() * otherLen.getY()));
  }
  //-----------------------------------------------------------------------------------
  bool operator<(const Rectangle<TYPE> &other) const {
    Point<TYPE> myLen = getLength();
    Point<TYPE> otherLen = other.getLength();
    return ((myLen.getX() * myLen.getY()) <
            (otherLen.getY() * otherLen.getY()));
  }
  //-----------------------------------------------------------------------------------
  void operator+=(const Rectangle<TYPE> &other) {
    /*
      Start
    */
    Point<TYPE> nStart(std::min<TYPE>(other.start_.getX(), start_.getX()),
                       std::min<TYPE>(other.start_.getY(), start_.getY()));

    Point<TYPE> nEnd(std::max<TYPE>(other.end_.getX(), end_.getX()),
                     std::max<TYPE>(other.end_.getY(), end_.getY()));

    setStart(nStart);
    setEnd(nEnd);
  }
};