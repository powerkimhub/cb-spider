package resources

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	idrv "../../../interfaces"
	irs "../../../interfaces/resources"
	compute "google.golang.org/api/compute/v1"
)

type GCPPublicIPHandler struct {
	Region     idrv.RegionInfo
	Ctx        context.Context
	Client     *compute.Service
	Credential idrv.CredentialInfo
}

// @TODO: PublicIP 리소스 프로퍼티 정의 필요
type PublicIPInfo struct {
	Id                string
	Name              string
	Region            string // GCP
	CreationTimestamp string // GCP
	Address           string // GCP
	NetworkTier       string // GCP : PREMIUM, STANDARD
	AddressType       string // GCP : External, INTERNAL, UNSPECIFIED_TYPE
	Status            string // GCP : IN_USE, RESERVED, RESERVING
	InstanceId        string // GCP : 연결된 VM
}

func (publicIpHandler *GCPPublicIPHandler) CreatePublicIP(publicIPReqInfo irs.PublicIPReqInfo) (irs.PublicIPInfo, error) {
	projectID := publicIpHandler.Credential.ProjectID
	region := publicIpHandler.Region.Region
	publicIpName := publicIPReqInfo.Name
	address := &compute.Address{
		Name: publicIpName,
	}
	publicIpHandler.Client.Addresses.Insert(projectID, region, address).Do()
	time.Sleep(time.Second * 3)

	publicIPInfo, err := publicIpHandler.GetPublicIP(publicIpName)
	if err != nil {
		log.Fatal(err)
	}

	return publicIPInfo, err
}

func (publicIpHandler *GCPPublicIPHandler) ListPublicIP() ([]*irs.PublicIPInfo, error) {
	projectID := publicIpHandler.Credential.ProjectID
	region := publicIpHandler.Region.Region

	list, err := publicIpHandler.Client.Addresses.List(projectID, region).Do()
	if err != nil {
		log.Fatal(err)
	}
	var publicIpInfoArr []*irs.PublicIPInfo
	for _, item := range list.Items {
		var publicInfo irs.PublicIPInfo
		publicInfo.Name = item.Name
		publicInfo.PublicIP = item.Address
		publicInfo.Status = item.Status
		//publicInfo.KeyValueList = GetKeyValueList()
	}
	return publicIpInfoArr, nil
}

func (publicIpHandler *GCPPublicIPHandler) GetPublicIP(publicIPID string) (irs.PublicIPInfo, error) {
	projectID := publicIpHandler.Credential.ProjectID
	region := publicIpHandler.Region.Region
	name := publicIPID // name or resource ID

	info, err := publicIpHandler.Client.Addresses.Get(projectID, region, name).Do()
	if err != nil {
		log.Fatal(err)
	}

	//바인딩 하기위해 []byte로 변환 처리
	infoByte, err := info.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}

	var publicInfo irs.PublicIPInfo
	var keyValueList []irs.KeyValue

	publicInfo.Name = info.Name
	publicInfo.PublicIP = info.Address
	if users := info.Users; users != nil {
		vmArr := strings.Split(users[0], "/")
		publicInfo.OwnedVMID = vmArr[len(vmArr)-1]
	}
	publicInfo.Status = info.Status

	// 구조체 안에 해당값을 바인딩해준다.
	var result map[string]interface{}

	json.Unmarshal(infoByte, &result)
	keyValueList = GetKeyValueList(result)
	// for key, value := range result {
	// 	keyValueList = append(keyValueList, irs.KeyValue{key, value})
	// }

	publicInfo.KeyValueList = keyValueList

	if err != nil {
		log.Fatal(err)
	}

	return publicInfo, err
}

func (publicIpHandler *GCPPublicIPHandler) DeletePublicIP(publicIPID string) (bool, error) {

	return true, nil
}

// func (*GCPPublicIPHandler) mappingPublicIpInfo(infos []byte) (irs.PublicIPInfo, error) {
// 	var publicInfo irs.PublicIPInfo
// 	err := json.Unmarshal(infos, &publicInfo)

// 	return publicInfo
// }
