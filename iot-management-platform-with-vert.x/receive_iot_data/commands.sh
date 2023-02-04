# gateway
multipass launch --name gateway --cpus 1 --mem 512M  --disk 3G --network "name=Wi-Fi,mode=auto"  --cloud-init .\vm.cloud-init.yaml
multipass mount app gateway:app

# mqtt
multipass launch --name mqtt --cpus 1 --mem 512M --disk 3G  --network "name=Wi-Fi,mode=auto" --cloud-init .\vm.cloud-init.yaml
multipass mount app mqtt:app


multipass launch --name devices --cpus 1 --mem 512M --disk 3G --network "name=Wi-Fi,mode=auto" --cloud-init .\vm.cloud-init.yaml
multipass mount app devices:app

multipass launch --name things --cpus 2 --mem 2G --disk 6G --network "name=Wi-Fi,mode=auto" --cloud-init .\vm.cloud-init.yaml
multipass mount app things:app



192.168.8.106 gateway.home.smart
192.168.8.109 things.home.smart
192.168.8.108 mqtt.home.smart
192.168.8.110 devices.home.smart
