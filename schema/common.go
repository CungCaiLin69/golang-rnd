package schema

import "github.com/golang-jwt/jwt/v5"

// IUserMatrix equivalent
type UserMatrix struct {
	IsCreate   *bool `json:"is_create,omitempty"`
	IsRead     *bool `json:"is_read,omitempty"`
	IsUpdate   *bool `json:"is_update,omitempty"`
	IsDelete   *bool `json:"is_delete,omitempty"`
	IsUpload   *bool `json:"is_upload,omitempty"`
	IsDownload *bool `json:"is_download,omitempty"`
}

// IPagination equivalent
type Pagination struct {
	Page        int  `json:"page"`
	PageSize    int  `json:"pageSize"`
	TotalRows   *int `json:"totalRows,omitempty"`
	TotalPage   *int `json:"totalPage,omitempty"`
	CurrentPage *int `json:"currentPage,omitempty"`
}

// IDataWithPagination equivalent
type DataWithPagination struct {
	Items      []interface{} `json:"items"`
	Matrix     *UserMatrix   `json:"matrix,omitempty"`
	Pagination *Pagination   `json:"pagination,omitempty"`
}

// IMessages equivalent
type Messages struct {
	Messages []string               `json:"messages"`
	Payload  map[string]interface{} `json:"payload,omitempty"`
}

// IMessagesWith<T> equivalent (using generics in Go 1.18+)
type MessagesWith[T any] struct {
	Messages []string `json:"messages"`
	Payload  *T       `json:"payload,omitempty"`
}

// IJwtVerify equivalent
type JwtVerify struct {
	ID            int     `json:"id"`
	Username      string  `json:"username"`
	Email         *string `json:"email,omitempty"`
	Fullname      *string `json:"fullname,omitempty"`
	Type          string  `json:"type"`
	PrivilegeName *string `json:"privilegeName,omitempty"`
	PrivilegeMode *string `json:"privilegeMode,omitempty"`
	GroupID       int     `json:"groupId"`
	Device        *string `json:"device,omitempty"`
	IPAddress     *string `json:"ipAddress,omitempty"`
	Token         string  `json:"token"`
}

// IJwtCommunicator equivalent (extending JwtVerify)
type JwtCommunicator struct {
	JwtVerify
	UserMatrix *UserMatrix `json:"userMatrix"`
	jwt.RegisteredClaims
}
