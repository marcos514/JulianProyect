// Package core implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package core

import (
	"time"
)

//"time"

//Factura this is a factura
type Factura struct {
	ID               int
	ClienteID        int
	cliente          Cliente
	Fecha            time.Time
	PrecioTotal      float32
	facturaProductos []FacturaProducto
}

//FacturaProducto los productos de la factura
type FacturaProducto struct {
	ID         int
	producto   Producto
	ProductoID int
	FacturaID  int
	factura    Factura
	Precio     float32
	Cantidad   int
}

type FacturaProductoInterface interface {
	ThisIsFactProd()
}

//GetPublicFields Guardar Product
func (f *Factura) GetPublicFields() []string {
	return GetPublicFields(f)
}

//GetPublicValues Guardar Product
func (f *Factura) GetPublicValues() []string {
	return GetPublicValues(f)
}

func (f *Factura) GetFacturaProducto() []FacturaProducto {
	return f.facturaProductos
}

func (fp *FacturaProducto) GetProduct() Producto {
	return fp.producto
}

func (fp *FacturaProducto) GetFactura() Factura {
	return fp.factura
}

//GetPublicFields Guardar Product
func (fp *FacturaProducto) GetPublicFields() []string {
	return GetPublicFields(fp)
}

//GetPublicValues Guardar Product
func (fp *FacturaProducto) GetPublicValues() []string {
	return GetPublicValues(fp)
}

func (f *Factura) AddFacturaProduct(fp FacturaProducto) {
	f.facturaProductos = append(f.facturaProductos, fp)
}

func (f *Factura) AppendListFacturasProductos(lfp []FacturaProducto) {
	f.facturaProductos = append(f.facturaProductos, lfp...)
}

// func (f Factura) AppendFactura(lfp []FacturaProducto) []FacturaProducto {
// 	for i := 0; i < len(lfp); i++ {
// 		fp := lfp[i]
// 		fp.factura = f
// 		lfp[i] = fp
// 	}
// 	return lfp
// }

// func AgregarFactura(lfp []FacturaProducto, f Factura) []FacturaProducto {
// 	for i := 0; i < len(lfp); i++ {
// 		fp := lfp[i]
// 		fp.AddFactura(f)
// 		lfp[i] = fp
// 	}
// 	return lfp
// }

func (fp *FacturaProducto) AddFactura(f Factura) {
	fp.factura = f
}

func (fp *FacturaProducto) AddProducto(p Producto) {
	fp.producto = p
}

func (f *Factura) SetCliente(c Cliente) {
	f.cliente = c
}

func (f *Factura) GetCliente() Cliente {
	return f.cliente
}

func GetFacturaProductosByIds(lfp []FacturaProducto) map[int][]FacturaProducto {
	var facturaProductosDict map[int][]FacturaProducto
	for i := 0; i < len(lfp); i++ {
		fp := lfp[i]
		facturaProductosDict[fp.ID] = append(facturaProductosDict[fp.ID], []FacturaProducto{fp}...)
	}
	return facturaProductosDict
}

func (fp *FacturaProducto) IndexFacturaProductoEnLista(lfp []FacturaProducto) int {
	index := -1
	if fp.ID != -6 {
		for i := 0; i < len(lfp); i++ {
			fpaux := lfp[i]
			if fp.ID == fpaux.ID || (fp.ProductoID == fpaux.ProductoID && fp.FacturaID == fpaux.FacturaID) {
				index = i
				break
			}
		}
	}
	return index
}
