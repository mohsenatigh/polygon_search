package geospat;

import java.io.FileInputStream;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.ArrayList;

import geospat.GeoMatrix.GeoMatrixAccuracy;

public final class Geospat {

  static private GeoMatrix<Integer> matrix_=new GeoMatrix<Integer>();
  
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
        matrix_.add(new Polygon(points), Integer.parseInt(id));
      }
      str.close();
    } catch(Exception e){
      return false;
    }  
    return true;
  }
  //-------------------------------------------------------------------------------------
  private static void testPerformance(final Point testPoint,GeoMatrixAccuracy acc){
    ArrayList<Integer> out=new ArrayList<Integer>();
    long start = System.nanoTime();
    for (int i = 0; i < 1000; i++) {
      matrix_.query(testPoint,out,acc,1);
    }
    long end = System.nanoTime();
    System.out.printf("t = %d \n",(end-start)/1000);
  }
  //-------------------------------------------------------------------------------------
  public static void main(String[] args){
    
    final String help = """
    -f data file
    -a lat
    -o long
    -H search with high accuracy
    -L search with low accuracy
    -h show help
    """;

    String dataFile="../data.csv";
    Double lat=7.847786843776704;
    Double lng=47.99549694289439;
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
    ArrayList<Integer> out = matrix_.query(testPoint,acc,-1);
    for (Integer i : out) {
      System.out.println(i);
    }


    //JIT Warmup
    System.out.println("Warmup jit compiler");
    for(int i=0;i<8;i++){
      testPerformance(testPoint,acc);
    }
    
    System.out.println("Final test");
    testPerformance(testPoint,acc);
  }
}
