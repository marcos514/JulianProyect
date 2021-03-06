package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/marcos514/julianBE/modules/core"
	csvmodule "github.com/marcos514/julianBE/modules/csvModule"
	//"julian_project/morestrings"
)

type empData struct {
	Name string
	Age  string
	City string
}

// type Cliente struct {
// id        int
// mail      string
// nombre    string
// direccion string
// numero    string
// }
type Usuario struct {
	nombre   string
	apellido string
}

func (u Usuario) getDatosPersonales() string {
	return fmt.Sprintf("%s, %s", u.nombre, u.apellido)
}

type Administrador struct {
	Usuario
	sector string
}

func (a Administrador) getDatosCompletos() string {
	return fmt.Sprintf("%s - %s", a.getDatosPersonales(), a.sector)
}

type Primate interface {
	Alimentar(string)
}

type Antropoide struct{}

func (t Antropoide) Alimentar(fruta string) {
	fmt.Printf("Comiendo %s \n", fruta)
}

type Gorila struct {
	Antropoide
}

func DarDeComer(primate Primate) {
	primate.Alimentar("banana")
}

func main() {
	fmt.Println("THIS IS THE DIRECTORY")
	dir, _ := os.UserHomeDir()
	fmt.Println(dir)
	path, _ := os.Getwd()
	fmt.Println(path)

	csvmodule.SetFullPath(path)

	//fmt.Println(morestrings.ReverseRunes("!oG ,ollasdasdeH"))

	var clientes = core.Cliente{}

	core.ReadFromFile("cliente.json", &clientes)
	res2B, _ := json.Marshal(clientes)
	fmt.Println(string(res2B))

	var listClientes = []core.Cliente{
		{
			ID:        1,
			Mail:      "marmarreyer@gmail.com",
			Nombre:    "Marcos",
			Direccion: "Amenedo 622",
			Numero:    "1549168959",
		},
		{
			ID:        2,
			Mail:      "reymarcos51@gmail.com",
			Nombre:    "Marcos Rey",
			Direccion: "Amenedo 622",
			Numero:    "1549168959",
		},
	}
	core.WriteData("clientes_2.json", listClientes)

	var clienteTest = core.Cliente{
		ID:        1,
		Mail:      "marmarreyer@gmail.com",
		Nombre:    "Marcos",
		Direccion: "Amenedo 622",
		Numero:    "1549168959",
	}
	core.WriteData("marcos.json", clienteTest)

	// archivo.readFromFile("asd.json", cliente)
	// fmt.Println(archivo.addLine())
	// csvFile, err := os.Open("emp2.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(csvFile)
	// defer csvFile.Close()

	// csvLines, err := csv.NewReader(csvFile).ReadAll()
	// if err != nil {
	// 	fmt.Println(csvLines)
	// }
	// for other, line := range csvLines {
	// 	fmt.Println(other)
	// 	emp := empData{
	// 		Name: line[0],
	// 		Age:  line[1],
	// 		City: line[2],
	// 	}
	// 	fmt.Println(emp.Name + " " + emp.Age + " " + emp.City)
	// }
	var administrador = Administrador{Usuario{"Jose", "Luis"}, "Computos"}

	fmt.Println(administrador.getDatosPersonales())
	fmt.Println(administrador.getDatosCompletos())
	kong := Gorila{}
	DarDeComer(kong)

	products := []csvmodule.Producto{{
		Producto: core.Producto{
			ID: 0, CantidadUnidad: 50,
			Nombre: "Marcos", Descripcion: "Rey", Medidas: "15*45", Empresa: "MarcosSA", Codigo: "1234",
			Precio:     154,
			Categorias: []string{"marcos", "12345"},
			Activo:     true,
		},
	}}
	prod := csvmodule.Producto{
		core.Producto{
			ID: 1, CantidadUnidad: 50,
			Nombre: "Primer Producto Bueno", Descripcion: "Redsady", Medidas: "15*454", Empresa: "MarcosSA2", Codigo: "2459",
			Precio:     154,
			Categorias: []string{"sad", "sad", "15asdsa4"},
			Activo:     true,
		},
	}
	products = append(products, prod)
	csvmodule.GuardarProductos(products)

	fact := csvmodule.Factura{
		core.Factura{
			ID:          1,
			ClienteID:   0,
			Fecha:       time.Now(),
			PrecioTotal: 12,
		},
	}

	var facts []csvmodule.Factura

	factProd := csvmodule.FacturaProducto{
		core.FacturaProducto{
			ID:         1,
			ProductoID: prod.ID,
			FacturaID:  fact.ID,
			Precio:     prod.Precio,
			Cantidad:   5,
		},
	}
	factProd2 := csvmodule.FacturaProducto{
		core.FacturaProducto{
			ID:         2,
			ProductoID: 2,
			FacturaID:  fact.ID,
			Precio:     prod.Precio,
			Cantidad:   5,
		},
	}
	fact.AddFacturaProduct(factProd.FacturaProducto)
	fact.AddFacturaProduct(factProd2.FacturaProducto)
	fact.AddFacturaProduct(factProd2.FacturaProducto)
	fact.AddFacturaProduct(factProd2.FacturaProducto)
	facts = append(facts, fact)
	fact.PrecioTotal = 100000

	// factProd.AddFactura(fact.Factura)
	// factProd.AddProducto(prod.Producto)
	csvmodule.ActualizarFactura(fact)

	// ahora := time.Now()
	// fmt.Printf("Ahora %v\n", ahora)

	cliente := []csvmodule.Cliente{{
		core.Cliente{
			ID:        0,
			Mail:      "marcos@smallsforsmalls.com",
			Nombre:    "Marcos Rey",
			Direccion: "Amenedo 622",
			Numero:    "1549168959",
		},
	}}
	csvmodule.GuardarClientes(cliente)
	csvproductos := csvmodule.ReadProductos("./store/productos.csv")
	fmt.Printf("PROYEEEEEEEECCCSSS %v\n\n\n\n", csvproductos)
	// clientesCsv := csvmodule.ReadClientes()
	csvfac := csvmodule.ReadFacturas()
	res2B, err := json.Marshal(csvfac[0].GetFacturaProducto())
	if err != nil {

	}
	fmt.Println(string(res2B))
	fmt.Printf("facturaaaaaaaa fecha %v\n", csvfac[0].Fecha.Format(time.ANSIC))
	csvfac[0].PrintFactura()

	ids := []int{1}
	csvmodule.DeshabilitarProductosById(ids)

	fmt.Printf("%v", csvmodule.AgregarProductosDeArchivo("./store/newproductos.csv"))

	productos := []csvmodule.Producto{{
		Producto: core.Producto{
			ID: 0, CantidadUnidad: 50,
			Nombre: "MarcosSSSSSSSS", Descripcion: "ReyYYYYYYYYYYYYY", Medidas: "15*45", Empresa: "MarcosSA", Codigo: "1234",
			Precio:     6666,
			Categorias: []string{"marcos", "12345"},
			Activo:     true,
		},
	}}
	csvmodule.ActualizarProductos(productos)

	csvfac[0].ImprimirFacturaCSV()

}
