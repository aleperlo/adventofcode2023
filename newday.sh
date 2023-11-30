#!/bin/sh

day_dir=$(printf day%02d $1)
mkdir -p $day_dir/part01 $day_dir/part02
cp main.go $day_dir/part01
cp makefile $day_dir/part01
cp main.go $day_dir/part02
cp makefile $day_dir/part02
