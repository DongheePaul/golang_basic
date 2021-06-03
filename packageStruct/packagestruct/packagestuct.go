package structPackage

import "fmt"

type StructPackage struct {
	Name    string
	Version string
}

// func ShowPackageInfo() {
// 	sp := StructPackage{"goBLock", "v1.0.0"}
// 	fmt.Println("Package Name : %s, version : %s\n", sp.Name, sp.Version)
// }

func ShowPackageInfo(sp *StructPackage, name string, ver string) {
	//sp := StructPackage{"goBLock", "v1.0.0"}
	//sp = new(StructPackage)
	sp.Name = name
	sp.Version = ver
	fmt.Println("Package Name : %s, version : %s\n", sp.Name, sp.Version)
}
