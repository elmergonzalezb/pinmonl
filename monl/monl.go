package monl

import "errors"

var (
	// ErrVendorNotFound indicates that the vendor does not found by name or url
	ErrVendorNotFound = errors.New("monl vendor not found")
)

// Monl holds list of vendors
type Monl struct {
	vendors map[string]Vendor
}

// New creates an instance of Monl
func New() *Monl {
	return &Monl{
		vendors: make(map[string]Vendor),
	}
}

// Add pushes vendor into the list
func (m *Monl) Add(v Vendor) error {
	m.vendors[v.Name()] = v
	return nil
}

// Get finds vendor by name
func (m *Monl) Get(name string) (Vendor, error) {
	if v, ok := m.vendors[name]; ok {
		return v, nil
	}
	return nil, ErrVendorNotFound
}

// GuessURL find vendor by calling vendor's Check()
func (m *Monl) GuessURL(url string) (Vendor, error) {
	for _, v := range m.vendors {
		if v.Check(url) {
			return v, nil
		}
	}
	return nil, ErrVendorNotFound
}