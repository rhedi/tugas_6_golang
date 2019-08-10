package library

import "fmt"


func private()  {
  var s1 mahasiswa
  var s2 = mahasiswa {"Yosep", 20}
  s1.nama = "Aldo"
  s1.umur = 19
  s1.namasaya()
  s2.namasaya()
}

func Public()  {
  var s1 mahasiswa
  var s2 = mahasiswa {"Yosep", 20}
  s1.nama = "Aldo"
  s1.umur = 19
  s1.namasaya()
  s2.namasaya()
}
type mahasiswa struct {
  nama string
  umur int
}
func (s mahasiswa) namasaya()  {
  fmt.Println("nama saya adalah", s.nama, "Dipublic")
  fmt.Println("Umur Saya adalah", s.umur)
}
















// func iniprivate()  { //Jika huruf depanya kecil maka private(iniprivate)
//   fmt.Println("saya di private")
// }
//
// func Inipublic()  { //jika huruf depannya besar berarti public (Inipublic)
//   fmt.Println("saya di public")
//   iniprivate() //untuk memanggil data private diluar folder / berbeda
// }
