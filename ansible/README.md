# ANSIBLE

ansible all -m ping --ssh-extra-args='-p 39339' -u HOST -m shell -a "df -h"
ansible-playbook playbook.yml -i inventory
molecule init role my-new-role --driver-name docker





-----------------------------------------
errors:
pip3 install docker
docker pull 