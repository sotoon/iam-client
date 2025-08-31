package webhook

import (
	"encoding/json"
	"fmt"

	"github.com/sotoon/iam-client/pkg/types"
)

type IAMObject interface{}

func parseIAMObjectJson(objectType IAMObjectType, jsonData []byte) (IAMObject, error) {
	switch objectType {
	case IAMObjectTypeUser:
		var user types.WebhookUser
		if err := json.Unmarshal(jsonData, &user); err != nil {
			return nil, err
		}
		return user, nil
	case IAMObjectTypeWorkspace:
		var workspace types.WebhookWorkspace
		if err := json.Unmarshal(jsonData, &workspace); err != nil {
			return nil, err
		}
		return workspace, nil
	case IAMObjectTypeUserWorkspace:
		var userWorkspace types.WebhookUserWorkspaceRelation
		if err := json.Unmarshal(jsonData, &userWorkspace); err != nil {
			return nil, err
		}
		return userWorkspace, nil
	default:
		return nil, fmt.Errorf("unsupported object type %s", objectType)
	}
}
