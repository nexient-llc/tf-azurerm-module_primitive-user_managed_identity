package common

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/nexient-llc/lcaf-component-terratest-common/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("TestSkeletonDeployedIsInvokable", func(t *testing.T) {
		principalId := terraform.Output(t, ctx.TerratestTerraformOptions, "principal_id")

		assert.NotEmptyf(t, principalId, "principal_id must not be empty")
	})
}
