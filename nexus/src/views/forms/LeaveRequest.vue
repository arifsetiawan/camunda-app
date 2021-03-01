<template>
  <div>
    <CRow>
      <CCol lg="12">
        <CCard>
            <CCardHeader>
                <strong>Leave Request </strong><small>Form</small>
            </CCardHeader>
          <CCardBody>
            <CInput
              label="User ID"
              placeholder="User ID"
              readonly
              v-model="profile.id" 
            />
            <CInput
              label="Name"
              placeholder="Name"
              readonly
              v-model="userName"
            />
            <CInput
              label="Email"
              placeholder="Email"
              readonly
              v-model="profile.email" 
            />
            <CInput
              label="Employee Role"
              placeholder="Role"
              readonly
              v-model="employeeData.roleName" 
            />
            <CInput
              label="Leave Days"
              placeholder="Number of leave days"
              v-model="requestLeaveDays" 
              :is-valid="leaveDaysValidator"
            />
            <br>
            <CRow>
                <CCol col="12" class="text-right">
                    <CButton color="primary" class="px-4" @click="submit">Submit</CButton>
                </CCol>
            </CRow>
          </CCardBody>
        </CCard>
      </CCol>
    </CRow>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'LeaveRequest',
  data() {
    return {
      profile: {},
      employeeData: {},
      requestLeaveDays: 0
    };
  },
  computed: {
    userName () {
      return this.profile.firstName + " " + this.profile.lastName
    }
  },
  methods: {
    submit: function() {
      if (this.requestLeaveDays > 0 && this.requestLeaveDays < this.employeeData.leaveDays) {
        var data = {
          variables: {
            id: {
              value: this.profile.id,
              type: "String"
            },
            name: {
              value: this.userName,
              type: "String"
            },
            email: {
              value: this.profile.email,
              type: "String"
            },
            level: {
              value: this.employeeData.roleId,
              type: "String"
            },
            days: {
              value: this.requestLeaveDays,
              type: "Integer"
            }
          }
        }

        axios.post(process.env.VUE_APP_CAMUNDA_PROXY_URL + '/process/start', data)
        .then(resp => {
          console.log(resp)
        })
        .catch(err => {
          console.log(err);
        })
      }
    },
    leaveDaysValidator (val) {
      return val ? val > 0 && val < this.employeeData.leaveDays : false
    }
  },
  created() {
    if (localStorage.getItem('user-profile')) {
      this.profile = JSON.parse(localStorage.getItem('user-profile'))
    }

    axios
      .get(process.env.VUE_APP_CAMUNDA_PROXY_URL + '/employee')
      .then(resp => {
        this.employeeData = resp.data
      })
  }
}
</script>
