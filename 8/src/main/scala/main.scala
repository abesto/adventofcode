import scala.io.Source

class Display {
  var pixels: Array[Array[Boolean]] = 0.until(50).map(
    _ => 0.until(6).map(_ => false).toArray
  ).toArray

  val pixelRepr: Map[Boolean, Char] = Map(false -> '.', true -> '#')

  override def toString(): String = {
    pixels.transpose.map(row => {
                           row.map(pixelRepr(_)).mkString
                         }).mkString("\n")
  }

  def rect(columns: Int, rows: Int): Unit = {
    0.until(columns).foreach(column =>
      0.until(rows).foreach(row => pixels(column)(row) = true)
    )
  }

  def rotate(xs: Array[Boolean], n: Int): Array[Boolean] =
    0.until(xs.length).map(index => xs((xs.length * 2 + index - n) % xs.length)).toArray

  def down(column: Int, by: Int): Unit = {
    pixels(column) = rotate(pixels(column), by)
  }

  def right(row: Int, by: Int): Unit = {
    val transposed = pixels.transpose
    transposed(row) = rotate(transposed(row), by)
    pixels = transposed.transpose
  }
}

object Main {
  def applyWithPrefix(d: Display, line: String, prefix: String, splitBy: String, f: ((Int, Int) => Unit)): Unit = {
    if (line.startsWith(prefix)) {
      val rest = line.substring(prefix.length)
      val items = rest.split(splitBy)
      f(items(0).toInt, items(1).toInt)
    }
  }

  def applyLine(d: Display, line: String): Unit = {
    applyWithPrefix(d, line, "rect ", "x", d.rect(_, _))
    applyWithPrefix(d, line, "rotate row y=", " by ", d.right(_, _))
    applyWithPrefix(d, line, "rotate column x=", " by ", d.down(_, _))
  }

  def animate(lines: Iterator[String]): Display = {
    val d = new Display()
    println(d)
    println("Init")
    lines.foreach(line => {
                    Thread.sleep(1000)
                    print(27.toChar + "[7A")
                    applyLine(d, line)
                    println(d)
                    println(line + "                                   ")
                  })
    d
  }

  def solveOne() {
    val d = animate(Source.fromURL(getClass.getResource("/input")).getLines)
    val solution = d.pixels.map(_.count(_ == true)).sum
    println(s"Number of lit pixels: ${solution}")
  }

  def main(args: Array[String]) {
    solveOne()
  }
}
