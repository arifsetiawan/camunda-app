
function createGroup() {
	groupId=$1
	groupName=$2
	groupType=$3
	userId=$4

    echo ""
    echo "---------------------------"
    echo create group: "$userEmail"
    result=$(curl -s -X POST \
  		--url http://localhost:8080/engine-rest/group/create \
  		--header 'Content-Type: application/json' \
  		--data '{"id":"'"$groupId"'","name":"'"$groupName"'","type":"'"$groupType"'"}')
    echo result: "$result"

	echo add member to group: "$groupId"
    result=$(curl -s -X PUT \
  		--url http://localhost:8080/engine-rest/group/$groupId/members/$userId)
    echo result: "$result"
}

createGroup "juniors" "Junior Engineer" "Organization Unit" "sampurple"
createGroup "seniors" "Senior Engineer" "Organization Unit" "annepink"
createGroup "hr" "HR" "Organization Unit" "johnblack"
createGroup "manager" "Managers" "Organization Unit" "sophiagreen"
createGroup "ceo" "CEO" "Organization Unit" "markwhite"



