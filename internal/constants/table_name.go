package constants

type tableNameStruct struct {
	SecretManagement string
}

var TableName = tableNameStruct{
	SecretManagement: "secret_management",
}
