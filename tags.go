package helpers

import (
	tfjson "github.com/hashicorp/terraform-json"
	. "github.com/onsi/gomega"
)

func ExpectProtectedTag(obj *tfjson.StateResource) {
	Expect(obj.AttributeValues["tags"].(map[string]any)["protected"]).To(Equal("true"))
}
