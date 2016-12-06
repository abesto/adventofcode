import scala.io.Source



object Main {
  def solve(inputFile: String, selectBy: ((Char, Int)) => Int) {
    print(inputFile + ": ")
    println(Source.fromURL(getClass.getResource(inputFile)).getLines
              .foldLeft(Map[Int, Map[Char, Int]]().withDefaultValue(Map().withDefaultValue(0)))(
              (m, line) => line.view.zipWithIndex.foldLeft(m)((m, c) => m + ((c._2, (m(c._2) + ((c._1, m(c._2)(c._1) + 1)))))))  // muhahaha
              .toSeq.sortBy(_._1).map(_._2.maxBy(selectBy)._1).mkString  // MUHAHAHAHA
    )
  }

  def solveOne(inputFile: String) = solve(inputFile, _._2)
  def solveTwo(inputFile: String) = solve(inputFile, -_._2)  // Because obviously -_._2 is valid Scala

  def main(args: Array[String]) {
    solveOne("/input.example1")
    solveOne("/input")
    solveTwo("/input.example1")
    solveTwo("/input")
  }
}
