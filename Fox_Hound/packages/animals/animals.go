package animals

import (
	"errors"
)

type Animal interface {
	Makesound() string
	Move() string
	Eat() string
	Sleep() string
}

type Swimmer interface {
	Animal
	Swim() string
}

var ErrCannotSwim = errors.New("это животное не умеет плавать")

type Shoebill struct{}

func (s Shoebill) Makesound() string {
	return "Стучит клювом в знак приветствия"
}

func (s Shoebill) Move() string {
	return "Передвигается по заболоченной местности"
}

func (s Shoebill) Eat() string {
	return "Питается рыбой и другими водными обитателями"
}

func (s Shoebill) Sleep() string {
	return "Спит стоя, сложив тяжелую голову на шею"
}

func (s Shoebill) Swim() string {
	return "Медленно плавает в неглубокой воде"
}

type Manul struct{}

func (m Manul) Makesound() string {
	return "Манул издает короткий глухой звук."
}

func (m Manul) Move() string {
	return "Манул движется быстро и бесшумно."
}

func (m Manul) Eat() string {
	return "Манул охотится на грызунов."
}

func (m Manul) Sleep() string {
	return "Манул спит в своем логове."
}

type Kakapo struct{}

func (k Kakapo) Makesound() string {
	return "Какапо издает характерные звуки, напоминающие ухание или бормотание."
}

func (k Kakapo) Move() string {
	return "Какапо двигается медленно, предпочитая передвигаться по земле."
}

func (k Kakapo) Eat() string {
	return "Какапо питается фруктами, листьями и цветами деревьев."
}

func (k Kakapo) Sleep() string {
	return "Какапо спит в укромных местах среди деревьев."
}
