package io

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
)

func Test_IoUtil(t *testing.T) {
	bytes, err := ioutil.ReadFile("high_level_permission.json")
	if err != nil {
		log.Errorf("read file fail, cause: %v", err)
		return
	}

	highLevelPermissions := make([]*HighLevelPermission, 0, 0)
	err = json.Unmarshal(bytes, &highLevelPermissions)
	if err != nil {
		log.Errorf("parse json fail, cause: %v", err)
		return
	}

	var viewType, subViewType string
	tempHighLevelPermissions := make([]*HighLevelPermission, 0, 0)
	for _, highLevelPermission := range highLevelPermissions {
		if highLevelPermission.ParentId == "0" {
			if len(tempHighLevelPermissions) > 0 {
				bytes, _ := json.Marshal(tempHighLevelPermissions)
				log.Infof("%s-%s: %s", viewType, subViewType, string(bytes))
				tempHighLevelPermissions = tempHighLevelPermissions[0:0]
			}
			viewType = highLevelPermission.Name
			continue
		}
		if highLevelPermission.OperationOffset == 0 {
			if len(tempHighLevelPermissions) > 0 {
				bytes, _ := json.Marshal(tempHighLevelPermissions)
				log.Infof("%s-%s: %s", viewType, subViewType, string(bytes))
				tempHighLevelPermissions = tempHighLevelPermissions[0:0]
			}
			subViewType = highLevelPermission.Name
			continue
		}
		tempHighLevelPermissions = append(tempHighLevelPermissions, highLevelPermission)
	}
	if len(tempHighLevelPermissions) > 0 {
		bytes, _ := json.Marshal(tempHighLevelPermissions)
		log.Infof("%s-%s: %s", viewType, subViewType, string(bytes))
		tempHighLevelPermissions = tempHighLevelPermissions[0:0]
	}
}

type HighLevelPermission struct {
	ResourceId         string `json:"resourceId"`
	PermissionId       string `json:"permissionId"`
	ResourceType       string `json:"resourceType"`
	ParentResourceType string `json:"parentResourceType"`
	ParentResourceId   string `json:"parentResourceId"`
	ParentId           string `json:"parentId"`
	Name               string `json:"name"`
	OperationOffset    int    `json:"operationOffset"`
}
