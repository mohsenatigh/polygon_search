package geospat;

import java.io.FileInputStream;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Random;

import geospat.GeoMatrix.GeoMatrixAccuracy;

public final class Geospat {

  static private GeoMatrix<Float> matrix_=new GeoMatrix<Float>();
  
  //-------------------------------------------------------------------------------------
  static private Point readPoint(String in) throws Exception {
    String[] parts=in.split(",",-1);
    if(parts.length!=2){
      throw new Exception("Invalid point");
    }  
    return new Point(Float.parseFloat(parts[0]),Float.parseFloat(parts[1]));
  }
  //-------------------------------------------------------------------------------------
  static private boolean loadFile(String fileName){
    try{
      FileInputStream str=new FileInputStream(fileName);
      BufferedReader br = new BufferedReader(new InputStreamReader(str));      
      String strLine;
      
      while ((strLine = br.readLine()) != null){
        String[] parts=strLine.split("\\|",-1);
        if (parts.length < 2) {
          continue;
        }
        String id= parts[0];
        ArrayList<Point> points=new ArrayList<Point>();
        for(int i=1;i<parts.length;i++){
            points.add(readPoint(parts[i]));
        }
        matrix_.add(new Polygon(points), Float.parseFloat(id));
      }
      str.close();
    } catch(Exception e){
      return false;
    }  
    return true;
  }
  //-------------------------------------------------------------------------------------
  private static void testPerformance(final Point testPoint,GeoMatrixAccuracy acc,boolean useRandomPoints){
    ArrayList<Float> out=new ArrayList<Float>();
    final int COUNT=10000000;

    ArrayList<Point> points=new ArrayList<Point>();
    
    Random r = new Random();
    for(int i=0;i<COUNT;i++){
      Double x=r.nextDouble()/10000;
      Double y=r.nextDouble()/10000;
      points.add(new Point(testPoint.getX()+x,testPoint.getY()+y));
    }


    long start = System.nanoTime();
    for (int i = 0; i < COUNT; i++) {
      if(useRandomPoints){
        matrix_.query(points.get(i),out,acc,1);
      } else {
        matrix_.query(testPoint,out,acc,1);
      }
      out.clear();
    }
    long end = System.nanoTime();
    System.out.printf("t = %d Micro second\n",(end-start)/1000);
  }
  //-------------------------------------------------------------------------------------
  public static void main(String[] args){
    
    final String help = 
    "-f data file\n"+
    "-a lat\n"+
    "-o long\n"+
    "-H search with high accuracy\n"+
    "-L search with low accuracy\n"+
    "-R use random points\n"+
    "-h show help\n";

    String dataFile="../data.csv";
    Double lat=7.847786843776704;
    Double lng=47.99549694289439;
    Boolean useRandomPoints=false;
    GeoMatrixAccuracy acc=GeoMatrixAccuracy.MEDIUM;

    //parse command lines
    try{
      int size=args.length;
      for (int i=0;i<size;i++){
        String arg=args[i];
        switch(arg){
          case "-h":
            System.out.println(help);  
            System.exit(0);
            break;
          case "-f":
            dataFile=args[++i];
            break;
          case "-a":
            lat=Double.parseDouble(args[++i]);
            break;
          case "-o":
            lng=Double.parseDouble(args[++i]);
            break;
          case "-L":
            acc=GeoMatrixAccuracy.LOW;
            break;
          case "-H":
            acc=GeoMatrixAccuracy.HIGH;
            break;
          case "-R":
            useRandomPoints=true;
            break;
        }
      }
    } catch(Exception e){
        System.out.println("invalid command line args");
        System.exit(0);
    }
    

    System.out.printf("try load data %s \n",dataFile);
    if(loadFile(dataFile)){
      System.out.println("load data successfully\n");
    } else {
      System.out.println("load data failed\n");
      return;
    }
    matrix_.build();

    Point testPoint=new Point(lat,lng);
    ArrayList<Float> out = matrix_.query(testPoint,acc,-1);
    for (Float i : out) {
      System.out.println(i);
    }

    System.out.println("Final test");
    testPerformance(testPoint,acc,useRandomPoints);
  }
}
