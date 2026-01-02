package item

type FishingRod struct{}

func (FishingRod) MaxCount() int {
	return 1
}

func (FishingRod) EncodeItem() (name string, meta int16) {
	return "minecraft:fishing_rod", 0
}
