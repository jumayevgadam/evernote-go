package reqvalidator

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates struct.
func ValidateStruct(ctx context.Context, s interface{}) error {
	return validate.StructCtx(ctx, s)
}

// ReadRequest func reads request.
func ReadRequestJSON(c *gin.Context, s interface{}) error {
	if err := c.ShouldBindJSON(s); err != nil {
		return fmt.Errorf("error binding request: %w", err)
	}

	if err := validate.StructCtx(c, s); err != nil {
		return fmt.Errorf("error validating struct: %w", err)
	}

	return nil
}
