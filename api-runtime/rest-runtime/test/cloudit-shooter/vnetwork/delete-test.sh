source ../setup.env

VNETWORK_ID=CB-VNet-powerkim
curl -X DELETE http://$RESTSERVER:1024/vnetwork/$VNETWORK_ID?connection_name=cloudit-config01 |json_pp
