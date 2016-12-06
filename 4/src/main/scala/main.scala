import scala.collection.immutable.ListMap
import scala.io.Source


case class Room(encryptedName: String, sectorId: Int, checksum: String) {
  lazy val correctChecksum: String = encryptedName
    .filter(_.isLetter)
    .foldLeft(ListMap[Char, Int]().withDefaultValue(0))(
      (acc: Map[Char, Int], c: Char) => acc + ((c, acc(c) + 1)))
    .toIndexedSeq.sortBy(x => (-x._2, x._1))
    .map(_._1).take(5).mkString

  lazy val isValid: Boolean = correctChecksum == checksum

  lazy val name: String = encryptedName.map{
    case '-' => ' '
    case c => ((c + sectorId - 'a') % ('z' - 'a' + 1) + 'a').toChar
  }.mkString
}


object Room {
  val regex = "([-a-z]+)-([0-9]+)\\[([a-z]+)\\]".r
  def parseString(s: String): Room = s match {
    case regex(encryptedName, sectorId, checksum) =>
      Room(encryptedName, sectorId.toInt, checksum)
  }
}


object Four {
  def test(room: String, expected: Boolean) {
    if (Room.parseString(room).isValid == expected) {
      print("PASS")
    } else {
      print("FAIL")
    }
    println(s" ${room}")
  }

  def sumSectorIds(input: String) {
    print(input + ": ")
    println(Source.fromURL(getClass.getResource(input))
              .getLines.map(Room.parseString)
              .collect { case r: Room if r.isValid => r.sectorId }
              .sum
    )
  }

  def decrypt(input: String) {
    print(input + ": ")
    println(Source.fromURL(getClass.getResource(input))
              .getLines.map(Room.parseString)
              .map(r => s"${r.name}\t ${r.sectorId}").mkString("\n")
    )

  }

  def main(args: Array[String]) {
    test("aaaaa-bbb-z-y-x-123[abxyz]", true)
    test("a-b-c-d-e-f-g-h-987[abcde]", true)
    test("not-a-real-room-404[oarel]", true)
    test("totally-real-room-200[decoy]", false)
    sumSectorIds("/input.example")
    sumSectorIds("/input")
    decrypt("/input.example-2")
    decrypt("/input")
  }
}
