RESTSERVER=localhost

VM_ID=41928fab-9362-4bca-87ea-83ebfe15262c
curl -X GET "http://$RESTSERVER:1024/controlvm/$VM_ID?connection_name=cloudit-config01&action=reboot"
