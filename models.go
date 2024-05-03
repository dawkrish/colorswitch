package main

import "image/color"

type Mode int

type Ball struct{
	cntrX int32
	cntrY int32
	radius float32
	color color.RGBA
}