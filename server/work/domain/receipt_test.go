package domain

import "testing"

func Test_Validate_Success(t *testing.T) {
	receipt_test := Receipt{
		Date: "2024/09/01",
		Category: "meal",
		Content: "A portion of dinner",
		Money: 500,
	}

	err := receipt_test.Validate()
	if err != nil{
		t.Fatal("Failed test")
	}
}

func Test_Validate_Failure(t *testing.T) {
	receipt_test := Receipt{
		Date: "2024/09/01",
		Category: "meal",
		Content: "A portion of dinner",
	}

	err := receipt_test.Validate()
	if err == nil{
		t.Fatal("Failed test")
	}

}