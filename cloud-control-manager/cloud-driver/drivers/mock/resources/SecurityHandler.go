// Cloud Driver Interface of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is Mock Driver.
//
// by CB-Spider Team, 2020.10.

package resources

import (
        _ "github.com/sirupsen/logrus"
        cblog "github.com/cloud-barista/cb-log"
	irs "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces/resources"
	"fmt"
)

var securityInfoMap map[string][]*irs.SecurityInfo

type MockSecurityHandler struct {
	MockName      string
}

func init() {
        // cblog is a global variable.
	securityInfoMap = make(map[string][]*irs.SecurityInfo)
}

// (1) create securityInfo object
// (2) insert securityInfo into global Map
func (securityHandler *MockSecurityHandler) CreateSecurity(securityReqInfo irs.SecurityReqInfo) (irs.SecurityInfo, error) {
        cblogger := cblog.GetLogger("CB-SPIDER")
        cblogger.Info("Mock Driver: called CreateSecurity()!")

	mockName := securityHandler.MockName
	securityReqInfo.IId.SystemId = securityReqInfo.IId.NameId

	// (1) create securityInfo object
	securityInfo := irs.SecurityInfo{securityReqInfo.IId,
			securityReqInfo.VpcIID,
			securityReqInfo.Direction,
			securityReqInfo.SecurityRules,
			nil}

	// (2) insert SecurityInfo into global Map
	infoList, _ := securityInfoMap[mockName]
	infoList = append(infoList, &securityInfo)
	securityInfoMap[mockName]=infoList

	return securityInfo, nil
}

func (securityHandler *MockSecurityHandler) ListSecurity() ([]*irs.SecurityInfo, error) {
        cblogger := cblog.GetLogger("CB-SPIDER")
        cblogger.Info("Mock Driver: called ListSecurity()!")
	
	mockName := securityHandler.MockName
	infoList, ok := securityInfoMap[mockName]
	if !ok {
		return []*irs.SecurityInfo{}, nil
	}
	// cloning list of Security
	resultList := make([]*irs.SecurityInfo, len(infoList))
	copy(resultList, infoList)
	return resultList, nil
}

func (securityHandler *MockSecurityHandler) GetSecurity(iid irs.IID) (irs.SecurityInfo, error) {
        cblogger := cblog.GetLogger("CB-SPIDER")
        cblogger.Info("Mock Driver: called GetSecurity()!")

	infoList, err := securityHandler.ListSecurity()
	if err != nil {
		cblogger.Error(err)
		return irs.SecurityInfo{}, err
	}

	for _, info := range infoList {
		if((*info).IId.NameId == iid.NameId) {
			return *info, nil
		}
	}
	
	return irs.SecurityInfo{}, fmt.Errorf("%s SecurityGroup does not exist!!")
}

func (securityHandler *MockSecurityHandler) DeleteSecurity(iid irs.IID) (bool, error) {
        cblogger := cblog.GetLogger("CB-SPIDER")
        cblogger.Info("Mock Driver: called DeleteSecurity()!")

        infoList, err := securityHandler.ListSecurity()
        if err != nil {
                cblogger.Error(err)
                return false, err
        }

	mockName := securityHandler.MockName
        for idx, info := range infoList {
                if(info.IId.NameId == iid.NameId) {
			infoList = append(infoList[:idx], infoList[idx+1:]...)
			securityInfoMap[mockName]=infoList
			return true, nil
                }
        }
	return false, nil
}