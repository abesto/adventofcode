import scala.io.Source
import com.twmacinta.util.MD5

object Main {
  MD5.initNativeLibrary(true)

  def md5String(s: String) = new MD5(s).asHex

  def passwordOne(doorId: String, iteration: Long = 0, acc: String = ""): String = if (acc.length == 8) {
    acc
  } else {
    if (iteration % 10000 == 0) {
      println(s"${iteration}\t${doorId}\t${iteration}\t${doorId + iteration.toString}\t${md5String(doorId + iteration.toString)}\t${acc}")
    }
    md5String(doorId + iteration.toString) match {
      case hash if hash.startsWith("00000") => passwordOne(doorId, iteration + 1, acc + hash(5))
      case _ => passwordOne(doorId, iteration + 1, acc)
    }
  }

  def solveOne(doorId: String) {
    println(s"${doorId}: ${passwordOne(doorId)}")
  }

  def passwordTwo(doorId: String, iteration: Long = 0, acc: Array[Option[Char]] = Array(None, None, None, None, None, None, None, None)): Array[Option[Char]] =
    if (acc.forall(_.isDefined)) {
      acc
    } else {
      if (iteration % 10000 == 0) {
        val pwd = acc.map(_.getOrElse('_')).mkString
        println(s"${iteration}\t${doorId}\t${iteration}\t${doorId + iteration.toString}\t${md5String(doorId + iteration.toString)}\t${pwd}")
      }
      md5String(doorId + iteration.toString) match {
        case hash if hash.startsWith("00000") && hash(5).isDigit && hash(5) < '8' && acc(hash(5) - '0').isEmpty =>
          val position: Int = hash(5) - '0'
          val value = hash(6)
          passwordTwo(doorId, iteration + 1, acc.patch(position, Seq(Some(value)), 1))
        case _ => passwordTwo(doorId, iteration + 1, acc)
      }
    }

  def solveTwo(doorId: String) {
    println(s"${doorId}: ${passwordTwo(doorId).map(_.get).mkString}")
  }

  def main(args: Array[String]) {
    // solveOne("abc")
    // solveOne("abbhdwsy")
    // solveTwo("abc")
    solveTwo("abbhdwsy")
  }
}
