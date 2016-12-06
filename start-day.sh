#!/bin/bash
set -xeuo pipefail

n=$1; shift

mkdir -p $n/src/main/{scala,resources}
echo "name := \"$n\"" > $n/build.sbt
cat > $n/src/main/scala/main.scala <<EOF
import scala.io.Source

object Main {
  def main(args: Array[String]) {
  }
}
EOF
cd $n
sbt ensimeConfig ensimeConfigProject
emacsclient -n src/main/scala/main.scala
