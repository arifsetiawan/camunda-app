
function createUser() {
	userId=$1
	userFirstName=$2
	userLastName=$3
	userEmail=$4
	userPassword=$5

    echo ""
    echo "---------------------------"
    echo create user: "$userEmail"
    result=$(curl -s -X POST \
  		--url http://localhost:8080/engine-rest/user/create \
  		--header 'Content-Type: application/json' \
  		--data '{"profile":{"id":"'"$userId"'","firstName":"'"$userFirstName"'","lastName":"'"$userLastName"'","email":"'"$userEmail"'"},"credentials":{"password":"'"$userPassword"'"}}')
    echo result: "$result"
}

createUser "sampurple" "Sam" "Purple" "sam.purple@company.com" "password"
createUser "annepink" "Anne" "Pink" "anne.pink@company.com" "password"
createUser "johnblack" "John" "Black" "john.black@company.com" "password"
createUser "sophiagreen" "Sophia" "Green" "sophia.green@company.com" "password"
createUser "markwhite" "Mark" "White" "mark.white@company.com" "password"
