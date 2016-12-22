import scala.io.Source
import scala.collection.mutable.HashMap

object Main {
  trait CanReceiveChip {
    def receiveChip(chip: Int)
  }

  case class Output(id: Int) extends CanReceiveChip {
    var chips: Seq[Int] = Seq()
    override def receiveChip(chip: Int) {
      chips :+= chip
      println(s"$this received chip $chip")
    }

    override def toString(): String = s"Output $id"
  }

  case class Bot(id: Int, var lowerTo: CanReceiveChip, var higherTo: CanReceiveChip) extends CanReceiveChip {
    var chips: Seq[Int] = Seq()

    override def receiveChip(chip: Int) {
      println(s"$this received chip $chip")
      chips :+= chip
      if (chips.length == 2) {
        val sorted = chips.sorted
        val (lower, higher) = (sorted(0), sorted(1))
        println(s"$this has two chips, passing $lower to $lowerTo and $higher to $higherTo")
        lowerTo.receiveChip(lower)
        higherTo.receiveChip(higher)
        chips = Seq()
      }
    }

    override def toString(): String = s"Bot $id"
  }

  class Factory {
    val bots: HashMap[Int, Bot] = HashMap.empty
    val outputs: HashMap[Int, Output] = HashMap.empty

    def bot(id: Int): Bot = {
      if (!bots.contains(id)) {
        bots.put(id, Bot(id, null, null))
      }
      bots(id)
    }

    def output(id: Int): Output = {
      if (!outputs.contains(id)) {
        outputs.put(id, Output(id))
      }
      outputs(id)
    }

    def get(t: String, id: Int): CanReceiveChip = t match {
      case "bot" => bot(id)
      case "output" => output(id)
    }

    def gives(who: Bot, lowerTo: CanReceiveChip, higherTo: CanReceiveChip): Unit = {
      who.lowerTo = lowerTo
      who.higherTo = higherTo
    }
  }

  def solveOne(): Factory = {
    var valueGoesToLines: Seq[String] = Seq()
    val factory = new Factory()
    val botGivesRegex = """bot (\d+) gives low to (\w+) (\d+) and high to (\w+) (\d+)""".r
    val valueGoesToRegex = """value (\d+) goes to bot (\d+)""".r
    Source.fromURL(getClass.getResource("/input")).getLines
      .foreach({
                 case botGivesRegex(giverIdStr, lowType, lowToStr, highType, highToStr) =>
                   val giver = factory.bot(giverIdStr.toInt)
                   val lowTo = factory.get(lowType, lowToStr.toInt)
                   val highTo = factory.get(highType, highToStr.toInt)
                   factory.gives(giver, lowTo, highTo)
                 case line => valueGoesToLines :+= line
      })
    valueGoesToLines.foreach({
                               case valueGoesToRegex(valueStr, botStr) =>
                                 factory.bot(botStr.toInt).receiveChip(valueStr.toInt)
                             })
    factory
  }

  def solveTwo() {
    val o = solveOne().outputs
    println(o(0).chips.head * o(1).chips.head * o(2).chips.head)
  }

  def test() {
    val f = new Factory()
    f.gives(f.bot(2), f.bot(1), f.bot(0))
    f.gives(f.bot(1), f.output(1), f.bot(0))
    f.gives(f.bot(0), f.output(2), f.output(0))
    f.bot(2).receiveChip(5)
    f.bot(1).receiveChip(3)
    f.bot(2).receiveChip(2)
  }

  def main(args: Array[String]) {
    solveTwo()
  }
}
