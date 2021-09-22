package main

func main() {
	dev1 := &Developer{
		name: "John",
		salary: 100000,
	}
	dev2 := &Developer{
		name: "Julia",
		salary: 150000,
	}
	dev3 := &Developer{
		name: "Mark",
		salary: 100000,
	}
	dev4 := &Developer{
		name: "David",
		salary: 150000,
	}

	tmL1 := &TeamLeader{
		name: "Carl",
		salary: 200000,
	}

	tmL2 := &TeamLeader{
		name: "Sophie",
		salary: 220000,
	}

	man1 := &Manager{
		name: "Daniel",
		salary: 25000,
	}

	tmL1.add(dev1)
	tmL1.add(dev2)
	tmL2.add(dev1)
	tmL2.add(dev3)
	tmL2.add(dev4)
	man1.add(tmL1)
	man1.add(tmL2)
	man1.print()
}
