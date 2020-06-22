package importedstructs

type FullNameType func(string, string) string

type ImpPerson struct {
	FirstName, LastName string
	FullName            FullNameType
}
