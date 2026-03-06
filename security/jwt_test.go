package security

import (
	"testing"

	"example.com/gin_realworld/utils"
)

func TestGenerateJWT(t *testing.T){
	token, err := GenerateJWT("jack", "jack@gmail.com")
	if err != nil {
		t.Error(err)
	}
	t.Logf("token: %v\n", token)
}

func TestVerifyJWT(t *testing.T){
	claim, valid, err := VerifyJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjI2NzgyMTEsImlhdCI6MTc2MjU5MTgxMSwidXNlciI6eyJlbWFpbCI6ImphY2tAZ21haWwuY29tIiwidXNlcm5hbWUiOiJqYWNrIn19.3JELMob51FosWVlncBWVb-0WjbsUQLnIA8Jb21hnBV0")
	if err != nil {
		t.Error(err)
		return 
	}
	t.Logf("current claim: %v\n", utils.JsonMarshal(claim))
	t.Logf("verify jwt: %v\n", valid)
}