#pragma once
#include "Point.hpp"
#include "Polygon.hpp"
#include "Rectangle.hpp"
#include <iostream>
#include <unordered_map>
#include <vector>

template <typename TYPE, typename DATATYPE, uint32_t ROW = 1000>
class GeoMatrix {
public:
  enum GeoMatrixAccuracy { HIGH, MEDIUM, LOW };
  
private:
  using ItemInfo = std::pair<Polygon<TYPE>, DATATYPE>;
  using MatrixPoint = Point<uint64_t>;
  using ItemsPointerList = std::vector<const ItemInfo *>;

  std::unordered_map<MatrixPoint, ItemsPointerList, PointHash<uint64_t>> map_;
  std::vector<ItemInfo> items_;
  Rectangle<TYPE> bBox_;
  //-------------------------------------------------------------------------------------
  MatrixPoint getMatrixPoint(const Point<TYPE> &in) const {
    Point<TYPE> ext = bBox_.getLength();
    TYPE slotX = (ext.getX() / ROW);
    TYPE slotY = (ext.getY() / ROW);

    TYPE x = ((in.getX() - bBox_.getStart().getX()) / slotX);
    TYPE y = ((in.getY() - bBox_.getStart().getY()) / slotY);

    x = std::min<TYPE>(x, ROW - 1);
    y = std::min<TYPE>(y, ROW - 1);

    return MatrixPoint(static_cast<uint64_t>(x), static_cast<uint64_t>(y));
  }
  //-------------------------------------------------------------------------------------
  void coverArea(const MatrixPoint &start, const MatrixPoint &end,
                 const ItemInfo &info) {
    for (uint64_t x = start.getX(); x <= end.getX(); x++) {
      for (uint64_t y = start.getY(); y <= end.getY(); y++) {
        map_[MatrixPoint(x, y)].push_back(&info);
      }
    }
  }

public:
  //-------------------------------------------------------------------------------------
  void add(const Polygon<TYPE> &poly, const DATATYPE &data) {
    auto rect = poly.getBoundingBox();

    if (items_.empty()) {
      bBox_ = rect;
    } else {
      bBox_ += rect;
    }

    items_.push_back(ItemInfo(poly, data));
  }
  //-------------------------------------------------------------------------------------
  void build() {
    for (auto &i : items_) {
      const Rectangle<TYPE> &rect = i.first.getBoundingBox();
      auto start = getMatrixPoint(rect.getStart());
      auto end = getMatrixPoint(rect.getEnd());
      coverArea(start, end, i);
    }
  }
  //-------------------------------------------------------------------------------------
  bool query(const Point<TYPE> &point, std::vector<DATATYPE> &out,
             GeoMatrixAccuracy acc = MEDIUM, size_t maxCount = -1) const {
    MatrixPoint mp = getMatrixPoint(point);

    auto items = map_.find(mp);

    if (items == map_.end()) {
      return false;
    }

    if(items->second.size()==1){
      out.push_back(items->second[0]->second);
      return true;
    }

    for (const auto &i : items->second) {
      if (out.size() == maxCount) {
        return true;
      }

      switch (acc) {
      case HIGH:
        if (i->first.pointIsInside(point)) {
          out.push_back(i->second);
          return true;
        }
        break;
      case MEDIUM:
        if (i->first.getBoundingBox().pointInside(point)) {
          out.push_back(i->second);
        }
        break;
      case LOW:
        out.push_back(i->second);
        break;
      }
    }
    return true;
  }
  //-------------------------------------------------------------------------------------
  std::vector<DATATYPE> query(const Point<TYPE> &point,
                              GeoMatrixAccuracy acc = HIGH,
                              int maxCount = -1) const {
    std::vector<DATATYPE> out;
    query(point, out, acc, maxCount);
    return out;
  }
  //-------------------------------------------------------------------------------------
};