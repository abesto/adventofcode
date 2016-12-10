import scala.io.Source

object Main {
  def isAbba(cs: Seq[Char]): Boolean =
    cs.length == 4 &&
      cs(0) == cs(3) &&
      cs(1) == cs(2) &&
      cs(0) != cs(1)

  def isAba(cs: Seq[Char]): Boolean =
    cs.length == 3 &&
      cs(0) == cs(2) &&
      cs(1) != cs(2)

  def abaToBab(cs: Seq[Char]): Seq[Char] = Seq(cs(1), cs(0), cs(1))

  case class IPFragment(cs: Seq[Char], isHypernet: Boolean) {
    def hasAbba: Boolean = cs.length > 0 && cs.sliding(4).exists(isAbba)
    def allAbas = cs.sliding(3).filter(isAba)
    def hasAba = cs.sliding(3).exists(isAba)
    def hasBab(aba: Seq[Char]) = cs.sliding(3).contains(abaToBab(aba))
    def +(c: Char): IPFragment = IPFragment(cs :+ c, isHypernet)
  }

  case class IP(fragments: Seq[IPFragment]) {
    def hasTLS: Boolean = !fragments.filter(_.isHypernet).exists(_.hasAbba) && fragments.filterNot(_.isHypernet).exists(_.hasAbba)

    def hasSSL: Boolean =
      fragments.exists(_.hasAba) &&
      fragments.filterNot(_.isHypernet).flatMap(_.allAbas)
        .exists(aba => fragments.filter(_.isHypernet).exists(_.hasBab(aba)))
  }

  def parseChar(fragments: Seq[IPFragment], c: Char): Seq[IPFragment] = c match {
    case '[' => fragments :+ IPFragment(Seq(), true)
    case ']' => fragments :+ IPFragment(Seq(), false)
    case _   => fragments.dropRight(1) :+ (fragments.last + c)
  }

  def parse(s: String): IP = IP(s.foldLeft(Seq[IPFragment](IPFragment(Seq(), false)))(parseChar))

  def hasTLS(s: String) = parse(s).hasTLS

  def hasSSL(s: String) = parse(s).hasSSL

  def testOne(s: String, expected: Boolean) {
    val actual = hasTLS(s)
    if (actual == expected) {
      print("PASS")
    } else {
      print("FAIL")
    }
    println(s" ${s} actual=${actual} expected=${expected}")
  }

  def testTwo(s: String, expected: Boolean) {
    val actual = hasSSL(s)
    if (actual == expected) {
      print("PASS")
    } else {
      print("FAIL")
    }
    println(s" ${s} actual=${actual} expected=${expected}")
  }

  def solveOne() {
    print("Number of IPs supporting TLS in input: ")
    println(
      Source.fromURL(getClass.getResource("input"))
        .getLines.filter(hasTLS).length)
  }

  def solveTwo() {
    print("Number of IPs supporting SSL in input: ")
    println(
      Source.fromURL(getClass.getResource("input"))
        .getLines.filter(hasSSL).length)
  }

  def main(args: Array[String]) {
    testOne("abba[mnop]qrst", true)
    testOne("abcd[bddb]xyyx", false)
    testOne("aaaa[qwer]tyui", false)
    testOne("ioxxoj[asdfgh]zxcvbn", true)
    solveOne()
    testTwo("aba[bab]xyz", true)
    testTwo("xyx[xyx]xyx", false)
    testTwo("aaa[kek]eke", true)
    testTwo("zazbz[bzb]cdb", true)
    solveTwo()
  }
}
