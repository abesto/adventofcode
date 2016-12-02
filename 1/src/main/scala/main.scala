import scala.io.Source

case class Vector(x: Int, y: Int) {
  def +(v: Vector) = Vector(x + v.x, y + v.y)
  def *(s: Int) = Vector(x * s, y * s)
  def left = Vector(-y, x)
  def right = Vector(y, -x)
}

case class State(position: Vector = Vector(0, 0),
                 heading: Vector = Vector(0, -1),
                 distance: Int = 0)
{
  def left = State(position, heading.left, distance)
  def right = State(position, heading.right, distance)
  def distanceDigit(d: Int) = State(position, heading, distance * 10 + d)
  def move = State(position + heading * distance, heading, 0)
}

object One {
  def applyChar(s: State, c: Char) =
    c match {
      case 'L' => s.left
      case 'R' => s.right
      case ' ' => s
      case ',' | '\n' => s.move
      case _ => {
        assert(c.isDigit, c)
        s.distanceDigit(c.asDigit)
      }
    }

  def solveOne(inputFile: String) {
    val hq: Vector = Source.fromURL(getClass.getResource(inputFile))
      .foldLeft(State())(applyChar).move.position
    print(inputFile + ": ")
    println(math.abs(hq.x) + math.abs(hq.y))
  }

  def solveTwo(inputFile: String) {
    print(inputFile + ": ")

    var touched: Seq[Vector] = Seq()
    var state: State = State()

    for (c <- Source.fromURL(getClass.getResource(inputFile))) {
      val newState = applyChar(state, c)
      if (newState.distance == 0) {
        // We just moved
        val touchedNow = 1.to(state.distance).map(state.position + state.heading * _)
        for (pos <- touchedNow) {
          // This would break if going in a straight line could hit the same point twice. For example if our plane was on a sphere.
          if (touched.contains(pos)) {
            println(math.abs(pos.x) + math.abs(pos.y))
            return
          }
        }
        touched = touched ++ touchedNow
      }
      state = newState
    }
  }

  def main(args: Array[String]) {
    println("First part")
    solveOne("/input.example1-1")
    solveOne("/input.example1-2")
    solveOne("/input.example1-3")
    solveOne("/input.test1-1")
    solveOne("/input")
    println("Second part")
    solveTwo("/input.example2-1")
    solveTwo("/input")
  }
}
