package global

var (
	AUTH_INFO      = "get_auth_info"
	ADD_VISIT_INFO = "add_visit_info"
)
var TOKENS = map[string]map[string]string{
	"admin": {
		"token": "admin-token",
		"password":"admin1234",
	},
	"editor": {
		"token": "editor-token",
		"password":"editor1234",
	},
}
var USERS = map[string]map[string]interface{}{
	"admin-token": {
		"roles": []string{"admin"},
		"introduction": "I am a super administrator",
		"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"name": "Super Admin",
	},	
	"editor-token": {								
		"roles": []string{"editor"},								
		"introduction": "I am an editor",								
		"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",								
		"name": "Normal Editor",								
	},								
}
