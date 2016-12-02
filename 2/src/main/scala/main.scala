import scala.io.Source

object Two {
  case class Keypad(layout: String) {
    val matrix: Seq[Seq[Char]] = layout.stripMargin.lines.map(_.toSeq).toSeq

    def isValid(p: Vector): Boolean =
        math.max(0, math.min(matrix(0).length - 1, p.x)) == p.x &&
        math.max(0, math.min(matrix.length - 1, p.y)) == p.y &&
        this(p) != '.'
    def apply(position: Vector) = matrix(position.y)(position.x)
  }

  case class Vector(x: Int, y: Int) {
    def +(v: Vector) = Vector(x + v.x, y + v.y)
  }

  case class State(position: Vector, code: Seq[Char] = Seq())(implicit keypad: Keypad) {
    def value = keypad(position)
    def move(v: Vector) = {
      val newPosition = position + v
      if (keypad.isValid(newPosition)) {
        State(newPosition, code)
      } else {
        State(position, code)
      }
    }
    def press = State(position, code :+ value)
    def apply(o: Operation) = o.apply(this)
  }

  trait Operation {
    def apply(s: State): State
  }
  case class MoveOperation(v: Vector) extends Operation {
    def apply(s: State) = s.move(v)
  }
  val up = MoveOperation(Vector(0, -1))
  val left = MoveOperation(Vector(-1, 0))
  val down = MoveOperation(Vector(0, 1))
  val right = MoveOperation(Vector(1, 0))
  object press extends Operation {
    def apply(s: State) = s.press
  }

  def solve(startPosition: Vector, inputFile: String)(implicit keypad: Keypad) =
    Source.fromURL(getClass.getResource(inputFile)).map {
      case 'U' => up
      case 'D' => down
      case 'L' => left
      case 'R' => right
      case '\n' => press
    }.foldLeft(State(startPosition))(_.apply(_)).code.mkString

  def solveOne(inputFile: String): String = {
    solve(Vector(1, 1), inputFile)(
      Keypad("""123
               |345
               |789"""))
  }

  def solveTwo(inputFile: String): String = {
    solve(Vector(0, 2), inputFile)(
      Keypad(
        """..1..
          |.234.
          |56789
          |.ABC.
          |..D.."""))
  }

  def main(args: Array[String]) {
    println(s"1/example-1: ${solveOne("/input.example-1")}")
    println(s"1: ${solveOne("/input")}")
    println(s"2/example-1: ${solveTwo("/input.example-1")}")
    println(s"2: ${solveTwo("/input")}")
  }
}
