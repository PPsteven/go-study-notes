package main

import (
	"fmt"
	"strconv"
)

var uid int

type Company struct {
	id   int
	name string
	taxCode string
	bankAccount string
}

func NewCompany(name string) *Company {
	uid = uid + 1
	return &Company{id: uid, name:name}
}

func (c *Company) GetCompanyId() string {
	return strconv.Itoa(c.id)
}

func (c *Company) SetTaxCode(code string) {
	c.taxCode = code
}

func (c *Company) SetBankAccount(account string) {
	c.bankAccount = account
}

func (c *Company) Display() {
	fmt.Printf("company name %s, taxcode %s, bank account %s\n", c.name, c.taxCode, c.bankAccount)
}

type Taxation struct {}

func NewTaxation() *Taxation {
	return &Taxation{}
}

func (t *Taxation) applyTaxCode(companyId string) string {
	return fmt.Sprintf("taxcode_%s", companyId)
}

type Bank struct {}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) openAccount(companyId string) string {
	return fmt.Sprintf("account_%s", companyId)
}


type AdminOfIndustry struct {}

func NewAdminOfIndustry() *AdminOfIndustry {
	return &AdminOfIndustry{}
}

func (a *AdminOfIndustry) register(companyName string) *Company {
	return NewCompany(companyName)
}

type CompanyFacade struct {
	admin *AdminOfIndustry
	taxation *Taxation
	bank *Bank
}

func NewCompanyFacade() *CompanyFacade {
	return &CompanyFacade{
		admin: NewAdminOfIndustry(),
		taxation: NewTaxation(),
		bank: NewBank(),
	}
}

func (f *CompanyFacade) OpenCompany(name string) *Company {
	com := f.admin.register(name)
	taxcode := f.taxation.applyTaxCode(com.GetCompanyId())
	account := f.bank.openAccount(com.GetCompanyId())
	com.SetTaxCode(taxcode)
	com.SetBankAccount(account)
	return com
}

func main() {
	facade := NewCompanyFacade()
	alibaba := facade.OpenCompany("Alibaba")
	alibaba.Display()

	tencent := facade.OpenCompany("Tencent")
	tencent.Display()
}
