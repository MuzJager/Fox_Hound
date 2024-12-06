package animals

import (
	"errors"
)

type Animal interface {
	Makesound() string
	Move() string
	Eat() string
	Sleep() string
	HighSpeed() string
	Size() string
	ClimbTree() string
	RecognizeDiseases() string
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

func (s Shoebill) HighSpeed() string {
	return "Медленная скорость передвижения."
}

func (s Shoebill) Size() string {
	return "Среднего размера птица."
}

func (s Shoebill) ClimbTree() string {
	return "Не умеет лазить по деревьям."
}

func (s Shoebill) RecognizeDiseases() string {
	return "Неизвестно, как распознает болезни."
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

func (m Manul) HighSpeed() string {
	return "Высокая скорость, короткие рывки."
}

func (m Manul) Size() string {
	return "Маленький по размеру, как кошка."
}

func (m Manul) ClimbTree() string {
	return "Манул умеет лазить по деревьям."
}

func (m Manul) RecognizeDiseases() string {
	return "Предположительно распознает болезни по поведению."
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

func (k Kakapo) HighSpeed() string {
	return "Очень медленно передвигается."
}

func (k Kakapo) Size() string {
	return "Большая птица."
}

func (k Kakapo) ClimbTree() string {
	return "Да, может лазить по деревьям."
}

func (k Kakapo) RecognizeDiseases() string {
	return "Неизвестно, как распознает болезни."
}
