
# Deploy and Setup Camunda

These steps are mentioned in my blog series [My Journey with Camunda](https://medium.com/@arifsetiawan/my-journey-with-camunda-toc-3030da004511). Specifically, [Part 2 - Camunda Deployment](https://medium.com/@arifsetiawan/my-journey-with-camunda-part-2-camunda-deployment-279788a7386e) and [Part 3 - Business Process Case](https://medium.com/@arifsetiawan/my-journey-with-camunda-part-3-business-process-case-da3be25d7258)

## Deploy Camunda

```
docker run -d -p 5436:5432 \
    --name pg-camunda \
    -e POSTGRES_USER=camunda \
    -e POSTGRES_PASSWORD=camunda \
    -v $HOME/camunda/postgres:/var/lib/postgresql/data \
    postgres:12.2

docker run -d --name camunda -p 8080:8080 --link pg-camunda:db \
    -e DB_DRIVER=org.postgresql.Driver \
    -e DB_URL=jdbc:postgresql://db:5432/camunda \
    -e DB_USERNAME=camunda \
    -e DB_PASSWORD=camunda \
    -e WAIT_FOR=db:5432 \
    camunda/camunda-bpm-platform:7.14.0
```

Open Camunda in http://localhost:8080/camunda-welcome/index.html

## Business Case - Leave Request

In [Part 3 - Business Process Case](https://medium.com/@arifsetiawan/my-journey-with-camunda-part-3-business-process-case-da3be25d7258), we use Leave Request workflow as our case study. You can refer to the article for more details explanation.

Download Camunda Modeler from [Camunda website](https://camunda.com/download/modeler/). Open [Leave Request BPMN file](camunda/model/leave-request.bpmn) and [Select Approver DMN file](camunda/model/select-approver.dmn). To know more about BPMN and DMN, refer to Camunda's [BPMN reference](https://docs.camunda.org/manual/7.14/reference/bpmn20/) and [DMN reference](https://docs.camunda.org/manual/7.13/reference/dmn/). 

Using Camunda Modeler, deploy both BPMN and DMN file to your Camunda instance. 

## Create Users and Group

We will add users and groups using [Camunda Admin UI](http://localhost:8080/camunda/app/admin/default/#/) to test our application. 

Add following users and groups

1. Junior Engineer (users: Sam Purple)
1. Senior Engineer (users: Anne Pink)
1. HR (users:John Black)
1. Manager (users: Sophia Green)
1. CEO (users: Mark White)

Create users by running [create-users.sh](camunda/create-users.sh)

```
./camunda/create-users.sh 
```

Create groups by running [create-groups.sh](camunda/create-groups.sh)

```
./camunda/create-groups.sh 
```

