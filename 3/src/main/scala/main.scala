import scala.io.Source

object Three {
  def solveOne() {
    println(Source.fromURL(getClass.getResource("/input"))
              .getLines
              .map(_.split(' ').filter(_.length > 0).map(_.toInt).sorted)
              .filter((t: Array[Int]) => t(2) < t(0) + t(1))
              .length)
  }

  def solveTwo() {
    println(Source.fromURL(getClass.getResource("/input"))
              .getLines
              .map(_.split(' ').filter(_.length > 0).map(_.toInt))
              .sliding(3, 3).flatMap(_.transpose)
              .map(_.sorted)
              .filter((t: Seq[Int]) => t(2) < t(0) + t(1))
              .length)
  }

  def main(args: Array[String]) {
    solveOne()
    solveTwo()
  }
}
