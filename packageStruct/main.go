package main

import "donghee.com/structPackage"

func main() {
	//structPackage.ShowPackageInfo()
	original_sp := structPackage.StructPackage{}
	structPackage.ShowPackageInfo(&original_sp, "stopBlcck", "v1.1.1")
}
