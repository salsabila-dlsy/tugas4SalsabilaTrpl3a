package common

//Struct API
// Order struct (Model) ...

//Message is
type Message struct {
	Code    int     `json:"code"`
	Remark  string  `json:"remark"`
	OrderID string  `json:"orderID"`
	Orders  *Orders `json:"orders,omitempty"`
	Result  *Result `json:"result,omitempty"`
}

//Orders is
type Orders struct {
	OrderID    string         `json:"orderID"`
	CustomerID string         `json:"customerID"`
	EmployeeID string         `json:"employeeID"`
	OrderDate  string         `json:"orderDate"`
	OrdersDet  []OrdersDetail `json:"ordersDetail"`
}

//OrdersDetail is
type OrdersDetail struct {
	OrderID     string  `json:"orderID"`
	ProductID   string  `json:"ProductID"`
	ProductName string  `json:"ProductName"`
	UnitPrice   float64 `json:"UnitPrice"`
	Quantity    int     `json:"Quantity"`
}

//Result is
type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

//Customers is
type Customers struct {
	CustomerID   string `json:"CustomerID"`
	CompanyName  string `json:"CompanyName"`
	ContactName  string `json:"ContactName"`
	ContactTitle string `json:"ContactTitle"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	Country      string `json:"Country"`
	Phone        string `json:"Phone"`
	PostalCode   string `json:"PostalCode"`
}

//Products is
type Products struct {
	ProductID       string `json:"ProductID"`
	SupplierID      string `json:"SupplierID"`
	Category        string `json:"Category"`
	QuantityPerUnit string `json:"QuantityPerUnit"`
	UnitePrice      string `json:"UnitePrice"`
	UnitsInStock    string `json:"UnitsInStock"`
	UnitsOrder      string `json:"UnitsOrder"`
	ReorderLevel    string `json:"ReorderLevel"`
	Discontinued    string `json:"Discontinued"`
}

//End Struct API
