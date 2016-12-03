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
              .flatMap(_.split(' ').filter(_.length > 0).map(_.toInt))
              .sliding(9, 9)
              .flatMap((x: Seq[Int]) => Seq(
                         x(0), x(3), x(6),
                         x(1), x(4), x(7),
                         x(2), x(5), x(8)
                       ))
              .sliding(3, 3).map(_.sorted)
              .filter((t: Seq[Int]) => t(2) < t(0) + t(1))
              .length)
  }

  def main(args: Array[String]) {
    solveOne()
    solveTwo()
  }
}
