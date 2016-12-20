// want: (3x2)(1x5)a -> (1x(1x5)a -> (1xaaaaa  (8)
// 11 -> 9 -> 8

// want: (6x2)(1x5)a -> (1x5)a(1x5)a -> aaaaaaaaaa (10)
// 12 -> 11 -> 10

// want: (6x2)(3x5)abc -> (3x5)a(3x5)abc -> a(3a(3a(3a(3a(3x5)abc -> a(3a(3a(3a(3 a abcabcabcabcabc (28)
//                                                                             12+1+15 = 28
// 13 -> 14 -> 9 + 12 = 21 -> 16 + 12 = 28


import scala.io.Source
import scala.collection.mutable.Queue
import scala.util.matching.Regex.Match

object Main {
  val regex = """(.*?)\((\d+)x(\d+)\)(.*)""".r

  def repeat(input: String, charCount: Int, times: Int): (String, String) =
    (0.until(times).map(_ => input.take(charCount)).mkString, input.drop(charCount))

  def expandMarkersV1(s: String): String = s match {
    case regex(before, charCount, times, after) =>
      val (repeated, rest) = repeat(after, charCount.toInt, times.toInt)
      before + repeated + expandMarkersV1(rest)
    case s => s
  }

  def solveOne() {
    println(
      expandMarkersV1(
        Source.fromURL(getClass.getResource("/input")).getLines.next().filter(_ != ' ')
      ).length()
    )
  }

  def matchToTuple(m: Match): (Int, Int) = {
    val charCount = m.group(1).toInt
    val times = m.group(2).toInt
    (charCount, times)
  }

  case class Marker(m: Match, var repeatCount: Int = 1) {
    val charCount = m.group(1).toInt
    val times = m.group(2).toInt
    def apply(length: Long) = length + ((charCount * (times - 1)) - m.matched.length) * repeatCount
  }

  def expandedV2Length(input: String): Long = {
    val markerRegex = """\((\d+)x(\d+)\)""".r
    val markers = markerRegex.findAllMatchIn(input).map(Marker(_)).toList
    markers.zipWithIndex.foreach((pair: (Marker, Int)) =>
      {
        val (marker, index) = pair
        markers.drop(index + 1).takeWhile(other =>
          other.m.end <= marker.m.end + marker.charCount
        ).foreach(other =>
          other.repeatCount *= marker.times
        )
      })

    markers.foldLeft(input.length.toLong)((length: Long, marker: Marker) => marker(length))
  }

  def solveTwo() {
    val input = Source.fromURL(getClass.getResource("/input")).getLines.next().filter(_ != ' ')
    println(s"Solution: ${expandedV2Length(input)}")
  }

  def test(input: String, expected: Int) {
    val actual = expandedV2Length(input)
    if (actual == expected) {
      print("PASS")
    } else {
      print("FAIL")
    }
    println(s"${input} expected=${expected} actual=${actual}")
  }

  def main(args: Array[String]) {
    test("(3x2)(1x5)a", 8)
    test("(6x2)(1x5)a", 10)
    test("(6x2)(3x5)abc", 28)
    test("(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920)
    solveTwo()
  }
}
