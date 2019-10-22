RESTSERVER=node12

curl -X POST http://$RESTSERVER:1024/vm?connection_name=azure-config01 -H 'Content-Type: application/json' -d '{ "VMName": "CBVm", "ImageId": "/subscriptions/cb592624-b77b-4a8f-bb13-0e5a48cae40f/resourceGroups/CB-GROUP/providers/Microsoft.Compute/images/CB-IMG", "NetworkInterfaceId": "/subscriptions/cb592624-b77b-4a8f-bb13-0e5a48cae40f/resourceGroups/CB-GROUP/providers/Microsoft.Network/networkInterfaces/CB-VNic", "SecurityGroupIds": ["/subscriptions/cb592624-b77b-4a8f-bb13-0e5a48cae40f/resourceGroups/CB-GROUP/providers/Microsoft.Network/networkSecurityGroups/CB-SecGroup"], "VMSpecId": "Standard_B1ls", "VMUserId": "cb-user", "VMUserPasswd": "cb-userCBUSER!@#"}'
