package world

import (
	"agar-life/math/crd"
	"agar-life/object/alive/animal"
	"agar-life/object/alive/animal/behavior"
	"agar-life/object/alive/animal/behavior/ai"
	"agar-life/object/alive/animal/species"
	"agar-life/object/alive/plant"
	sp "agar-life/object/alive/plant/species"
	_const "agar-life/world/const"
	"agar-life/world/frame"
	"agar-life/world/frame/grid"
	"encoding/json"
	"io"
)

type WorldJSON struct {
	W           float64  `json:"W"`
	H           float64  `json:"H"`
	Cycle       uint   `json:"Cycle"`
	Plants      []Plant  `json:"Plants"`
	Animals     []Animal `json:"Animals"`
	CountPlant  int      `json:"CountPlant"`
	CountAnimal int      `json:"CountAnimal"`
}

type Plant struct {
	Color    string  `json:"Color"`
	Size     float64 `json:"Size"`
	ViewSize float64 `json:"ViewSize"`
	GrowSize float64 `json:"GrowSize"`
	Hidden   bool    `json:"Hidden"`
	X        float64 `json:"X"`
	Y        float64 `json:"Y"`
	Deed     bool    `json:"Deed"`
	Group    string  `json:"Group"`
	ID       int     `json:"ID"`
	Danger   bool    `json:"Danger"`
	Edible   bool    `json:"Edible"`
}

type Animal struct {
	Color    string  `json:"Color"`
	Size     float64 `json:"Size"`
	ViewSize float64 `json:"ViewSize"`
	GrowSize float64 `json:"GrowSize"`
	Hidden   bool    `json:"Hidden"`
	X        float64 `json:"X"`
	Y        float64 `json:"Y"`
	Deed     bool    `json:"Deed"`
	Group    string  `json:"Group"`
	ID       int     `json:"ID"`
	Danger   bool    `json:"Danger"`
	Edible   bool    `json:"Edible"`
	ChCrd    struct {
		X float64 `json:"X"`
		Y float64 `json:"Y"`
	} `json:"ChCrd"`
	OldDirection struct {
		X float64 `json:"X"`
		Y float64 `json:"Y"`
	} `json:"OldDirection"`
	OldDist float64 `json:"OldDist"`
	Inertia struct {
		Direction struct {
			X float64 `json:"X"`
			Y float64 `json:"Y"`
		} `json:"Direction"`
		Speed        float64 `json:"Speed"`
		Acceleration float64 `json:"Acceleration"`
	} `json:"Inertia"`
	Speed     float64  `json:"Speed"`
	Vision    float64  `json:"Vision"`
	CycleGlue uint   `json:"CycleGlue"`
	Children  []Animal `json:"Children"`
	Behavior  struct {
		Name string `json:"Name"`
	} `json:"Behavior"`
	Parent *int `json:"Parent"`
}

type export struct {
	W, H                    float64
	Cycle                   uint
	Plants                  []plant.Plant
	Animals                 []animal.Animal
	CountPlant, CountAnimal uint
}

func NewWorldFromFile(reader io.Reader) World {
	var data WorldJSON
	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		panic(err)
	}
	w, h := data.W, data.H
	world := World{
		w:          w,
		h:          h,
		gridPlant:  grid.NewArray(_const.GridSize, w, h),
		gridAnimal: grid.NewArray(_const.GridSize, w, h),
		animal:     frame.NewFrame(len(data.Animals), w, h),
		plant:      frame.NewFrame(len(data.Plants), w, h),
		cycle:      data.Cycle,
	}
	index := 0
	for i := 0; i < len(data.Animals); i++ {
		an := data.Animals[i]
		if an.Parent == nil {
			continue
		}
		el := createAnimalFromJSON(an, data.W, data.H, nil)
		world.gridAnimal.Set(el.GetX(), el.GetY(), el.GetSize(), index)
		world.animal.Set(index, el)
		index++
		if len(an.Children) > 0 {
			for _, an := range an.Children {
				el1 := createAnimalFromJSON(an, data.W, data.H, el)
				world.gridAnimal.Set(el1.GetX(), el1.GetY(), el1.GetSize(), index)
				world.animal.Set(index, el1)
				index++
			}
		}
	}
	for i := 0; i < len(data.Plants); i++ {
		el := createPlantFromJSON(data.Plants[i])
		world.gridPlant.Set(el.GetX(), el.GetY(), el.GetSize(), i)
		world.plant.Set(i, el)
	}
	return world
}

func (w *World) ExportWorld() string {
	animalExport := w.animal.All()
	animals := make([]animal.Animal, 0, len(animalExport))
	for _, el := range animalExport {
		if el == nil {
			continue
		}
		animals = append(animals, el.(animal.Animal))
	}
	plantExport := w.plant.All()
	plants := make([]plant.Plant, 0, len(plantExport))
	for _, el := range plantExport {
		if el == nil {
			continue
		}
		plants = append(plants, el.(plant.Plant))
	}
	exp := export{
		W:           w.w,
		H:           w.h,
		Cycle:       w.cycle,
		Plants:      plants,
		Animals:     animals,
		CountPlant:  w.countPlant,
		CountAnimal: w.countAnimal,
	}
	jData, err := json.MarshalIndent(exp, "", "\t")
	if err != nil {
		println(err)
	}
	return string(jData)
}

func createPlantFromJSON(an Plant) plant.Plant {
	var el plant.Plant
	if an.Danger {
		el = sp.NewPoison()
	} else {
		el = sp.NewPlant()
	}
	el.SetColor(an.Color)
	el.SetSize(an.Size)
	el.SetViewSize(an.ViewSize)
	el.Revive()
	el.SetCrd(crd.NewCrd(an.X, an.Y))
	el.SetGroup(an.Group)
	el.SetID(an.ID)
	el.SetDanger(an.Danger)
	el.SetEdible(an.Edible)
	return el
}

func createAnimalFromJSON(an Animal, w, h float64, parent animal.Animal) animal.Animal {
	var el animal.Animal
	if an.Behavior.Name == ai.AIV1Name {
		el = species.NewBeast(ai.NewAiv1(w, h))
	}
	if an.Behavior.Name == behavior.FollowerName {
		el = species.NewBeast(behavior.NewFollower())
	}
	if el == nil {
		panic("unexpected behavior")
	}
	if parent != nil {
		el.SetParent(parent)
	}
	el.SetColor(an.Color)
	el.SetSize(an.Size)
	el.SetViewSize(an.ViewSize)
	el.Revive()
	el.SetCrd(crd.NewCrd(an.X, an.Y))
	el.SetGroup(an.Group)
	el.SetID(an.ID)
	el.SetDanger(an.Danger)
	el.SetEdible(an.Edible)
	el.SetSpeed(an.Speed)
	el.SetVision(an.Vision)
	el.SetGlueTime(an.CycleGlue)
	el.SetInertiaImport(crd.NewCrd(an.Inertia.Direction.X, an.Inertia.Direction.Y), an.Inertia.Speed, an.Inertia.Acceleration)
	return el
}
