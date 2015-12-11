#!/bin/bash

down=$(cat input | tr -d '(' | wc -c)
up=$(cat input | tr -d ')' | wc -c)
expr $up - $down
