package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interactive Loker")
	fmt.Println("---------------------")

	fmt.Print("init ")
	var total int
	fmt.Scan(&total)

	if total <= 0 {
		fmt.Println("Jumlah loker minimal 1 atau lebih dan format input harus berupa angka!")
	} else {
		fmt.Println("Berhasil membuat loker dengan jumlah", total)

		i := 0

		for {
			var cmd, jenis, id string

			fmt.Scan(&cmd)

			switch strings.ToLower(cmd) {
			case "input":
				i += 1

				if i > total {
					fmt.Println("Maaf loker sudah penuh")
				} else {
					fmt.Scan(&jenis, &id)
					fmt.Println("Kartu identitas tersimpan di loker nomer ", i)
					//fmt.Println("no", i, "\nid:", id, "\ntipe:", jenis)

					f, err := os.OpenFile("dataloker.txt", os.O_APPEND|os.O_WRONLY, 0600)
					if err != nil {
						panic(err)
					}
					defer f.Close()

					new_data := strconv.Itoa(i) + " " + jenis + " " + id
					if _, err = f.WriteString(new_data + "\n"); err != nil {
						panic(err)
					}
				}
			case "status":

				fmt.Println("No Loker", "Tipe Identitas", "No Identitas")

				data, err := ioutil.ReadFile("dataloker.txt")
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(string(data))

			case "leave":
				data, err := ioutil.ReadFile("dataloker.txt")
				if err != nil {
					fmt.Println(err)
				}

				a := strings.Split(string(data), "\n")

				fmt.Scan(&i)
				// Remove the element at index i from a.
				a[i-1] = a[len(a)-1] // Copy last element to index i.
				a[len(a)-1] = ""     // Erase last element (write zero value).
				a = a[:len(a)-1]     // Truncate slice.

				modified_data := []byte(strings.Join(a, "\n") + "\n")
				// the WriteFile method returns an error if unsuccessful
				errn := ioutil.WriteFile("dataloker.txt", modified_data, 0777)
				// handle this error
				if errn != nil {
					// print it out
					fmt.Println(err)
				}

				fmt.Println("Loker nomer", i, "berhasil dikosongkan")
				i = i - 1
			case "find":
				fmt.Scan(&id)
				data, err := ioutil.ReadFile("dataloker.txt")
				if err != nil {
					fmt.Println(err)
				}
				a := strings.Split(string(data), "\n")
				for x := 0; x < total; x++ {
					listdata := a[x]
					r := strings.Split(listdata, " ")
					k, found := Find(r, id)
					if id == r[2] {
						fmt.Println("Kartu identitas tersebut tersimpan di loker nomer", r[0])
					} else if id != r[2] {

					} else if !found {
						fmt.Println("Kartu identitas tidak ditemukan", "Result:", k)
					}
				}
			case "search":
				fmt.Scan(&jenis)
				data, err := ioutil.ReadFile("dataloker.txt")
				if err != nil {
					fmt.Println(err)
				}
				a := strings.Split(string(data), "\n")
				var ids string
				for x := 0; x < total; x++ {
					listloker := a[x]
					d := strings.Split(listloker, " ")
					if d[1] == jenis {
						ids += d[2] + ", "
					} else {

					}
				}
				fmt.Println(ids)

			case "exit":
				fmt.Println("Program Terminated")
				null_data := []byte("")
				// the WriteFile method returns an error if unsuccessful
				err := ioutil.WriteFile("dataloker.txt", null_data, 0777)
				// handle this error
				if err != nil {
					// print it out
					fmt.Println(err)
				}
				os.Exit(0)
			case "help":
				fmt.Println("List command: \n status -> to see loker status \n leave [number of loker] -> to delete the spesific order of loker \n find [loker id] -> to see spesific loker by id \n search [type of loker] -> to see spesific loker by type of loker \n exit -> to exit the program \n help -> to see list of command")
			default:
				fmt.Println("Unknow Command \n to see list command insert command 'help'")
			}
		}
	}
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return 0, false
}
